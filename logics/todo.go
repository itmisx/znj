package logics

import (
	"strconv"
	"time"

	"znj/models"

	"github.com/go-playground/validator/v10"
	"github.com/itmisx/errorx"
	"github.com/itmisx/go-helper"
	"gorm.io/gorm"
)

// Todo 待办
type Todo struct{}

// BoardTasks 待办任务列表
type BoardTasks struct {
	models.Todos
	Name        string `json:"name"`
	ProjectName string `json:"project_name"`
	IsExpired   int8   `json:"is_expired"`
}

// List 待办列表
func (td Todo) List(db *gorm.DB,
	todoTypeID string,
	doneStatus int16,
	assignTo string,
	adminID string,
) (list []BoardTasks, err error) {
	// 待办列表定义
	var (
		count    int64 = 0
		todoType models.TodoType
	)
	// 项目待办
	if todoTypeID != "" {

		// 验证待办类型合法性
		db.Table("todo_type").
			Where("id = ?", todoTypeID).
			First(&todoType).
			Count(&count)
		if count <= 0 {
			err = errorx.New("", 203111)
			return
		}

		// 项目权限验证
		pl := Project{}
		err = pl.HasAccessAuthority(db, todoType.ProjectID, adminID)
		if err != nil {
			return
		}

		query := db.Table("todos").
			Select("todos.*,u.name").
			Joins("left join users u on u.id = todos.assign_to").
			Where("todo_type_id = ?", todoTypeID)

		// 默认待办状态为已打开，包括待完成和已完成
		if doneStatus == -1 {
			query = query.Where("done_status <> 4")
		} else {
			query = query.Where("done_status = ?", doneStatus)
		}

		if assignTo != "-1" && assignTo != "" {
			query = query.Where("assign_to = ?", assignTo)
		}

		query.Order("FIELD(`done_status`,2,1,3,4)").
			Order("level asc").
			Order("create_time desc").
			Find(&list)
	} else {
		// 我的待办
		db.Table("todos td").
			Where("assign_to = ?", adminID).
			Where("done_status = ?", 1).
			Joins("left join projects p on p.id = td.project_id").
			Select("td.*,p.name as project_name").
			Order("td.level asc").
			Order("td.create_time desc").
			Find(&list)
	}
	for i, td := range list {
		endTimeInt, _ := strconv.ParseInt(td.EndTime, 10, 64)
		if endTimeInt > 0 {
			// 时间戳转日期
			list[i].EndTime = (&helper.Datetime{}).WithTimestamp(endTimeInt).ToDate("Y-m-d")
			// 判断任务是否过期
			if time.Now().Unix() >= endTimeInt {
				list[i].IsExpired = 1
			} else {
				list[i].IsExpired = 0
			}
		} else {
			list[i].EndTime = ""
		}
	}

	return
}

// Save 保存待办(含新增和编辑)
// return errCode
// return rowsAffected
func (td Todo) Save(db *gorm.DB,
	todo models.Todos,
	adminID string,
) (err error) {
	var (
		todoType         models.TodoType
		todoExist        models.Todos
		statusChangeFlag int64
		count            int64  = 0
		primaryKeyID     string = ""
		rowsAffected     int64  = 0
	)
	// 获取操作人id

	todo.CreateID = adminID

	// 设置更新时间
	todo.UpdateTime = time.Now().Unix()

	// 编辑
	if todo.ID != "" {

		// 如果为编辑，先判断id的合法性
		rowsAffected = db.Model(&todoExist).Where("id = ?", todo.ID).
			First(&todoExist).RowsAffected

		if rowsAffected != 1 {
			err = errorx.New("", 203003)
			return
		}

		// 优先等级
		if todo.Level > 0 {
			todoExist.Level = todo.Level
		}

		// 待办内容
		if todo.Content != "" {
			todoExist.Content = todo.Content
		}

		// 指派人
		if todo.AssignTo != "" {
			todoExist.AssignTo = todo.AssignTo
		}

		if todo.Desc != "" {
			todoExist.Desc = todo.Desc
		}
		// 截止时间
		if todo.EndTime != "" {
			timestamp, _ := (&helper.Datetime{}).ToTime(todo.EndTime + " 23:59:59")
			todoExist.EndTime = strconv.FormatInt(timestamp, 10)
		} else {
			todoExist.EndTime = "0"
		}
		// 完成状态
		if todo.DoneStatus > 0 {
			if todo.DoneStatus != todoExist.DoneStatus {
				statusChangeFlag = 1
			}
			todoExist.DoneStatus = todo.DoneStatus
		}

		// 更新时间
		todoExist.UpdateTime = time.Now().Unix()

		// 字段验证
		err = td.FieldCheck(&todoExist)
		if err != nil {
			return
		}

		// 更新内容
		rowsAffected := db.Model(&models.Todos{}).
			Where("id = ?", todo.ID).
			Updates(&todoExist).
			RowsAffected
		if rowsAffected == 1 {
			// 添加操作记录
			if statusChangeFlag == 1 {
				switch todo.DoneStatus {
				case 1:
					SaveHistory(db, "todo", "reopen", todo.ID, adminID)
				case 2:
					SaveHistory(db, "todo", "start", todo.ID, adminID)
				case 3:
					SaveHistory(db, "todo", "finish", todo.ID, adminID)
				case 4:
					SaveHistory(db, "todo", "close", todo.ID, adminID)
				}
			}
			return
		}
		err = errorx.New("", 203002)
		return

	}
	/* 新增 */

	// 设置主键id
	primaryKeyID, err = helper.UUID{}.String()
	if err != nil {
		return
	}
	todo.ID = primaryKeyID
	// 设置创建时间
	todo.CreateTime = time.Now().Unix()
	todo.UpdateTime = time.Now().Unix()

	// 完成状态
	if todo.DoneStatus <= 0 {
		todo.DoneStatus = 1
	}
	if todo.AssignTo == "" {
		todo.AssignTo = adminID
	}
	// 字段验证
	err = td.FieldCheck(&todo)
	if err != nil {
		return
	}

	// 待办类型
	if todo.TodoTypeID != "" {
		// 验证待办类型合法性
		db.Table("todo_type").
			Where("id = ?", todo.TodoTypeID).
			First(&todoType).
			Count(&count)
		if count <= 0 {
			err = errorx.New("", 203111)
		}

		// 项目权限验证
		pl := Project{}
		err = pl.HasAccessAuthority(db, todoType.ProjectID, adminID)
		if err != nil {
			return
		}
		// 关联项目
		todo.ProjectID = todoType.ProjectID
	}
	// 截止时间
	if todo.EndTime != "" {
		timestamp, _ := (&helper.Datetime{}).ToTime(todo.EndTime + " 23:59:59")
		todo.EndTime = strconv.FormatInt(timestamp, 10)
	} else {
		todo.EndTime = "0"
	}

	// 添加操作记录
	SaveHistory(db, "todo", "create", primaryKeyID, adminID)

	/* 保存数据 */
	rowsAffected = db.Create(&todo).RowsAffected

	if rowsAffected < 1 {
		err = errorx.New("", 203001)
		return
	}

	return
}

// GetHistory 获取任务历史记录
// 历史记录除了记录待办的操作还会记录其它，这里默认的记录类型为“todo”
//
// @param relatedID 关联id
func (td Todo) GetHistory(db *gorm.DB,
	relatedID string,
) (historyList interface{}, err error) {
	var list []struct {
		models.History
		UserName      string `json:"user_name"`
		CreateTimeStr string `json:"create_time_str"`
	}

	db.Table("history h").
		Joins("left join users u on u.id = h.create_id").
		Select("h.*,FROM_UNIXTIME(h.create_time) as create_time_str,u.name as user_name").
		Where("h.history_type = ?", "todo").
		Where("h.related_id = ?", relatedID).
		Order("h.create_time desc").
		Find(&list)
	historyList = list
	return
}

// FieldCheck 字段验证
// return int errCode 错误码
func (td Todo) FieldCheck(todolist *models.Todos) (err error) {
	validate := validator.New()
	err = validate.Struct(todolist)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			field := e.Field()
			tag := e.ActualTag()
			switch field {
			case "ID":
				if tag == "required" {
					err = errorx.New("", 203101)
					return
				}
			case "Level":
				if tag == "required" {
					err = errorx.New("", 203102)
					return
				} else if tag == "min" || tag == "max" {
					err = errorx.New("", 203109)
					return
				}
			case "Content":
				if tag == "required" {
					err = errorx.New("", 203103)
					return
				}
			case "DoneStatus":
				if tag == "required" {
					err = errorx.New("", 203104)
					return
				} else if tag == "min" || tag == "max" {
					err = errorx.New("", 203110)
					return
				}
			case "AssignTo":
				if tag == "required" {
					err = errorx.New("", 203105)
					return
				}
			case "CreateId":
				if tag == "required" {
					err = errorx.New("", 203106)
					return
				}
			case "CreateTime":
				if tag == "required" {
					err = errorx.New("", 203107)
					return
				}
			case "UpdateTime":
				if tag == "required" {
					err = errorx.New("", 203108)
					return
				}
			}
		}
	}
	return
}
