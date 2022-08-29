<template>
  <div class="main">
    <div class="left"></div>
    <div class="center">
      <div class="center_body_header">
        <el-button
          size="mini"
          style="width: 120px; float: right; margin: 5px 0px"
          class="el-icon-plus"
          @click="addDoc"
          type="primary"
          >{{ $t("lang.addDoc") }}
        </el-button>
        <el-input
          size="mini"
          style="width: 200px; float: right; margin: 5px"
          :placeholder="$t('lang.docKeyword')"
          v-model="searchDocKeyword"
        >
          <i slot="prefix" class="el-input__icon el-icon-search"></i>
        </el-input>
      </div>
      <div class="center_body">
        <div v-for="doc in docListTemp" :key="doc.id">
          <div @click="openDoc(doc.id, doc.name)" class="doc">
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
    <div class="right"></div>
    <!-- 文档内容显示 -->
    <div
      v-if="docViewVisible"
      style="
        position: absolute;
        background: #fff;
        height: 100%;
        width: 100%;
        left: 0;
        top: 0;
      "
    >
      <DocView
        :docID="docID"
        :docName="docName"
        v-on:docViewHide="docViewHide"
      ></DocView>
    </div>
    <!-- 添加或编辑文档 -->
    <el-dialog
      :title="add ? $t('lang.addDoc') : $t('lang.editDoc')"
      :close-on-click-modal="add_doc_close_on_click_modal"
      :visible.sync="add_doc_dialog_visible"
      width="30%"
    >
      <el-form ref="form" :model="docForm" label-width="120px">
        <el-form-item :label="$t('lang.docName')" size="mini" required>
          <el-input
            v-model="docForm.name"
            :placeholder="$t('lang.docName')"
          ></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button size="mini" @click="add_doc_dialog_visible = false">{{
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
  name: "ProjectDoc",
  components: { DocView },
  data() {
    return {
      projectID: "",
      docID: "",
      docName: "",
      docViewVisible: false,
      docEditVisible: false,
      searchDocKeyword: "",
      bodyStyle:
        "display: -webkit-flex;display: flex;flex-direction: column;padding:0 0;height:100%;",
      docList: [],
      docListTemp: [],
      add: true,
      add_doc_close_on_click_modal: false,
      add_doc_dialog_visible: false,
      docForm: {
        id: "",
        name: "",
        project_id: "",
      },
    };
  },
  watch: {
    docList(val) {
      this.docListTemp = val;
    },
    searchDocKeyword(val) {
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
    // 获取文档列表
    getDocList() {
      let data = {
        project_id: this.projectID,
      };
      this.Request.docList(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.docList = response.data.data;
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    // 打开文档
    openDoc(docID, docName) {
      this.docViewVisible = true;
      this.docID = docID;
      this.docName = docName;
    },
    // 关闭文档显示
    docViewHide() {
      this.docViewVisible = false;
    },
    // 新增文档
    addDoc() {
      this.add = true;
      this.add_doc_dialog_visible = true;
      this.docForm.id = "";
      this.docForm.name = "";
      this.docForm.project_id = this.projectID;
    },
    // 编辑文档
    editDoc(doc) {
      this.add = false;
      this.add_doc_dialog_visible = true;
      this.docForm.id = doc.id;
      this.docForm.name = doc.name;
      this.docForm.project_id = this.projectID;
    },
    // 保存文档
    saveDoc() {
      this.Request.docSave(this.docForm)
        .then((response) => {
          if (response.data.code == 0) {
            this.add_doc_dialog_visible = false;
            this.getDocList();
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
    //获取文档列表
    this.getDocList();
    //打开指定的文档
    let docID = sessionStorage.getItem("docID");
    let docName = sessionStorage.getItem("docName");
    if (docID) {
      this.openDoc(docID, docName);
      sessionStorage.removeItem("docID");
      sessionStorage.removeItem("docName");
    }
  },
};
</script>

<style scoped>
.main {
  width: 1200px;
  margin-left: auto;
  margin-right: auto;
  display: flex;
  display: -webkit-flex;
  flex-direction: row;
}
.left {
  width: 200px;
}
.center {
  flex: 1;
  margin-top: 5px;
  display: -webkit-flex;
  display: flex;
  flex-direction: column;
}
.center_body {
  flex: 1;
  margin: 10px 0px;
  display: flex;
  display: -webkit-flex;
  flex-direction: row;
  flex-wrap: wrap;
  align-content: flex-start;
  justify-content: flex-start;
}
.doc {
  text-align: center;
  margin: 0px 10px 0px 0px;
  cursor: pointer;
}
.right {
  width: 200px;
}
.box-card {
  width: 70px;
  height: 100px;
  margin-bottom: 12px;
  margin-right: 12px;
  padding: 10px 10px;
  cursor: pointer;
}
.docEdit {
  display: none;
}
.box-card:hover .docEdit {
  display: block;
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
