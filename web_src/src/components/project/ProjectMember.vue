<template>
  <div class="main">
    <div class="top-bar">
      <!-- 成员设置 -->
      <el-button
        size="mini"
        style="width: auto; float: right; margin: 5px 0px"
        @click="memberSetDialogVisible = true"
        icon="el-icon-user"
        type="primary"
        >{{ $t("lang.projectMembersSet") }}
      </el-button>
      <!-- 成员搜索 -->
      <el-input
        size="mini"
        style="width: 200px; float: right; margin: 5px"
        :placeholder="$t('lang.memberKeyword')"
        v-model="searchMemberKeyword"
      >
        <i slot="prefix" class="el-input__icon el-icon-search"></i>
      </el-input>
    </div>
    <div class="body">
      <div
        v-for="item in projectMembersTemp"
        style="margin: 20px; width: 90px"
        :key="item.id"
      >
        <i class="el-icon-user"></i>
        <span style="font-size: 14px">{{ item.name }}</span>
      </div>
    </div>
    <!-- 设置项目成员 -->
    <el-dialog
      :title="$t('lang.projectMembersSet')"
      :close-on-click-modal="closeOnClickModal"
      :visible.sync="memberSetDialogVisible"
      width="30%"
    >
      <el-checkbox
        :indeterminate="isIndeterminate"
        v-model="checkAll"
        @change="handleCheckAllChange"
        >{{ $t("lang.selectAll") }}</el-checkbox
      >
      <div style="margin: 15px 0"></div>
      <el-checkbox-group v-model="checkedItems" @change="handleCheckedMembers">
        <el-checkbox
          v-for="(item, index) in totalItems"
          :label="item"
          :key="item"
          >{{ teamMembers[index].name }}
        </el-checkbox>
      </el-checkbox-group>
      <span slot="footer" class="dialog-footer">
        <el-button size="mini" @click="memberSetDialogVisible = false">{{
          $t("lang.cancel")
        }}</el-button>
        <el-button size="mini" type="primary" @click="memberSet">{{
          $t("lang.confirm")
        }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "ProjectMember",
  data() {
    return {
      projectID: 0,
      searchMemberKeyword: "",
      memberSetDialogVisible: false,
      closeOnClickModal: false,
      teamMembers: [],
      projectMembers: [],
      projectMembersTemp: [],
      totalItems: [],
      checkedItems: [],
      isIndeterminate: true,
      checkAll: false,
    };
  },
  watch: {
    //获取项目成员
    projectMembers(val) {
      this.projectMembersTemp = val;
    },
    //成员搜索
    searchMemberKeyword(val) {
      if (val) {
        this.projectMembersTemp = this.projectMembers.filter((member) => {
          return member.name.indexOf(val) != -1;
        });
      } else {
        this.projectMembersTemp = this.projectMembers;
      }
    },
  },
  methods: {
    //全选
    handleCheckAllChange(val) {
      this.checkedItems = val ? this.totalItems : [];
      this.isIndeterminate = false;
    },
    //选择设置
    handleCheckedMembers(value) {
      let checkedCount = value.length;
      this.checkAll = checkedCount === this.totalItems.length;
      this.isIndeterminate =
        checkedCount > 0 && checkedCount < this.totalItems.length;
    },
    //保存设置
    memberSet() {
      let ids = "";
      let checkedItems = this.checkedItems;
      for (let i in checkedItems) {
        if (ids != "") ids = ids + "," + checkedItems[i].toString();
        else ids = checkedItems[i].toString();
      }
      //提交checkedMembers
      let data = {
        project_id: this.projectID,
        ids: ids,
      };
      this.Request.projectMemberSet(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.getMemberList();
            this.memberSetDialogVisible = false;
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    //获取项目成员
    getMemberList() {
      let data = {
        project_id: this.projectID,
      };
      this.Request.projectMemberList(data)
        .then((response) => {
          if (response.data.code == 0) {
            // 这个地方必须是数值的集合，非二维
            let list = response.data.data;
            this.projectMembers = list;
            this.checkedItems = [];
            for (var i in list) {
              this.checkedItems.push(list[i].id);
            }
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    //获取团队成员
    getTeamMember() {
      let data = {
        project_id: this.projectID,
      };
      this.Request.teamMemberList(data)
        .then((response) => {
          if (response.data.code == 0) {
            let list = response.data.data.list;
            this.teamMembers = list;
            for (var i in list) {
              this.totalItems.push(list[i].id);
            }
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
  },
  mounted() {
    //设置工作项目
    this.projectID = sessionStorage.getItem("projectID");
    //加载项目成员
    this.getMemberList();
    //加载团队成员
    this.getTeamMember();
  },
};
</script>

<style scoped>
.main {
  width: 800px;
  height: 100%;
  margin: 0px auto;
}
.top-bar {
  margin: 5px;
}
.body {
  display: flex;
  display: -webkit-flex;
  flex-direction: row;
  flex-wrap: wrap;
  align-content: flex-start;
  justify-content: flex-start;
}
</style>
