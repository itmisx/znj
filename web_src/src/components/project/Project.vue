<template>
  <div class="main">
    <div class="project-header">
      <!--项目名称-->
      <div
        style="
          display: flex;
          display: -webkit-flex;
          flex-direction: row;
          align-items: center;
        "
      >
        <div style="font-size: 24px; font-weight: bold; margin-right: 30px">
          {{ projectName }}
        </div>
      </div>
      <!--项目管理按钮-->
      <div style="margin-top: 15px">
        <span
          class="projectManageBtn"
          @click="switchItem('board')"
          :style="{ fontWeight: activeProjectItem == 0 ? 'bold' : 'normal' }"
          >{{ $t("lang.board") }}
        </span>
        <span
          class="projectManageBtn"
          @click="switchItem('doc')"
          :style="{ fontWeight: activeProjectItem == 1 ? 'bold' : 'normal' }"
          >{{ $t("lang.documents") }}
        </span>
        <span
          class="projectManageBtn"
          @click="switchItem('member')"
          :style="{ fontWeight: activeProjectItem == 2 ? 'bold' : 'normal' }"
          >{{ $t("lang.members") }}
        </span>
      </div>
    </div>
    <el-divider style="margin: 0"></el-divider>
    <div class="project-body">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: "Project",
  data() {
    return {
      projectName: "",
      projectID: 0,
      activeProjectItem: 0,
    };
  },
  methods: {
    switchItem(item) {
      switch (item) {
        case "board":
          this.activeProjectItem = 0;
          this.$router.push({
            path: "/main/project/board",
          });
          break;
        case "doc":
          this.activeProjectItem = 1;
          this.$router.push({
            path: "/main/project/doc",
          });
          break;
        case "member":
          this.activeProjectItem = 2;
          this.$router.push({
            path: "/main/project/member",
          });
          break;
        default:
          break;
      }
    },
  },
  mounted() {
    //设置项目子页面
    if (this.$route.path == "/main/project") {
      this.$router.push({
        path: "/main/project/board",
      });
    }
    if (this.$route.path == "/main/project/board") {
      this.activeProjectItem = 0;
    } else if (this.$route.path == "/main/project/doc") {
      this.activeProjectItem = 1;
    } else if (this.$route.path == "/main/project/member") {
      this.activeProjectItem = 2;
    }
    // 获取项目ID
    this.projectID = parseInt(sessionStorage.getItem("projectID"));
    // 获取项目名称
    this.projectName = sessionStorage.getItem("projectName");
  },
};
</script>

<style scoped>
.main {
  height: 100%;
  width: 100%;
  margin: 0;
  display: flex;
  display: -webkit-flex;
  flex-direction: column;
}
.project-header {
  height: 75px;
  width: 800px;
  margin: 0px auto;
}
.el-divider--horizontal {
  margin: 0;
}
.projectManageBtn {
  font-size: 14px;
  margin-right: 25px;
}
.projectManageBtn:hover {
  color: #25b864;
  cursor: pointer;
}
.project-body {
  flex: 1;
  overflow-x: show;
}
</style>
