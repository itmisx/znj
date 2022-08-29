<template>
  <div class="main">
    <div style="margin: 0px 0px 10px 0px">
      <el-button
        type="primary"
        v-if="teamAddButtonShow"
        class="el-icon-plus"
        @click="showTeamAddForm"
        size="mini"
        >{{ $t("lang.addTeam") }}
      </el-button>
      <el-input
        size="small"
        :placeholder="$t('lang.inputTeamName')"
        v-model="team.name"
        @keyup.enter.native="teamAddSave"
        v-if="teamAddFormShow"
      >
        <div slot="append">
          <el-button @click="teamAddSave">{{ $t("lang.confirm") }}</el-button>
          <el-divider direction="vertical"></el-divider>
          <el-button @click="teamAddCancel">{{ $t("lang.cancel") }}</el-button>
        </div>
      </el-input>
    </div>
    <el-table
      ref="singleTable"
      :data="teamList"
      highlight-current-row
      style="width: 100%"
    >
      <el-table-column type="index" width="50"> </el-table-column>
      <el-table-column property="name" :label="$t('lang.teamName')">
        <template slot-scope="scope">
          <el-input
            size="small"
            v-model="scope.row.name"
            :placeholder="$t('lang.inputTeamName')"
            @keyup.enter.native="teamEditSave(scope.row)"
            :style="{ display: scope.$index == scopeIndex ? 'block' : 'none' }"
          ></el-input>
          <span
            :style="{ display: scope.$index != scopeIndex ? 'block' : 'none' }"
            >{{ scope.row.name }}</span
          >
        </template>
      </el-table-column>
      <el-table-column :label="$t('lang.action')" align="center" width="200">
        <template slot-scope="scope">
          <el-button
            @click="teamEdit(scope.$index)"
            :style="{
              display: scope.$index != scopeIndex ? 'inline-block' : 'none',
            }"
            type="text"
            size="small"
            >{{ $t("lang.edit") }}
          </el-button>
          <el-button
            @click="teamEditSave(scope.row)"
            :style="{
              display: scope.$index == scopeIndex ? 'inline-block' : 'none',
            }"
            type="text"
            size="small"
            >{{ $t("lang.save") }}
          </el-button>
          <el-button
            @click="teamEditCancel()"
            :style="{
              display: scope.$index == scopeIndex ? 'inline-block' : 'none',
            }"
            type="text"
            size="small"
            >{{ $t("lang.cancel") }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "TeamList",
  data() {
    return {
      teamAddButtonShow: true,
      teamAddFormShow: false,
      teamList: [],
      currentRow: {}, //指示当前要编辑的行
      scopeIndex: null,
      team: {},
    };
  },
  methods: {
    //获取团队列表
    loadTeamList: function () {
      this.Request.teamList()
        .then((response) => {
          if (response.data.code == 0) {
            this.teamList = response.data.data;
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    // 显示团队新增表单
    showTeamAddForm() {
      this.scopeIndex = null;
      this.teamAddFormShow = true;
      this.teamAddButtonShow = false;
    },
    // 新增团队取消
    teamAddCancel() {
      this.teamAddFormShow = false;
      this.teamAddButtonShow = true;
    },
    // 添加保存
    teamAddSave() {
      let data = {
        name: this.team.name,
      };
      this.Request.teamSave(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.teamAddCancel();
            this.loadTeamList();
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    // 编辑保存
    teamEditSave(row) {
      let data = {
        id: row.id,
        name: row.name,
      };
      this.Request.teamSave(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.scopeIndex = null;
          } else {
            this.$message(response.data.msg);
          }
          this.scopeIndex = null;
        })
        .catch((err) => {
          this.scopeIndex = null;
        });
    },
    // 团队编辑
    teamEdit(index) {
      this.scopeIndex = index;
    },
    // 编辑取消
    teamEditCancel() {
      this.scopeIndex = null;
    },
  },
  mounted() {
    //获取团队列表
    this.loadTeamList();
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
