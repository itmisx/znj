<template>
  <div class="todotype-main">
    <el-card class="todotype-box-card" :body-style="bodyStyle" shadow="never">
      <!-- 任务类型，卡片header -->
      <div slot="header" class="todotype-header">
        <div v-if="!todoTypeEditShow">
          <span style="font-weight: 500">{{ todoTypeName }}</span>
          <div
            style="
              float: right;
              padding: 0;
              max-height: 250px;
              overflow-y: auto;
            "
          >
            <i class="el-icon-plus task-category-manage" @click="addTodo"></i>
            <!-- 指派人筛选 -->
            <el-dropdown @command="handleTodoAssignSelect" trigger="click">
              <i
                class="iconfont icon-user task-category-manage"
                style="color: gray; margin: auto 5px"
              ></i>
              <el-dropdown-menu
                slot="dropdown"
                style="width: 140px; text-align: left; font-size: 14px"
              >
                <!-- 全部 -->
                <el-dropdown-item command="-1">
                  {{ $t("lang.all") }}
                </el-dropdown-item>
                <!-- 人员列表-->
                <el-dropdown-item
                  :command="item.id"
                  v-for="item in projectMembers"
                  :key="item.id"
                >
                  {{ item.name }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
            <!-- 待办清单操作 -->
            <el-dropdown @command="handleTodoTypeCommand" trigger="click">
              <i
                class="iconfont icon-menu task-category-manage"
                style="color: gray; margin: auto 5px"
              ></i>
              <el-dropdown-menu
                slot="dropdown"
                style="width: 160px; text-align: left; font-size: 14px"
              >
                <!-- 编辑 -->
                <el-dropdown-item command="todoTypeEdit">
                  <i class="iconfont icon-edit-square"></i>{{ $t("lang.edit") }}
                </el-dropdown-item>
                <!-- 查看已打开 -->
                <el-dropdown-item command="listOpenedTask">
                  <i class="iconfont icon-reloadtime"></i
                  >{{ $t("lang.viewOpenedTask") }}
                </el-dropdown-item>
                <!-- 查看已关闭 -->
                <el-dropdown-item command="listClosedTask">
                  <i class="iconfont icon-close-circle"></i
                  >{{ $t("lang.viewClosedTask") }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </div>
        </div>
        <el-input
          size="small"
          :placeholder="$t('lang.pleaseInputTaskType')"
          v-model="todoTypeName"
          @keyup.enter.native="editTodoType"
          v-if="todoTypeEditShow"
        >
          <div slot="append">
            <el-button @click="editTodoType">{{
              $t("lang.confirm")
            }}</el-button>
            <el-divider direction="vertical"></el-divider>
            <el-button @click="todoTypeEditShow = false">{{
              $t("lang.cancel")
            }}</el-button>
          </div>
        </el-input>
      </div>
      <!-- 任务列表，body -->
      <div
        style="
          overflow-y: auto;
          height: 100%;
          padding: 10px 10px;
          background: #fdfdfd;
        "
      >
        <!-- 任务项 -->
        <el-card
          :style="{
            borderLeftColor: levelColor[item.level - 1],
            backgroundColor: item.done_status == 3 ? '#e9e9e9' : '#fff',
          }"
          shadow="always"
          body-style="padding:10px"
          v-for="item in todoList"
          :key="item.id"
          class="task"
          style="
            border-left-width: 4px;
            border-top-color: #ccc;
            border-right-color: #ccc;
            border-bottom-color: #ccc;
          "
        >
          <div>
            <div
              style="
                display: flex;
                display: -webkit-flex;
                flex-direction: row;
                align-items: flex-start;
              "
            >
              <!-- 任务标题 -->
              <div :style="{ flex: 1 }">
                <!-- 进行状态 -->
                <el-tag
                  size="mini"
                  v-if="item.done_status == 2"
                  type="success"
                  effect="light"
                  >进行中
                  <i class="el-icon-loading"></i>
                </el-tag>
                <!-- 已完成 -->
                <el-tag
                  size="mini"
                  v-if="item.done_status == 3"
                  type="info"
                  effect="plain"
                  >已完成
                  <i class="el-icon-circle-check"></i>
                </el-tag>
                {{ item.content }}
              </div>
              <!-- 任务操作菜单 -->
              <el-dropdown
                @command="
                  (command) => {
                    handleTaskCommand(command, item);
                  }
                "
                @click.stop.native
                trigger="click"
              >
                <i
                  class="iconfont icon-ellipsis"
                  style="color: gray; margin: auto 5px; font-size: 28px"
                ></i>
                <el-dropdown-menu
                  slot="dropdown"
                  style="width: 120px; text-align: left; font-size: 14px"
                >
                  <!-- 编辑 -->
                  <el-dropdown-item command="edit">
                    <i class="el-icon-edit"></i
                    >{{ $t("lang.todoControl.edit") }}
                  </el-dropdown-item>
                  <el-dropdown-item
                    command="start"
                    v-if="item.done_status == 1"
                  >
                    <i class="el-icon-video-play"></i
                    >{{ $t("lang.todoControl.start") }}
                  </el-dropdown-item>
                  <el-dropdown-item command="stop" v-if="item.done_status == 2">
                    <i class="el-icon-video-pause"></i
                    >{{ $t("lang.todoControl.stop") }}
                  </el-dropdown-item>
                  <!-- 已解决 -->
                  <el-dropdown-item
                    command="finish"
                    v-if="item.done_status < 3"
                  >
                    <i class="iconfont icon-check-circle"></i
                    >{{ $t("lang.todoControl.finish") }}
                  </el-dropdown-item>
                  <!-- 关闭 -->
                  <el-dropdown-item command="close" v-if="item.done_status < 4">
                    <i class="iconfont icon-close-circle"></i
                    >{{ $t("lang.todoControl.close") }}
                  </el-dropdown-item>
                  <!-- 重新打开 -->
                  <el-dropdown-item
                    command="reopen"
                    v-if="item.done_status == 4"
                  >
                    <i class="iconfont icon-undo"></i
                    >{{ $t("lang.todoControl.reopen") }}
                  </el-dropdown-item>
                  <!-- 历史记录 -->
                  <el-dropdown-item command="history">
                    <i class="iconfont icon-branches"></i
                    >{{ $t("lang.history") }}
                  </el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </div>
            <!-- 分割线 -->
            <el-divider></el-divider>
            <!-- 任务截止时间、指派人 -->
            <div
              style="
                display: flex;
                flex-direction: row-reverse;
                align-items: center;
              "
            >
              <!-- 指派人 -->
              <span>{{ item.name }}</span>
              <el-avatar shape="circle" fit="fill" :size="20">{{
                item.name[0]
              }}</el-avatar>
              <!-- 截止日期 -->
              <div style="flex: 1">
                <!-- 已完成 -->
                <el-tag
                  size="mini"
                  v-if="item.end_time && item.end_time != '0'"
                  :type="
                    item.done_status <= 2 && item.is_expired == 1
                      ? 'danger'
                      : 'info'
                  "
                  effect="plain"
                >
                  <i class="el-icon-date"></i>{{ item.end_time }}
                </el-tag>
              </div>
            </div>
          </div>
        </el-card>
      </div>
    </el-card>
    <!-- 新增任务 -->
    <el-dialog
      :title="editTodoDialogTitle"
      :close-on-click-modal="closeOnClickModal"
      :visible.sync="editTodoDialogVisible"
      :destroy-on-close="true"
      width="960px"
    >
      <el-form ref="todoForm" label-width="auto" :rules="rules" :model="todo">
        <el-form-item
          :label="$t('lang.taskTitle')"
          size="mini"
          prop="content"
          required
        >
          <el-input v-model="todo.content" style="width: 200px" placeholder="">
          </el-input>
        </el-form-item>
        <el-form-item
          :label="$t('lang.taskAssign')"
          size="mini"
          prop="assign_to"
          required
        >
          <el-select
            v-model="todo.assign_to"
            style="width: 200px"
            placeholder=""
          >
            <el-option
              v-for="item in projectMembers"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('lang.endDate')" size="mini" prop="end_time">
          <el-date-picker
            v-model="todo.end_time"
            style="width: 200px"
            value-format="yyyy-MM-dd"
            type="date"
            placeholder="选择日期"
          >
          </el-date-picker>
        </el-form-item>
        <el-form-item
          :label="$t('lang.taskLevel')"
          size="mini"
          prop="level"
          required
        >
          <el-select v-model="todo.level" style="width: 200px" placeholder="">
            <el-option label="最高" :value="1">
              <div
                style="
                  height: 100%;
                  display: -webkit-flex;
                  display: flex;
                  flex-direction: row;
                  align-items: center;
                "
              >
                <div
                  style="width: 15px; height: 15px; background: #f56c6c"
                ></div>
                <span>最高</span>
              </div>
            </el-option>
            <el-option label="较高" :value="2">
              <div
                style="
                  height: 100%;
                  display: -webkit-flex;
                  display: flex;
                  flex-direction: row;
                  align-items: center;
                "
              >
                <div
                  style="width: 15px; height: 15px; background: #e6a23c"
                ></div>
                <span>较高</span>
              </div>
            </el-option>
            <el-option label="普通" :value="3">
              <div
                style="
                  height: 100%;
                  display: -webkit-flex;
                  display: flex;
                  flex-direction: row;
                  align-items: center;
                "
              >
                <div
                  style="width: 15px; height: 15px; background: #2693fe"
                ></div>
                <span>普通</span>
              </div>
            </el-option>
            <el-option label="较低" :value="4">
              <div
                style="
                  height: 100%;
                  display: -webkit-flex;
                  display: flex;
                  flex-direction: row;
                  align-items: center;
                "
              >
                <div
                  style="width: 15px; height: 15px; background: #67c23a"
                ></div>
                <span>较低</span>
              </div>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('lang.taskDesc')" size="mini">
          <Editor :catchData="getEditorContent" ref="wangEditor"></Editor>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button size="mini" @click="editTodoDialogVisible = false">{{
          $t("lang.cancel")
        }}</el-button>
        <el-button size="mini" type="primary" @click="todoSave">{{
          $t("lang.confirm")
        }}</el-button>
      </span>
    </el-dialog>
    <!-- 待办变动历史 -->
    <el-dialog
      :title="$t('lang.history')"
      :visible.sync="showTodoHistory"
      width="600px"
    >
      <div style="height: 250px; overflow-y: auto">
        <el-timeline>
          <el-timeline-item
            v-for="item in historyList"
            :key="item.id"
            :timestamp="item.create_time_str"
          >
            <b>{{ item.user_name }}</b
            >{{ $t("lang.todoOperationText." + item.operation_type) }}
          </el-timeline-item>
        </el-timeline>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="showTodoHistory = false">关闭</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import Editor from "../editor/TextEditor";
import md5 from "js-md5";

export default {
  name: "TodoType",
  components: { Editor },
  props: {
    typeName: "",
    typeID: "",
  },
  data() {
    return {
      historyList: [],
      showTodoHistory: false,
      projectID: 0,
      todoList: [],
      projectMembers: [],
      todoTypeEditShow: false,
      todoTypeName: this.typeName,
      todoTypeID: this.typeID,
      bodyStyle: "padding:0px;font-size:12px;flex:1;height:0;",
      closeOnClickModal: false,
      editTodoDialogVisible: false,
      editTodoDialogTitle: "新增任务",
      levelColor: ["#F56C6C", "#E6A23C", "#2693FE", "#67C23A"],
      taskListType: "opened",
      taskListDoneStatus: -1,
      todo: {
        id: 0,
        content: "", //任务标题
        assign_to: "", //指派人
        end_time: "", //截止时间
        level: 1, //重要程度
        desc: "", //任务描述
      },
      rules: {
        content: [
          {
            type: "string",
            required: true,
            message: "待办内容必填",
            trigger: "blur",
          },
        ],
        assign_to: [
          {
            type: "string",
            required: true,
            message: "任务必须指派人",
            trigger: "blur",
          },
        ],
        level: [
          {
            type: "number",
            required: true,
            message: "任务紧急程度必填",
            trigger: "blur",
          },
        ],
      },
    };
  },
  methods: {
    dragstart() {
      console.log("dragstart");
    },
    hashName(name) {
      let hashName =
        "https://secure.gravatar.com/avatar/" + md5(name) + "?d=identicon";
      return hashName;
    },
    // 获取编辑器内容
    getEditorContent(content) {
      this.todo.desc = content;
    },
    //待办类型处理
    handleTodoTypeCommand(command) {
      switch (command) {
        case "todoTypeEdit":
          this.todoTypeEditShow = true;
          break;
        case "todoTypeStar":
          this.starTodoType();
          break;
        case "listOpenedTask":
          this.taskListDoneStatus = -1;
          this.getTodoList();
          break;
        case "listClosedTask":
          this.taskListDoneStatus = 4;
          this.getTodoList();
          break;
        default:
          break;
      }
    },
    // 任务处理
    handleTaskCommand(commond, item) {
      switch (commond) {
        case "edit":
          this.editTodo(item);
          break;
        case "start":
          this.changeDoneStatus(2, item.id);
          break;
        case "stop":
          this.changeDoneStatus(1, item.id);
          break;
        case "finish":
          this.changeDoneStatus(3, item.id);
          break;
        case "close":
          this.changeDoneStatus(4, item.id);
          break;
        case "reopen":
          this.changeDoneStatus(1, item.id);
          break;
        case "history":
          this.Request.todoHistory({ related_id: item.id })
            .then((response) => {
              this.showTodoHistory = true;
              if (response.data.code == 0) {
                this.historyList = response.data.data;
              } else {
                this.$message(response.data.msg);
              }
            })
            .catch((err) => {});
          break;
        default:
          break;
      }
    },
    //指派人筛选
    handleTodoAssignSelect(assign_to) {
      let data = {
        todo_type_id: this.todoTypeID,
        done_status: this.taskListDoneStatus,
        assign_to: assign_to,
      };
      this.Request.todoList(data).then((response) => {
        if (response.data.code == 0) {
          this.todoList = response.data.data;
        } else {
          this.$message(response.data.msg);
        }
      });
    },
    //待办类型编辑
    editTodoType() {
      let data = {
        id: this.todoTypeID,
        name: this.todoTypeName,
      };
      this.Request.projectTodoTypeSave(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.todoTypeEditShow = false;
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    //获取待办任务列表
    getTodoList() {
      let data = {
        todo_type_id: this.todoTypeID,
        done_status: this.taskListDoneStatus,
      };
      this.Request.todoList(data).then((response) => {
        if (response.data.code == 0) {
          this.todoList = response.data.data;
        } else {
          this.$message(response.data.msg);
        }
      });
    },
    //获取项目成员
    getMemberList() {
      let data = {
        project_id: this.projectID,
      };
      this.Request.projectMemberList(data)
        .then((response) => {
          if (response.data.code == 0) {
            let list = response.data.data;
            this.projectMembers = list;
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    // 新增任务
    addTodo() {
      this.editTodoDialogTitle = "新增任务";
      this.editTodoDialogVisible = true;
      this.$nextTick(() => {
        this.resetForm();
        this.$refs.wangEditor.setValue("");
      });
    },
    // 编辑任务
    editTodo(item) {
      let todo = {
        id: item.id,
        content: item.content, //任务标题
        assign_to: item.assign_to, //指派人
        end_time: item.end_time, //截止时间
        level: item.level, //重要程度
        desc: item.desc, //任务描述
      };
      this.todo = todo;
      this.editTodoDialogTitle = "编辑任务";
      this.editTodoDialogVisible = true;
      this.$nextTick(() => {
        this.$refs.wangEditor.setValue(item.desc);
      });
    },
    // 任务编辑保存
    todoSave() {
      this.$refs.todoForm.validate((success, object) => {
        if (success == false) {
          this.$message("任务输入有误");
        } else {
          let data = {
            id: this.todo.id,
            level: parseInt(this.todo.level),
            content: this.todo.content,
            desc: this.todo.desc,
            assign_to: this.todo.assign_to,
            end_time: this.todo.end_time,
            todo_type_id: this.todoTypeID,
          };
          this.Request.todoSave(data)
            .then((response) => {
              if (response.data.code == 0) {
                this.getTodoList();
              } else {
                this.$message(response.data.msg);
              }
              this.resetForm();
              this.cancelTodoEdit();
            })
            .catch((err) => {
              this.cancelTodoEdit();
            });
        }
      });
    },
    // 复位表格内容
    resetForm() {
      this.$refs.todoForm.resetFields();
      this.todo = {
        id: 0,
        content: "",
        assign_to: "",
        level: 1,
        desc: "",
      };
    },
    //取消待办编辑
    cancelTodoEdit() {
      this.editTodoDialogVisible = false;
    },
    //待办完成状态变更
    changeDoneStatus(status, todo_id) {
      let data = {
        id: todo_id,
        done_status: status,
      };
      this.Request.todoSave(data)
        .then((response) => {
          if (response.data.code === 0) {
            this.getTodoList();
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
  },
  mounted() {
    // 设置工作项目
    this.projectID = sessionStorage.getItem("projectID");
    // 加载项目成员
    this.getMemberList();
    //加载待办任务
    this.getTodoList();
  },
};
</script>

<style scoped>
.todotype-header {
  cursor: move;
}
.todotype-main {
  width: 100%;
  height: 100%;
  padding: 5px 5px;
  overflow-y: hidden;
  overflow-x: hidden;
}
.todotype-box-card {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  overflow-x: hidden;
}
.task-category-manage {
  cursor: pointer;
}
.task-category-manage:hover {
  color: #25b864;
}
.task {
  margin-bottom: 5px;
}
.el-divider--horizontal {
  margin: 10px 0px;
}
</style>
