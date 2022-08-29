<template>
  <div class="main">
    <div
      v-if="todoItemShow"
      class="todoItem"
      style="text-decoration: line-through"
      :style="{
        color: color,
        background: background,
        textDecoration: textDecoration,
      }"
    >
      <div class="content">{{ content }}</div>
      <div style="flex: 1">
        <div
          v-if="this.done_status === 1"
          @click.stop
          style="float: right; display: table-cell"
        >
          <el-dropdown @command="handleCommand" trigger="click">
            <span class="el-dropdown-link">
              <i
                class="iconfont icon-menu"
                style="color: gray; margin: auto 10px"
              ></i>
            </span>
            <el-dropdown-menu
              slot="dropdown"
              style="width: 120px; text-align: center; font-size: 14px"
            >
              <el-dropdown-item command="todoedit">
                <i class="iconfont icon-edit"></i>编辑
              </el-dropdown-item>
              <el-dropdown-item command="tododone">
                <i class="iconfont icon-check-circle"></i>完成
              </el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "TodolistItem",
  props: {
    // 待办的id
    todoid: {
      type: String,
      required: true,
    },
    // 紧急程度
    level: {
      type: Number,
      required: true,
    },
    // 完成状态
    done_status: {
      type: Number,
      required: true,
    },
    // 待办内容
    content: {
      type: String,
      required: true,
    },
    // 截止时间
    end_time: {
      type: String,
    },
    // 关联的项目名称
    project_name: {
      type: String,
    },
  },
  data() {
    return {
      item: "padding:5px 5px;height:38px;",
      todoId: this.todoid,
      todoLevel: this.level,
      todoContent: this.content,
      todoItemShow: true,
      todoEditShow: false,
      todo: {
        id: this.todoid,
        projectName: this.project_name,
        content: this.content,
        level: this.level,
        end_time: this.end_time,
        desc: this.desc,
      },
    };
  },
  computed: {
    // 文字颜色
    color() {
      switch (this.level) {
        case 1:
          return "#F56C6C";
        case 2:
          return "#E6A23C";
        case 3:
          return "#2693FE";
        case 4:
          return "#67C23A";
      }
    },
    // 背景色
    background() {
      switch (this.level) {
        case 1:
          return "#FEF0F0";
        case 2:
          return "#FDF6EC";
        case 3:
          return "#E8F3FF";
        case 4:
          return "#F0F9EB";
      }
    },
    // 删除线
    textDecoration() {
      if (this.done_status === 2) return "line-through";
      else return "none";
    },
  },
  methods: {
    handleCommand(command) {
      switch (command) {
        case "todoedit":
          this.$emit("click", this.todo);
          break;
        case "tododone":
          this.todoDone();
          break;
      }
    },
    // 待办完成
    todoDone: function () {
      let data = {
        id: this.todoid,
        done_status: 2,
      };
      this.Request.todoSave(data)
        .then((response) => {
          if (response.data.code === 0) {
            this.$emit("todoList");
            this.cancelTodoEdit();
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {
          this.cancelTodoEdit();
        });
    },
  },
};
</script>

<style scoped>
.main {
  width: 100%;
}
.append-button {
  display: -webkit-flex;
  display: flex;
  flex-direction: row;
  align-items: center;
}
.todoItem {
  width: 100%;
  margin: 5px 0px;
  height: 38px;
  vertical-align: middle;
  cursor: pointer;
  display: -webkit-flex;
  display: flex;
  flex-direction: row;
  align-items: center;
}
.todoEdit {
  width: 100%;
  margin: 5px 0px;
  vertical-align: middle;
}
.content {
  margin-left: 5px;
  width: 700px;
  font-size: 14px;
  display: inline-block;
  vertical-align: middle;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
