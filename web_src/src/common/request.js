/**
 * 请求方法封装
 */
import fetch from "./fetch.js";
import * as api from "./api.js";

//是否登录
export function isSignedIn(params) {
  return fetch({
    url: api.index_is_signed_in,
    method: "post",
  });
}
//用户登录
export function signin(params) {
  return fetch({
    url: api.index_signin,
    method: "post",
    data: params,
  });
}
//用户登录
export function signout(params) {
  return fetch({
    url: api.index_signout,
    method: "post",
  });
}

/******************用户*******************/

// 修改密码
export function changePassword(params) {
  return fetch({
    url: api.user_change_password,
    method: "post",
    data: params,
  });
}

/******************待办*******************/

// 待办清单
export function todoList(params) {
  return fetch({
    url: api.todo_list,
    method: "post",
    data: params,
  });
}
// 增加待办
export function todoSave(params) {
  return fetch({
    url: api.todo_save,
    method: "post",
    data: params,
  });
}

// 待办状态变动历史
export function todoHistory(params) {
  return fetch({
    url: api.todo_history,
    method: "post",
    data: params,
  });
}

/******************团队*******************/

// 团队列表
export function teamList(params) {
  return fetch({
    url: api.team_list,
    method: "post",
    data: params,
  });
}
// 团队成员
export function teamMemberList(params) {
  return fetch({
    url: api.team_member_list,
    method: "post",
    data: params,
  });
}
// 添加团队成员
export function addTeamMember(params) {
  return fetch({
    url: api.add_team_member,
    method: "post",
    data: params,
  });
}
// 团队新增/编辑 保存
export function teamSave(params) {
  return fetch({
    url: api.team_save,
    method: "post",
    data: params,
  });
}
// 设置工作团队
export function setWorkTeam(params) {
  return fetch({
    url: api.set_work_team,
    method: "post",
    data: params,
  });
}
// 获取工作团队
export function getWorkTeam(params) {
  return fetch({
    url: api.get_work_team,
    method: "post",
    data: params,
  });
}
// 重置密码
export function resetPassword(params) {
  return fetch({
    url: api.team_reset_password,
    method: "post",
    data: params,
  });
}
// 设为管理员
export function setAdminRole(params) {
  return fetch({
    url: api.team_set_admin_role,
    method: "post",
    data: params,
  });
}
/******************项目******************/

//---项目

// 获取项目列表
export function projectList(params) {
  return fetch({
    url: api.project_list,
    method: "post",
    data: params,
  });
}
// 项目新增/编辑 保存
export function projectSave(params) {
  return fetch({
    url: api.project_save,
    method: "post",
    data: params,
  });
}

//---项目看板

//获取待办类型
export function projectTodoTypeList(params) {
  return fetch({
    url: api.project_todo_type_list,
    method: "post",
    data: params,
  });
}
//新增待办类型
export function projectTodoTypeSave(params) {
  return fetch({
    url: api.project_todo_type_save,
    method: "post",
    data: params,
  });
}
//待办类型批量排序
export function projectTodoTypeBatchSort(params) {
  return fetch({
    url: api.project_todo_type_batch_sort,
    method: "post",
    data: params,
  });
}

//---项目成员

// 获取项目成员列表
export function projectMemberList(params) {
  return fetch({
    url: api.project_member_list,
    method: "post",
    data: params,
  });
}
// 设置项目成员
export function projectMemberSet(params) {
  return fetch({
    url: api.project_member_set,
    method: "post",
    data: params,
  });
}

//---文档

//文档列表
export function docList(params) {
  return fetch({
    url: api.doc_list,
    method: "post",
    data: params,
  });
}
//保存文档
export function docSave(params) {
  return fetch({
    url: api.doc_save,
    method: "post",
    data: params,
  });
}
//章节批量排序
export function chapterBatchSort(params) {
  return fetch({
    url: api.chapter_batch_sort,
    method: "post",
    data: params,
  });
}
//章节列表
export function chapterList(params) {
  return fetch({
    url: api.chapter_list,
    method: "post",
    data: params,
  });
}
//章节保存
export function chapterSave(params) {
  return fetch({
    url: api.chapter_save,
    method: "post",
    data: params,
  });
}
//获取章节内容
export function chapterContent(params) {
  return fetch({
    url: api.chapter_content,
    method: "post",
    data: params,
  });
}
//保存章节内容
export function chapterContentSave(params) {
  return fetch({
    url: api.chapter_content_save,
    method: "post",
    data: params,
  });
}
//章节历史版本
export function chapterContentList(params) {
  return fetch({
    url: api.chapter_content_list,
    method: "post",
    data: params,
  });
}
//获取文档版本
export function getChapterRevision(params) {
  return fetch({
    url: api.get_chapter_revision,
    method: "post",
    data: params,
  });
}
//设置文档版本
export function setChapterRevision(params) {
  return fetch({
    url: api.set_chapter_revision,
    method: "post",
    data: params,
  });
}
//获取最近访问的项目和章节
export function getLatestProjectsAndChapters(params) {
  return fetch({
    url: api.get_latestProjectsAndChapters,
    method: "post",
    data: params,
  });
}
