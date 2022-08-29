<template>
  <div class="wiki-list-main">
    <div class="wiki-list-left"></div>
    <div class="wiki-list-center">
      <div>
        <el-button
          style="width: auto; float: left"
          class="el-icon-plus"
          size="mini"
          v-if="wikiType == 1"
          @click="addDoc('1')"
          type="primary"
          >{{ $t("lang.addTeamWiki") }}
        </el-button>
        <el-button
          style="width: auto; float: left"
          class="el-icon-plus"
          size="mini"
          v-if="wikiType == 2"
          @click="addDoc('2')"
          type="primary"
          >{{ $t("lang.addPersonalWiki") }}
        </el-button>
        <el-input
          size="mini"
          style="width: 200px; float: left; margin: 0px 10px"
          :placeholder="$t('lang.wikiKeyword')"
          v-model="searchWikiKeyword"
        >
          <i slot="prefix" class="el-input__icon el-icon-search"></i>
        </el-input>
      </div>
      <div class="wiki-list-center_body">
        <div v-for="doc in docListTemp" :key="doc.id">
          <div @click="openDoc(doc.id, doc.name)" class="wiki-list-doc">
            <el-card class="box-card" shadow="never" :body-style="bodyStyle">
              <div style="flex: 1">
                <i class="iconfont icon-book"></i>
                <div class="doc-title">{{ doc.name }}</div>
              </div>
              <div style="height: 20px" class="docEdit">
                <el-tooltip :content="$t('lang.edit')" placement="top">
                  <a
                    class="icon-btn"
                    style="float: right"
                    @click.stop="editDoc(doc)"
                  >
                    <i class="el-icon-edit" style="font-size: 20px"></i>
                  </a>
                </el-tooltip>
              </div>
            </el-card>
          </div>
        </div>
      </div>
    </div>
    <div class="wiki-list-right">
      <el-menu
        background-color="#F9F9F9"
        default-active="1"
        @select="menuSelect"
        class="el-menu-vertical"
      >
        <el-menu-item index="1" class="el-memu-item">
          <i class="iconfont icon-team" style="font-size: 20px"></i>
          <span slot="title">{{ $t("lang.teamWiki") }}</span>
        </el-menu-item>
        <el-menu-item index="2" class="el-memu-item">
          <i class="iconfont icon-user" style="font-size: 20px"></i>
          <span slot="title">{{ $t("lang.personalWiki") }}</span>
        </el-menu-item>
      </el-menu>
    </div>
    <!-- 文档内容显示 -->
    <div
      v-if="docViewVisible"
      style="
        position: absolute;
        background: #fff;
        height: 100%;
        width: 100%;
        left: 0px;
        top: 0px;
      "
    >
      <DocView
        :docID="docID"
        :docName="docName"
        v-on:docViewHide="docViewHide"
      ></DocView>
    </div>
    <!-- 新增项目 -->
    <el-dialog
      :title="editDocTitle"
      :close-on-click-modal="doc_dialog_close_on_click_modal"
      :visible.sync="edit_doc_dialog_visible"
      width="30%"
    >
      <el-form ref="form" :model="docForm" label-width="auto">
        <el-form-item :label="$t('lang.docName')" size="mini" required>
          <el-input
            v-model="docForm.name"
            :placeholder="$t('lang.docName')"
          ></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button size="mini" @click="edit_doc_dialog_visible = false">{{
          $t("lang.cancel")
        }}</el-button>
        <el-button size="mini" type="primary" @click="saveDoc">{{
          $t("lang.confirm")
        }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import DocView from "../doc/DocView";
export default {
  name: "wikiList",
  components: { DocView },
  data() {
    return {
      wikiType: 1,
      editDocTitle: "新增文档",
      doc_dialog_close_on_click_modal: false,
      edit_doc_dialog_visible: false,
      searchWikiKeyword: "",
      bodyStyle:
        "display: -webkit-flex;display: flex;flex-direction: column;padding:0 0;height:100%;",
      docList: [],
      docListTemp: [],
      docForm: {
        id: "",
        name: "",
        team_id: "1",
        user_id: "",
      },
      docViewVisible: false,
      docID: "",
      docName: "",
    };
  },
  watch: {
    docList(val) {
      this.docListTemp = val;
    },
    searchWikiKeyword(val) {
      if (val) {
        this.docListTemp = this.docList.filter((doc) => {
          return doc.name.indexOf(val) != -1;
        });
      } else {
        this.docListTemp = this.docList;
      }
    },
  },
  methods: {
    // 获取项目列表
    getDocList: function (param) {
      this.Request.docList(param)
        .then((response) => {
          if (response.data.code == 0) {
            this.docList = response.data.data;
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    // 知识库类型选择
    menuSelect(index) {
      this.searchWikiKeyword = "";
      this.wikiType = index;
      switch (this.wikiType) {
        case "1":
          this.docForm.user_id = "";
          this.docForm.team_id = "1";
          break;
        case "2":
          this.docForm.user_id = "1";
          this.docForm.team_id = "";
          break;
        default:
          break;
      }
      this.getDocList(this.docForm);
    },
    // 新增文档
    addDoc(type) {
      this.resetForm();
      this.edit_doc_dialog_visible = true;
      this.editDocTitle =
        this.wikiType == "1"
          ? this.$t("lang.addTeamWiki")
          : this.$t("lang.addPersonalWiki");
      switch (type) {
        case "1":
          this.docForm.user_id = "";
          this.docForm.team_id = "1";
          break;
        case "2":
          this.docForm.user_id = "1";
          this.docForm.team_id = "";
          break;
        default:
          break;
      }
    },
    // 编辑文档
    editDoc(doc) {
      this.edit_doc_dialog_visible = true;
      this.editDocTitle =
        this.wikiType == "1"
          ? this.$t("lang.editTeamWiki")
          : this.$t("lang.editPersonalWiki");
      this.docForm.id = doc.id;
      this.docForm.name = doc.name;
    },
    // 重置form
    resetForm() {
      this.docForm = {
        id: "",
        name: "",
        team_id: "1",
        user_id: "",
      };
    },
    // 保存文档
    saveDoc() {
      this.Request.docSave(this.docForm)
        .then((response) => {
          if (response.data.code == 0) {
            this.edit_doc_dialog_visible = false;
            this.getDocList(this.docForm);
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    //查看文档内容
    openDoc(docID, docName) {
      this.docViewVisible = true;
      this.docID = docID;
      this.docName = docName;
    },
    // 关闭文档显示
    docViewHide() {
      this.docViewVisible = false;
    },
  },
  mounted() {
    //加载文档列表
    this.getDocList(this.docForm); //默认团队知识库
  },
};
</script>

<style scoped>
.wiki-list-main {
  width: 1200px;
  margin-left: auto;
  margin-right: auto;
  display: -webkit-flex;
  display: flex;
  flex-direction: row;
}
.wiki-list-left {
  width: 200px;
}
.wiki-list-right {
  width: 250px;
}
.wiki-list-center {
  flex: 1;
  display: -webkit-flex;
  display: flex;
  flex-direction: column;
}
.wiki-list-center_body {
  flex: 1;
  margin: 10px 0px;
  display: -webkit-flex;
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  align-content: flex-start;
  justify-content: flex-start;
}
.wiki-list-doc {
  text-align: center;
  margin: 0px 10px 0px 0px;
  cursor: pointer;
}
.box-card {
  width: 70px;
  height: 100px;
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
.el-memu-item {
  height: 40px;
  line-height: 40px;
}
.icon-btn {
  display: none;
}
.box-card:hover a {
  display: inline-block;
}
.doc-title {
  font-size: 12px;
  color: grey;
  word-break: break-all;
}
.el-dialog {
  display: flex;
  display: -webkit-flex;
  flex-direction: column;
}
.el-dialog__body {
  flex: 1;
}
</style>
