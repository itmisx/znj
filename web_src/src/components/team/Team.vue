<template>
  <div class="main">
    <div class="left"></div>
    <div class="center">
      <router-view></router-view>
    </div>
    <div class="right">
      <el-menu
        background-color="#F9F9F9"
        router
        default-active="/main/team/team-list"
        class="el-menu-vertical"
      >
        <el-menu-item index="/main/team/team-list" class="el-memu-item-new">
          <i class="el-icon-s-grid"></i>
          <span slot="title">{{ $t("lang.myTeam") }}</span>
        </el-menu-item>
        <el-menu-item index="/main/team/team-member" class="el-memu-item-new">
          <i class="el-icon-user"></i>
          <span slot="title">{{ $t("lang.teamMembers") }}</span>
        </el-menu-item>
      </el-menu>
    </div>
  </div>
</template>

<script>
export default {
  name: "Team",
  data() {
    return {};
  },
  methods: {},
  mounted() {
    //初始化页面
    this.$router.push({
      path: "/main/team/team-list",
    });
  },
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
  margin: 0;
  padding: 0;
}
.right {
  width: 200px;
}
.el-menu-vertical {
  padding-left: 0px;
  margin-right: 20px;
  border-right: none;
}
.el-memu-item-new {
  height: 40px;
  line-height: 40px;
}
</style>
>
