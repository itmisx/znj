<template>
  <div class="main">
    <div class="left"></div>
    <div class="center">
      <div class="list-tbar">
        <el-button
          style="width: 120px"
          class="el-icon-plus"
          size="mini"
          type="primary"
          @click="addToDo"
          v-if="addToDoBtnDisplay"
          >{{ $t("lang.addTodo") }}
        </el-button>
      </div>
      <div class="list">
        <TodolistItem
          v-for="(item, index) in todolist"
          v-on:todoList="todoList"
          @click.native="editTodo(item)"
          @click="editTodo(item)"
          :todoid="item.id"
          :level="item.level"
          :content="item.content"
          :end_time="item.end_time"
          :done_status="item.done_status"
          :project_name="item.project_name"
          :key="index"
        >
        </TodolistItem>
      </div>
    </div>
    <div class="right">
      <div class="latest-title">{{ $t("lang.latestProjects") }}</div>
      <el-card
        class="latest-projects"
        shadow="never"
        body-style="width:auto;display:flex;display:-webkit-flex;flex-direction:row;
                        flex-wrap:wrap;justify-content:flex-start;"
      >
        <div
          v-for="project in latestProjects"
          style="
            width: 33%;
            display: flex;
            display: -webkit-flex;
            flex-direction: column;
            align-items: center;
            cursor: pointer;
            margin-top: 10px;
          "
          @click="openProject(project.id, project.name)"
          :key="project.id"
        >
          <el-button type="primary" size="mini" circle>
            <div style="width: 12px">{{ project.name[0] }}</div>
          </el-button>
          <div
            style="
              font-size: 12px;
              width: 80px;
              text-overflow: ellipsis;
              overflow: hidden;
              white-space: nowrap;
              display: inline-block;
              text-align: center;
            "
          >
            {{ project.name }}
          </div>
        </div>
        <div v-if="!latestProjects.length">
          <i class="el-icon-cloudy">????????????</i>
        </div>
      </el-card>
      <div class="latest-title">{{ $t("lang.latestDocuments") }}</div>
      <el-card
        class="latest-projects"
        shadow="never"
        body-style="width:100%;display:flex;display:-webkit-flex;
                    flex-direction:column;align-items:flex-start;justify-content:center;"
      >
        <div
          v-for="chapter in latestChapters"
          :key="chapter.id"
          @click="
            openDoc(
              chapter.project_id,
              chapter.project_name,
              chapter.doc_id,
              chapter.doc_name,
              chapter.doc_type,
              chapter.id,
              chapter.name
            )
          "
          style="
            margin: 5px;
            display: flex;
            display: -webkit-flex;
            flex-direction: row;
            align-items: center;
            cursor: pointer;
          "
        >
          <i class="iconfont icon-file-text"></i>
          <span
            style="
              font-size: 14px;
              width: 190px;
              text-overflow: ellipsis;
              overflow: hidden;
              white-space: nowrap;
              display: inline-block;
              text-align: left;
            "
            >{{ chapter.name }}</span
          >
        </div>
        <div v-if="!latestChapters.length">
          <i class="el-icon-cloudy"></i>????????????
        </div>
      </el-card>
      <img src="@/assets/img/p1.jpeg" style="margin-top: 20px; width: 100%" />
      <div class="latest-title">Make it Work, Make it Right,Make it Fast!</div>
    </div>
    <!-- ???????????? -->
    <el-dialog
      :title="add ? $t('lang.addTodo') : $t('lang.editTodo')"
      :close-on-click-modal="closeOnClickModal"
      :visible.sync="editTodoDialogVisible"
      destroy-on-close
      append-to-body
      width="800px"
    >
      <el-form
        ref="todoForm"
        label-width="auto"
        :rules="rules"
        style="padding: 10px; height: 400px; overflow-y: auto"
        :model="todo"
      >
        <el-form-item
          :label="$t('lang.taskTitle')"
          size="mini"
          prop="content"
          required
        >
          <el-input v-model="todo.content" style="width: 200px" placeholder="">
          </el-input>
        </el-form-item>
        <!-- ???????????? -->
        <el-form-item
          :label="$t('lang.relatedProject')"
          size="mini"
          v-if="todo.projectName"
        >
          <el-tag type="success">{{ todo.projectName }}</el-tag>
        </el-form-item>
        <!-- ???????????? -->
        <el-form-item
          :label="$t('lang.taskLevel')"
          size="mini"
          prop="level"
          required
        >
          <el-select v-model="todo.level" style="width: 200px" placeholder="">
            <el-option label="??????" :value="1">
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
                <span>??????</span>
              </div>
            </el-option>
            <el-option label="??????" :value="2">
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
                <span>??????</span>
              </div>
            </el-option>
            <el-option label="??????" :value="3">
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
                <span>??????</span>
              </div>
            </el-option>
            <el-option label="??????" :value="4">
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
                <span>??????</span>
              </div>
            </el-option>
          </el-select>
        </el-form-item>
        <!-- ???????????? -->
        <el-form-item :label="$t('lang.endDate')" size="mini" prop="end_time">
          <el-date-picker
            v-model="todo.end_time"
            style="width: 200px"
            value-format="yyyy-MM-dd"
            type="date"
            placeholder="????????????"
          >
          </el-date-picker>
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
  </div>
</template>

<script>
import Editor from "../editor/TextEditor";
import TodolistItem from "./TodolistItem";
export default {
  name: "Todolist",
  components: { Editor, TodolistItem },
  data() {
    return {
      add: true,
      closeOnClickModal: false,
      editTodoDialogVisible: false,
      activeIndex: "0",
      searchKeyWord: "",
      addToDoDisplay: false,
      addToDoBtnDisplay: true,
      todoAddContent: "",
      todoAddLevel: 1,
      todolist: [],
      latestProjects: [],
      latestChapters: [],
      todo: {
        id: 0,
        projectName: "",
        content: "",
        level: 1,
        end_time: "",
        desc: ""
      },
      rules: {
        content: [
          {
            type: "string",
            required: true,
            message: "??????????????????",
            trigger: "blur"
          }
        ],
        level: [
          {
            type: "number",
            required: true,
            message: "????????????????????????",
            trigger: "change"
          }
        ]
      }
    };
  },
  methods: {
    // ?????????????????????
    getEditorContent(content) {
      this.todo.desc = content;
    },
    editTodo(item) {
      this.add = false;
      this.editTodoDialogVisible = true;
      this.todo.id = item.id;
      this.todo.projectName = item.project_name;
      this.todo.content = item.content;
      this.todo.level = item.level;
      this.todo.end_time = item.end_time;
      this.todo.desc = item.desc;
      this.$nextTick(() => {
        this.$refs.wangEditor.setValue(item.desc);
      });
    },
    //????????????????????????
    addToDo: function() {
      this.add = true;
      this.editTodoDialogVisible = true;
      this.$nextTick(() => {
        this.resetForm();
        this.$refs.wangEditor.setValue("");
      });
    },
    //????????????????????????
    cancelAddToDo: function() {
      this.addToDoDisplay = false;
      this.addToDoBtnDisplay = true;
    },
    // ??????????????????
    todoSave() {
      let data = {
        id: this.todo.id,
        level: parseInt(this.todo.level),
        content: this.todo.content,
        end_time: this.todo.end_time,
        desc: this.todo.desc
      };
      this.Request.todoSave(data)
        .then(response => {
          if (response.data.code == 0) {
            this.todoList();
          } else {
            this.$message(response.data.msg);
          }
          this.resetForm();
          this.cancelTodoEdit();
        })
        .catch(err => {
          this.cancelTodoEdit();
        });
    },
    // ??????????????????
    todoList: function() {
      this.Request.todoList().then(response => {
        if (response.data.code == 0) {
          this.todolist = response.data.data.todo_lists;
          this.latestProjects = response.data.data.latest_projects;
          this.latestChapters = response.data.data.latest_chapters;
        } else {
          this.$message(response.data.msg);
        }
      });
    },
    // ????????????????????????????????????
    getLatestProjectsAndChapters() {
      this.Request.getLatestProjectsAndChapters().then(response => {
        if (response.data.code == 0) {
          this.latestProjects = response.data.data.projects
            ? response.data.data.projects
            : [];
          this.latestChapters = response.data.data.chapters
            ? response.data.data.chapters
            : [];
        } else {
          this.$message(response.data.msg);
        }
      });
    },
    // ??????????????????
    openProject(projectID, projectName) {
      sessionStorage.setItem("projectID", projectID);
      sessionStorage.setItem("projectName", projectName);
      this.$router.push({
        path: "/main/project"
      });
    },
    // ????????????
    openDoc(
      projectID,
      projectName,
      docID,
      docName,
      docType,
      chapterID,
      chapterName
    ) {
      sessionStorage.setItem("projectID", projectID);
      sessionStorage.setItem("projectName", projectName);
      sessionStorage.setItem("docID", docID);
      sessionStorage.setItem("docType", docType);
      sessionStorage.setItem("docName", docName);
      sessionStorage.setItem("chapterID", chapterID);
      sessionStorage.setItem("chapterName", chapterName);
      this.$router.push({
        path: "/main/project/doc"
      });
    },
    // ??????????????????
    resetForm() {
      this.$refs.todoForm.resetFields();
      this.todo = {
        id: 0,
        content: "",
        level: 1,
        end_time: "",
        desc: ""
      };
    },
    //??????????????????
    cancelTodoEdit() {
      this.editTodoDialogVisible = false;
    }
  },
  mounted() {
    // ??????????????????
    this.todoList();
  }
};
</script>

<style scoped>
.main {
  width: 1200px;
  margin-left: auto;
  margin-right: auto;
  display: -webkit-flex;
  display: flex;
  flex-direction: row;
}
.left {
  width: 200px;
}
.center {
  flex: 1;
}
.right {
  margin-top: 40px;
  width: 300px;
  margin-left: 20px;
}
.list {
  margin-top: 10px;
  overflow-y: auto;
}
.latest-title {
  margin: 5px 5px 5px 5px;
  font-size: 14px;
  color: rgb(140, 140, 140);
}
.latest-projects {
  font-size: 14px;
  height: auto;
  color: rgb(140, 140, 140);
}
</style>
