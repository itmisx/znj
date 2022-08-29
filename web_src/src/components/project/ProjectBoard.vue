<template>
  <draggable
    handle=".todotype-header"
    class="board-main"
    v-model="todoTypeList"
    @start="beforeMove"
    @end="move"
  >
    <TodoType
      v-for="item in todoTypeList"
      style="flex: 0 0 320px; height: 99%"
      :key="item.id"
      :typeID="item.id"
      v-on:todoTypeList="getTodoTypeList"
      :typeName="item.name"
    >
    </TodoType>
    <div
      class="add-task-category"
      slot="footer"
      @dragstart.prevent
      @dragstart.stop
    >
      <el-card class="add-task-category-text">
        <div @click="todoTypeAddShow = true" v-if="!todoTypeAddShow">
          <i class="el-icon-plus"></i>
          <span>{{ $t("lang.addTaskList") }}</span>
        </div>
        <el-input
          size="small"
          :placeholder="$t('lang.pleaseInputTaskType')"
          v-model="todoTypeName"
          @keyup.enter.native="addTodoType"
          v-if="todoTypeAddShow"
        >
          <div slot="append">
            <el-button @click="addTodoType">{{ $t("lang.confirm") }}</el-button>
            <el-divider direction="vertical"></el-divider>
            <el-button @click="todoTypeAddShow = false">{{
              $t("lang.cancel")
            }}</el-button>
          </div>
        </el-input>
      </el-card>
    </div>
  </draggable>
</template>

<script>
import draggable from "vuedraggable";
import TodoType from "./TodoType";
export default {
  name: "ProjectBoard",
  components: { TodoType, draggable },
  data() {
    return {
      todoTypeAddShow: false,
      projectID: 0,
      todoTypeName: "",
      todoTypeList: [],
      todoTypeListBeforeMove: [],
    };
  },
  methods: {
    //新增待办类型对话框
    addTodoType() {
      let data = {
        project_id: this.projectID,
        name: this.todoTypeName,
      };
      this.projectTodoTypeSave(data);
    },
    projectTodoTypeSave(data) {
      this.Request.projectTodoTypeSave(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.getTodoTypeList();
            this.todoTypeAddShow = false;
            this.todoTypeName = "";
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    //开始移动
    beforeMove(evt) {
      //保存旧的任务类型顺序
      this.todoTypeListBeforeMove = this.todoTypeList;
    },
    //结束移动
    move(evt) {
      let idSorts = "";
      //比较移动前和移动后的数据顺序，仅提交变化的
      for (let i = 0; i < this.todoTypeList.length; i++) {
        if (
          this.todoTypeList[i].id !== this.todoTypeListBeforeMove[i].id ||
          this.todoTypeList[i].sort == 0
        ) {
          if (idSorts != "") {
            idSorts = idSorts + "," + this.todoTypeList[i].id + "-" + i;
          } else {
            idSorts = this.todoTypeList[i].id + "-" + i;
          }
        }
      }
      if (idSorts == "") {
        return;
      }
      //保存新的排序
      let data = {
        id_sorts: idSorts,
      };
      this.Request.projectTodoTypeBatchSort(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.getTodoTypeList();
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    //获取待办类型列表
    getTodoTypeList() {
      let data = {
        project_id: this.projectID,
      };
      this.Request.projectTodoTypeList(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.todoTypeList = response.data.data;
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
  },
  mounted() {
    //获取项目ID
    this.projectID = sessionStorage.getItem("projectID");
    //加载待办类型列表
    this.getTodoTypeList();
  },
};
</script>

<style scoped>
.board-main {
  height: 100%;
  width: 100%;
  box-sizing: border-box;
  -moz-box-sizing: border-box;
  padding: 5px 5px;
  display: flex;
  display: -webkit-flex;
  flex-direction: row;
  /* align-items:stretch; */
  overflow-y: hidden;
  overflow-x: auto;
}
.add-task-category {
  flex: 0 0 300px;
  padding: 5px 5px;
  height: 60px;
  text-align: center;
  cursor: pointer;
}
.add-task-category-text:hover {
  color: #25b864;
}
</style>
