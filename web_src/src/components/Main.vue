<template>
  <div class="main">
    <!--header-->
    <div class="header">
      <div class="header-body">
        <div class="header-body-center">
          <!--logo-->
          <div class="header-body-logo" style="color: #25b864; font-size: 32px">
            <i class="iconfont icon-edit" style="font-size: 34px"></i>
            <span style="font-size: 24px">{{ productName }}</span>
          </div>
          <!--团队切换-->
          <el-select
            v-model="currentTeam"
            @change="switchTeam"
            size="small"
            @visible-change="workTeamVisible"
            placeholder="团队切换"
          >
            <el-option
              v-for="item in teamList"
              :key="item.id"
              :label="item.name"
              :common="item.id"
              :value="item.id"
            >
            </el-option>
          </el-select>
          <!--导航菜单-->
          <el-menu
            :default-active="activeIndex"
            background-color="#f5f7ee"
            class="header-menu"
            text-color="#595959"
            active-text-color="#25b864"
            mode="horizontal"
            @select="select"
          >
            <el-menu-item class="menu-item" index="todolist">{{
              $t("lang.todo")
            }}</el-menu-item>
            <el-menu-item class="menu-item" index="project">{{
              $t("lang.project")
            }}</el-menu-item>
            <el-menu-item class="menu-item" index="wiki">{{
              $t("lang.wiki")
            }}</el-menu-item>
            <el-menu-item class="menu-item" index="team">{{
              $t("lang.team")
            }}</el-menu-item>
            <!-- <el-menu-item class="menu-item" index="help">{{$t('lang.help')}}</el-menu-item> -->
          </el-menu>
        </div>
        <!--语言切换-->
        <div style="width: 160px">
          <el-switch
            v-model="language"
            active-text="English"
            inactive-text="中文"
          >
          </el-switch>
        </div>
        <!--用户中心-->
        <div>
          <!--我的-->
          <el-dropdown
            placement="bottom-start"
            style="cursor: pointer"
            @command="handleCommand"
          >
            <span style="display: block">
              <i class="iconfont icon-user" style="font-size: 24px"></i>
              <i class="el-icon-arrow-down"></i>
            </span>
            <el-dropdown-menu
              slot="dropdown"
              style="width: auto; text-align: center"
            >
              <el-dropdown-item command="changePasswordDialog">
                <i class="iconfont icon-edit-square" style="font-size: 14px"></i
                >{{ $t("lang.changePassword") }}
              </el-dropdown-item>
              <el-dropdown-item command="signout">
                <i class="iconfont icon-logout" style="font-size: 14px"></i
                >{{ $t("lang.signOut") }}
              </el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
      </div>
    </div>
    <!--main body-->
    <div class="center">
      <router-view></router-view>
    </div>
    <div style="position: fixed; bottom: 60px; right: 100px">
      <el-backtop
        target=".center"
        :visibility-height="200"
        :bottom="40"
        :right="40"
      >
        <el-button size="medium" circle>
          <i
            class="el-icon-caret-top"
            style="font-size: 18px; color: #25b864"
          ></i>
        </el-button>
      </el-backtop>
    </div>
    <!--bottom-->
    <div class="bottom">
      <span>©znj</span>
    </div>
    <!-- 修改密码 -->
    <el-dialog
      :title="$t('lang.changePassword')"
      :close-on-click-modal="changepwd_close_on_click_modal"
      :visible.sync="change_password_dialog"
      width="30%"
    >
      <el-form
        ref="form"
        :model="form"
        label-width="120px"
        label-position="top"
      >
        <el-form-item :label="$t('lang.oldPassword')" size="mini" required>
          <el-input
            v-model="form.old_password"
            type="password"
            placeholder=""
          ></el-input>
        </el-form-item>
        <el-form-item :label="$t('lang.newPassword')" size="mini" required>
          <el-input
            v-model="form.new_password"
            type="password"
            placeholder=""
          ></el-input>
        </el-form-item>
        <el-form-item
          :label="$t('lang.confirmNewPassword')"
          size="mini"
          required
        >
          <el-input
            v-model="form.new_password_confirm"
            type="password"
            placeholder=""
          ></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button size="mini" @click="change_password_dialog = false"
          >取 消</el-button
        >
        <el-button size="mini" type="primary" @click="changePassword"
          >确 定</el-button
        >
      </span>
    </el-dialog>
  </div>
</template>

<script>
import * as productConfig from "../config/product";
export default {
  name: "Main",
  data() {
    return {
      productName: productConfig.name,
      language: false,
      activeIndex: "todolist",
      searchKeyWord: "",
      currentTeam: "",
      teamlist: [],
      change_password_dialog: false,
      changepwd_close_on_click_modal: false,
      form: {
        old_password: "",
        new_password: "",
        new_password_confirm: "",
      },
      teamList: [],
      currentTeam: null,
    };
  },
  watch: {
    $route(to, from) {
      let path = to.path;
      if (path.indexOf("/main/todo-list") != -1) this.activeIndex = "todolist";
      else if (path.indexOf("/main/project") != -1)
        this.activeIndex = "project";
      else if (path.indexOf("/main/wiki") != -1) this.activeIndex = "wiki";
      else if (path.indexOf("/main/team") != -1) this.activeIndex = "team";
      else if (path.indexOf("/main/help") != -1) this.activeIndex = "help";
    },
    language(val) {
      if (val === true) {
        this.$i18n.locale = "en";
        localStorage.setItem("lang", "en");
      } else {
        this.$i18n.locale = "zh";
        localStorage.setItem("lang", "zh");
      }
    },
  },
  methods: {
    // 导航切换
    select: function(index) {
      switch (index) {
        case "todolist":
          this.$router.push({
            path: "/main/todo-list",
          });
          break;
        case "project":
          this.$router.push({
            path: "/main/project-list",
          });
          break;
        case "wiki":
          this.$router.push({
            path: "/main/wiki",
          });
          break;
        case "team":
          this.$router.push({
            path: "/main/team",
          });
          break;
        case "help":
          this.$router.push({
            path: "/main/help",
          });
          break;
      }
    },
    // 我的-菜单
    handleCommand(command) {
      switch (command) {
        // 修改密码
        case "changePasswordDialog":
          this.change_password_dialog = true;
          break;
        // 登出
        case "signout":
          this.Request.signout().then((response) => {
            if (response.data.code == 0) {
              this.$router.push({
                path: "/login",
              });
            }
          });
          break;
      }
    },
    // 切换团队
    switchTeam: function() {
      this.setWorkTeam(this.currentTeam);
    },
    // 修改密码
    changePassword: function() {
      this.Request.changePassword(this.form).then((response) => {
        if (response.data.code == 0) {
          this.$router.push({
            path: "/login",
          });
        } else {
          this.$message(response.data.msg);
        }
      });
    },
    //工作团队-下拉
    workTeamVisible(visible) {
      if (visible) this.getTeamList();
    },
    //获取团队列表
    getTeamList: function() {
      this.Request.teamList()
        .then((response) => {
          if (response.data.code == 0) {
            this.teamList = response.data.data;
            if (this.teamList.length > 0) {
              this.getWorkTeam();
            }
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    // 获取工作团队
    getWorkTeam: function() {
      this.Request.getWorkTeam()
        .then((response) => {
          if (response.data.code == 0) {
            this.currentTeam = response.data.data;
          } else {
            this.setWorkTeam(this.teamList[0].id);
          }
        })
        .catch((err) => {});
    },
    // 设置工作团队
    setWorkTeam(team_id) {
      let data = {
        work_team_id: team_id,
      };
      this.Request.setWorkTeam(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.currentTeam = team_id;
            location.reload();
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
  },
  created() {
    //获取语言设置
    let lang = localStorage.getItem("lang");
    if (lang == "en") {
      this.language = true;
      this.$i18n.locale = "en";
    } else {
      this.language = false;
      localStorage.setItem("lang", "zh");
      this.$i18n.locale = "zh";
    }
  },
  mounted() {
    //获取工作团队
    this.getTeamList();
    //默认为待办页面
    if (this.$route.path == "/main") {
      this.$router.push({
        path: "/main/todo-list",
      });
    }
    let path = this.$route.path;
    if (path.indexOf("/main/todo-list") != -1) this.activeIndex = "todolist";
    else if (path.indexOf("/main/project") != -1) this.activeIndex = "project";
    else if (path.indexOf("/main/wiki") != -1) this.activeIndex = "wiki";
    else if (path.indexOf("/main/team") != -1) this.activeIndex = "team";
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.main {
  width: 100%;
  height: 100%;
  display: -webkit-flex;
  display: flex;
  flex-direction: column;
}
.header {
  height: 60px;
  background-color: #f5f7ee;
}
.header-body {
  height: 100%;
  width: 1200px;
  margin-left: auto;
  margin-right: auto;
  display: flex;
  display: -webkit-flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}
.header-body-logo {
  width: 200px;
  height: 100%;
  display: flex;
  display: -webkit-flex;
  flex-direction: row;
  align-items: center;
}
.header-body-center {
  display: flex;
  display: -webkit-flex;
  flex-direction: row;
  align-items: center;
}
.header-menu {
  padding-left: 20px;
  border: none;
}
.menu-item {
  padding-left: 15px;
  padding-right: 15px;
}
.center {
  flex: 1;
  padding-top: 15px;
  width: 100%;
  overflow-y: auto;
}
.bottom {
  font-size: 12px;
  width: 100%;
  height: 20px;
  text-align: center;
  margin-top: 15px;
}
</style>
