<template>
  <div class="main">
    <div class="left"></div>
    <div class="center">
      <div v-if="!showLatestProject">
        <el-button
          style="width: 120px; float: left"
          class="el-icon-plus"
          size="mini"
          @click="showAddProjectDialog"
          type="primary"
          >{{ $t("lang.addProject") }}
        </el-button>
        <el-input
          size="mini"
          style="width: 200px; float: left; margin: 0px 10px"
          :placeholder="$t('lang.projectKeyword')"
          v-model="searchProjectKeyword"
        >
          <i slot="prefix" class="el-input__icon el-icon-search"></i>
        </el-input>
      </div>
      <div>
        <!-- 全部项目 -->
        <div v-if="!showLatestProject" class="center_body">
          <div
            v-for="project in projectListTemp"
            :key="project.id"
            @click="openProject(project.id, project.name)"
          >
            <el-card class="box-card" shadow="never" :body-style="bodyStyle">
              <div style="font-size: 16px; color: #595959">
                <div
                  style="
                    width: 160px;
                    text-overflow: ellipsis;
                    overflow: hidden;
                    white-space: nowrap;
                    display: inline-block;
                  "
                >
                  {{ project.name }}
                </div>
                <el-tooltip :content="$t('lang.edit')" placement="top">
                  <a
                    class="icon-btn"
                    style="float: right"
                    @click.stop="showEditProjectDialog(project)"
                  >
                    <i class="el-icon-edit" style="font-size: 20px"></i>
                  </a>
                </el-tooltip>
                <el-tooltip :content="$t('lang.top')" placement="top">
                  <a
                    class="icon-btn"
                    @click.stop="starProject(project)"
                    style="float: right; margin-right: 10px"
                  >
                    <i class="el-icon-star-off" style="font-size: 20px"></i>
                  </a>
                </el-tooltip>
              </div>
              <div style="font-size: 12px; color: #898989; margin-top: 5px">
                <div style="word-wrap: break-word">{{ project.desc }}</div>
              </div>
            </el-card>
          </div>
        </div>
        <!-- 最近项目 -->
        <div v-if="showLatestProject" class="center_body">
          <div
            v-for="project in latestProjects"
            :key="project.id"
            @click="openProject(project.id, project.name)"
          >
            <el-card class="box-card" shadow="never" :body-style="bodyStyle">
              <div style="font-size: 16px; color: #595959">
                <div
                  style="
                    width: 160px;
                    text-overflow: ellipsis;
                    overflow: hidden;
                    white-space: nowrap;
                    display: inline-block;
                  "
                >
                  {{ project.name }}
                </div>
                <el-tooltip :content="$t('lang.edit')" placement="top">
                  <a
                    class="icon-btn"
                    style="float: right"
                    @click.stop="showEditProjectDialog(project)"
                  >
                    <i class="el-icon-edit" style="font-size: 20px"></i>
                  </a>
                </el-tooltip>
                <el-tooltip :content="$t('lang.top')" placement="top">
                  <a
                    class="icon-btn"
                    @click.stop="starProject(project)"
                    style="float: right; margin-right: 10px"
                  >
                    <i class="el-icon-star-off" style="font-size: 20px"></i>
                  </a>
                </el-tooltip>
              </div>
              <div style="font-size: 12px; color: #898989; margin-top: 5px">
                <div style="word-wrap: break-word">{{ project.desc }}</div>
              </div>
            </el-card>
          </div>
        </div>
      </div>
    </div>
    <div class="right">
      <el-menu
        background-color="#F9F9F9"
        default-active="1"
        @select="menuSelect"
        class="el-menu-vertical"
      >
        <el-menu-item index="1" class="el-memu-item-new">
          <i class="el-icon-s-grid"></i>
          <span slot="title">{{ $t("lang.allProjects") }}</span>
        </el-menu-item>
        <el-menu-item index="2" class="el-memu-item-new">
          <i class="el-icon-time"></i>
          <span slot="title">{{ $t("lang.latestProjects") }}</span>
        </el-menu-item>
      </el-menu>
      <img
        src="@/assets/img/p2.jpg"
        style="margin-top: 30px; width: 50%; margin-left: 25px"
      />
    </div>
    <!-- 新增项目 -->
    <el-dialog
      :title="add ? $t('lang.addProject') : $t('lang.editProject')"
      :close-on-click-modal="add_project_close_on_click_modal"
      :visible.sync="add_project_dialog_visible"
      width="30%"
    >
      <el-form ref="form" :model="projectForm" label-width="120px">
        <el-form-item :label="$t('lang.projectName')" size="mini">
          <el-input
            v-model="projectForm.name"
            :placeholder="$t('lang.pleaseInputProjectName')"
          ></el-input>
        </el-form-item>
        <el-form-item :label="$t('lang.projectDesc')" size="mini">
          <el-input
            v-model="projectForm.desc"
            :placeholder="$t('lang.pleaseInputProjectDesc')"
          ></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button size="mini" @click="add_project_dialog_visible = false">{{
          $t("lang.cancel")
        }}</el-button>
        <el-button size="mini" type="primary" @click="saveProject">{{
          $t("lang.confirm")
        }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "ProjectList",
  data() {
    return {
      add: true,
      add_project_close_on_click_modal: false,
      add_project_dialog_visible: false,
      searchProjectKeyword: "",
      latestProjects: [],
      bodyStyle:
        "display: -webkit-flex;display: flex;flex:1;flex-direction: column;padding:5px 0;",
      showLatestProject: false,
      projectList: [],
      projectListTemp: [],
      projectForm: {
        id: "",
        name: "",
        desc: ""
      }
    };
  },
  watch: {
    projectList(val) {
      this.projectListTemp = val;
    },
    searchProjectKeyword(val) {
      if (val) {
        this.projectListTemp = this.projectList.filter(project => {
          return project.name.indexOf(val) != -1;
        });
      } else {
        this.projectListTemp = this.projectList;
      }
    }
  },
  methods: {
    // reset ProjectForm
    resetProjectForm() {
      this.id = "";
      this.projectForm.name = "";
      this.projectForm.desc = "";
    },
    // 显示新增项目
    showAddProjectDialog: function() {
      this.add = true;
      this.resetProjectForm();
      this.add_project_dialog_visible = true;
    },
    // 新增项目
    saveProject: function() {
      this.Request.projectSave(this.projectForm)
        .then(response => {
          if (response.data.code == 0) {
            this.getProjectList();
            this.add_project_dialog_visible = false;
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch(err => {
          this.showAddProjectDialog = false;
        });
    },
    // 显示编辑项目
    showEditProjectDialog: function(project) {
      this.add = false;
      this.projectForm.id = project.id;
      this.projectForm.name = project.name;
      this.projectForm.desc = project.desc;
      this.add_project_dialog_visible = true;
    },
    // 打开项目页面
    openProject: function(projectID, projectName) {
      sessionStorage.setItem("projectID", projectID);
      sessionStorage.setItem("projectName", projectName);
      this.$router.push({
        path: "/main/project"
      });
    },
    // 获取项目列表
    getProjectList: function() {
      this.Request.projectList()
        .then(response => {
          if (response.data.code == 0) {
            this.projectList = response.data.data;
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch(err => {});
    },
    // 项目置顶
    starProject(project) {
      let data = project;
      data.star_time = 1;
      this.Request.projectSave(data)
        .then(response => {
          if (response.data.code == 0) {
            this.getProjectList();
            this.add_project_dialog_visible = false;
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch(err => {
          this.showAddProjectDialog = false;
        });
    },
    menuSelect(index) {
      switch (index) {
        case "1":
          this.showLatestProject = false;
          this.getProjectList();
          break;
        case "2":
          this.showLatestProject = true;
          this.getLatestProjectsAndChapters();
          break;
        default:
          break;
      }
    },
    // 获取最近访问的项目和章节
    getLatestProjectsAndChapters() {
      this.Request.getLatestProjectsAndChapters().then(response => {
        if (response.data.code == 0) {
          this.latestProjects = response.data.data.projects
            ? response.data.data.projects
            : [];
        } else {
          this.$message(response.data.msg);
        }
      });
    }
  },
  created() {},
  mounted() {
    //加载项目列表
    this.getProjectList();
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
.right {
  width: 250px;
}
.center {
  flex: 1;
  display: -webkit-flex;
  display: flex;
  flex-direction: column;
}
.center_body {
  flex: 1;
  padding-top: 10px;
  padding-left: 0px;
  display: -webkit-flex;
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  align-content: flex-start;
  justify-content: flex-start;
}
.box-card {
  width: 215px;
  height: 90px;
  margin-bottom: 12px;
  margin-right: 12px;
  padding: 10px 10px;
  cursor: pointer;
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
.icon-btn {
  display: none;
}
.box-card:hover a {
  display: inline-block;
}
</style>
