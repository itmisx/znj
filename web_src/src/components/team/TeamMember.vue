<template>
  <div>
    <div style="margin: 0px 0px 5px 0px">
      <el-button
        type="primary"
        class="el-icon-plus"
        style="margin: 0px 0px"
        @click="showAddMemberDialog"
        size="mini"
        >{{ $t("lang.addMember") }}
      </el-button>
      <!--团队切换-->
      <el-select
        v-model="currentTeam"
        size="mini"
        @change="teamChange"
        :placeholder="$t('lang.selectTeam')"
      >
        <el-option
          v-for="item in teamList"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        >
        </el-option>
      </el-select>
    </div>
    <el-table
      ref="singleTable"
      :data="memberList"
      highlight-current-row
      style="width: 100%"
    >
      <el-table-column type="index" width="50"> </el-table-column>
      <el-table-column property="user" width="150" :label="$t('lang.account')">
      </el-table-column>
      <el-table-column property="name" width="150" :label="$t('lang.name')">
      </el-table-column>
      <el-table-column property="role" width="150" :label="$t('lang.role')">
        <template slot-scope="scope">
          <span style="margin-left: 10px" v-if="scope.row.role == 1">{{
            $t("lang.superAdminRole")
          }}</span>
          <span style="margin-left: 10px" v-if="scope.row.role == 2">{{
            $t("lang.adminRole")
          }}</span>
          <span style="margin-left: 10px" v-if="scope.row.role == 3">{{
            $t("lang.userRole")
          }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('lang.action')">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="text"
            v-if="
              scope.row.role != 1 &&
                (currentMemberRole == 1 || currentMemberRole == 2)
            "
            @click="resetPassword(scope.row)"
            >{{ $t("lang.resetPassword") }}
          </el-button>
          <el-button
            size="mini"
            type="text"
            v-if="scope.row.role == 3 && currentMemberRole == 1"
            @click="setAdminRole(scope.row)"
            >{{ $t("lang.setAdminRole") }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 添加成员 -->
    <el-dialog
      :title="add ? $t('lang.addMember') : $t('lang.editMember')"
      :close-on-click-modal="changepwd_close_on_click_modal"
      :visible.sync="add_team_member_dialog"
      width="500px"
    >
      <el-form ref="form" :model="addMemberForm" label-width="auto">
        <el-form-item :label="$t('lang.account')" size="mini" required>
          <el-input v-model="addMemberForm.user" placeholder=""></el-input>
        </el-form-item>
        <el-form-item :label="$t('lang.name')" size="mini" required>
          <el-input v-model="addMemberForm.name" placeholder=""></el-input>
        </el-form-item>
        <el-form-item :label="$t('lang.password')" size="mini" required>
          <el-input
            v-model="addMemberForm.password"
            type="password"
            :placeholder="$t('lang.passwordRule')"
          >
          </el-input>
        </el-form-item>
        <el-form-item :label="$t('lang.confirmPassword')" size="mini" required>
          <el-input
            v-model="addMemberForm.password_confirm"
            type="password"
            :placeholder="$t('lang.passwordRule')"
          >
          </el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button size="mini" @click="add_team_member_dialog = false">{{
          $t("lang.cancel")
        }}</el-button>
        <el-button size="mini" type="primary" @click="addTeamMember">{{
          $t("lang.confirm")
        }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "TeamMember",
  data() {
    return {
      add: true,
      changepwd_close_on_click_modal: false,
      add_team_member_dialog: false,
      currentTeam: null,
      currentMemberRole: 3,
      memberList: [],
      teamList: [],
      addMemberForm: {
        name: "",
        user: "",
        password: "",
        password_confirm: "",
      },
    };
  },
  methods: {
    //显示添加成员对话框
    showAddMemberDialog() {
      this.add_team_member_dialog = true;
      this.resetAddMemberForm();
    },
    //获取团队列表
    getTeamList: function() {
      this.Request.teamList()
        .then((response) => {
          if (response.data.code == 0) {
            this.teamList = response.data.data;
            if (this.teamList.length > 0) {
              this.currentTeam = this.teamList[0].id;
              this.getMemberList();
            }
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    //获取团队成员
    getMemberList: function() {
      if (!this.currentTeam) return;
      let data = {
        team_id: this.currentTeam,
      };
      this.Request.teamMemberList(data)
        .then((response) => {
          if (response.data.code == 0) {
            let data = response.data.data;
            this.memberList = data.list;
            if (data.user_role) this.currentMemberRole = data.user_role;
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    //添加团队成员
    addTeamMember: function() {
      if (!this.currentTeam) return;
      let data = {
        team_id: this.currentTeam,
        name: this.addMemberForm.name,
        user: this.addMemberForm.user,
        password: this.addMemberForm.password,
        password_confirm: this.addMemberForm.password_confirm,
      };
      this.Request.addTeamMember(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.getMemberList();
            this.add_team_member_dialog = false;
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {
          this.add_team_member_dialog = false;
        });
    },
    //选择团队
    teamChange(val) {
      this.getMemberList();
    },
    //重置密码
    resetPassword(user) {
      this.$prompt("请输入密码", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
      })
        .then(({ value }) => {
          let data = {
            id: user.id,
            team_id: this.currentTeam,
            password: value,
          };
          this.Request.resetPassword(data)
            .then((response) => {
              if (response.data.code == 0) {
                this.getMemberList();
              } else {
                this.$message(response.data.msg);
              }
            })
            .catch((err) => {});
        })
        .catch(() => {});
    },
    //设为管理员
    setAdminRole(user) {
      let data = {
        id: user.id,
        team_id: this.currentTeam,
      };
      this.Request.setAdminRole(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.getMemberList();
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    //重置表单
    resetAddMemberForm() {
      this.addMemberForm = {
        name: "",
        user: "",
        password: "",
        password_confirm: "",
      };
    },
  },
  mounted() {
    //加载团队列表
    this.getTeamList();
  },
};
</script>

<style scoped>
.main {
  height: 100%;
  width: 100%;
  display: flex;
  display: -webkit-flex;
  flex-direction: column;
}
</style>
