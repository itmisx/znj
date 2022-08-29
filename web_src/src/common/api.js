/*------------------接口定义------------------*/
//公共入口
export const index_is_signed_in = "/index/isSignedIn"; //是否登录
export const index_signup = "/index/signup"; //注册
export const index_signin = "/index/signin"; //登录
export const index_signout = "/index/signout"; //登出
export const index_find_password = "/index/findpwd"; //找回密码
//用户
export const user_change_password = "/user/changePassword"; //修改密码

//todolist
export const todo_list = "/todo/list"; //获取待办列表
export const todo_save = "/todo/save"; //添加待办
export const todo_history = "/todo/history"; //待办任务的

//项目
export const project_list = "/project/list"; //获取项目列表
export const project_save = "/project/save"; //项目新增或编辑保存

//项目看板
export const project_todo_type_list = "/project/todoTypeList"; //项目任务清单
export const project_todo_type_save = "/project/todoTypeSave"; //增加任务\
export const project_todo_type_batch_sort = "project/todoTypeBatchSort"; //任务清单批量排序

//项目成员
export const project_member_list = "/project/memberList"; //项目成员列表
export const project_member_set = "/project/memberSet"; //项目成员设置

/*文档*/
export const doc_list = "/doc/list"; //文档列表
export const doc_save = "/doc/save"; //保存文档
export const chapter_list = "/doc/chapterList"; //章节列表
export const chapter_save = "/doc/chapterSave"; //章节保存
export const chapter_batch_sort = "/doc/chapterBatchSort"; //章节批量排序
export const chapter_content = "/doc/chapterContent"; //获取章节内容
export const chapter_content_save = "/doc/chapterContentSave"; //保存章节内容
export const chapter_content_list = "/doc/chapterContentList"; //章节历史版本
export const get_chapter_revision = "/doc/getChapterRevision"; //获取章节版本内容
export const set_chapter_revision = "/doc/setChapterRevision"; //章节内容回滚
export const get_latestProjectsAndChapters =
  "/doc/getLatestProjectsAndChapters"; //获取最近访问的项目和章节
/*团队*/

//团队列表
export const team_switch = "/team/switch"; //切换团队
export const team_save = "/team/save"; //创建团队
export const team_list = "/team/list"; //团队列表
export const team_member_list = "/team/memberList"; //团队成员
export const add_team_member = "/team/addTeamMember"; //添加团队成员
export const set_work_team = "/team/setWorkTeam"; //设置工作团队
export const get_work_team = "/team/getWorkTeam"; //获取工作团队

//团队分组
export const team_group_list = "/team/group-list"; //团队分组
export const team_group_add = "/team/group-add"; //新增分组
export const team_group_edit = "/team/group-edit"; //编辑分组

//团队成员
export const team_member_add = "/team/member-add"; //添加团队成员
export const team_member_del = "/team/member-del"; //删除团队成员
export const team_reset_password = "/team/resetPassword"; //重置密码
export const team_set_admin_role = "/team/setAdminRole"; //设为管理员
