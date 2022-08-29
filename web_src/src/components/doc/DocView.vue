<template>
  <div class="doc-view-main">
    <div class="doc-view-header">
      <div>
        <el-button
          size="mini"
          type="primary"
          style="margin-left: 5px"
          @click="back"
          icon="el-icon-arrow-left"
          >{{ $t("lang.goBack") }}
        </el-button>
      </div>
      <div style="font-size: 20px; font-weight: bold; padding: 10px">
        {{ docName }}
      </div>
    </div>
    <div class="doc-view-body">
      <!-- 章节列表 -->
      <div class="left">
        <el-input
          size="mini"
          style="width: 96%; margin: 5px"
          :placeholder="$t('lang.searchChapter')"
          v-model="searchChapterKeyword"
        >
          <i slot="prefix" class="el-input__icon el-icon-search"></i>
        </el-input>
        <el-tree
          node-key="id"
          ref="tree"
          draggable
          @node-drag-end="handleChapterDragEnd"
          :allow-drag="allowChapterDrag"
          :allow-drop="allowChapterDrop"
          :accordion="true"
          :default-expand-all="false"
          :default-expanded-keys="['0']"
          :highlight-current="true"
          :expand-on-click-node="false"
          :data="chapterList"
          :props="treeProps"
          v-on:node-click="getChapterContent"
          :filter-node-method="filterChapter"
        >
          <span class="custom-tree-node" slot-scope="{ node, data }">
            <span>{{ node.label }}</span>
            <span class="tree-button">
              <el-button
                type="text"
                size="mini"
                @click.stop="() => addChapter(data)"
                >{{ $t("lang.add") }}
              </el-button>
              <el-button
                type="text"
                size="mini"
                v-if="data.id != '0'"
                @click.stop="() => editChapter(data, node)"
                >{{ $t("lang.edit") }}
              </el-button>
            </span>
          </span>
        </el-tree>
      </div>
      <div class="center-left"></div>
      <div class="center-center" id="editor_preview">
        <div class="center-center-header">{{ currentChapter.name }}</div>
        <div
          class="center-center-body"
          :style="centerCenterBodyStyle"
          style="z-index: 200"
        >
          <!-- markdown显示 -->
          <div
            :style="{
              display: currentChapter.doc_type == 1 ? 'block' : 'none',
            }"
            style="width: 100%; height: 100%"
          >
            <div :id="vditorPreviewID" style="height: 98%"></div>
          </div>
          <!-- 流程图显示 -->
          <div
            :style="{
              display: currentChapter.doc_type == 2 ? 'flex' : 'none',
              display: currentChapter.doc_type == 2 ? '-webkit-flex' : 'none',
            }"
            style="
              width: 100%;
              height: 98%;
              overflow: hidden;
              border: 1px solid #d1d5da;
              flex-direction: column;
              align-items: stretch;
            "
            class="geEditor"
            @contextmenu.prevent
          >
            <div
              style="
                height: 35px;
                background: #f4f7f9;
                border-bottom: 1px solid #d1d5da;
                line-height: 35px;
                vertical-align: middle;
              "
            >
              <el-tooltip
                class="item"
                effect="dark"
                content="可右键拖拽移动"
                placement="top"
              >
                <i
                  class="iconfont icon-info-circle graphTool"
                  style="font-size: 20px; margin-left: 10px"
                ></i>
              </el-tooltip>
              <el-tooltip
                class="item"
                effect="dark"
                content="全屏预览"
                placement="top"
                v-if="!graphFullScreen"
              >
                <i
                  class="iconfont icon-fullscreen graphTool"
                  @click="graphFullScreenChange"
                  style="font-size: 20px"
                ></i>
              </el-tooltip>
              <el-tooltip
                class="item"
                effect="dark"
                content="退出全屏"
                placement="top"
                v-if="graphFullScreen"
              >
                <i
                  class="iconfont icon-fullscreen-exit graphTool"
                  @click="graphFullScreenChange"
                  style="font-size: 20px"
                ></i>
              </el-tooltip>
            </div>
            <div
              :id="graphPreviewID"
              style="flex: 1; width: 100%"
              :style="{ cursor: graphCursor }"
              @mousedown="graphPreviewMoveStart"
              @mouseup="graphPreviewMoveEnd"
            ></div>
          </div>
        </div>
      </div>
      <!-- 章节内容操作-编辑或查看历史版本 -->
      <div class="center-right">
        <el-button
          type="primary"
          icon="el-icon-edit"
          circle
          style="margin: 10px 10px 0px 5px"
          @click="editChapterContent"
          v-if="currentChapter.id != ''"
        >
        </el-button>
        <el-tooltip
          effect="dark"
          content="查看历史版本"
          placement="right"
          v-if="currentChapter.id != ''"
        >
          <el-button
            type="info"
            icon="iconfont icon-time-circle"
            circle
            plain
            style="margin: 10px 10px 0px 10px"
            size="mini"
            @click="showChapterContentList"
          >
          </el-button>
        </el-tooltip>
      </div>
    </div>
    <!-- 编辑章节 -->
    <el-dialog
      :title="addChapterDialog ? $t('lang.addChapter') : $t('lang.editChapter')"
      :close-on-click-modal="closeOnClickModal"
      :visible.sync="editChapterDialogVisible"
      width="350px"
    >
      <el-form
        ref="chapterForm"
        label-width="80px"
        :rules="rules"
        :model="chapter"
      >
        <el-form-item label="名称" size="mini" prop="name">
          <el-input
            v-model="chapter.name"
            style="width: 200px"
            placeholder="章节名称"
          >
          </el-input>
        </el-form-item>
        <el-form-item label="文档类型" size="mini" prop="doc_type">
          <el-select
            v-model="chapter.doc_type"
            placeholder="请选择"
            style="width: 200px"
          >
            <el-option label="Markdown" :value="1"></el-option>
            <el-option label="流程图" :value="2"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button size="mini" @click="editChapterDialogVisible = false">{{
          $t("lang.cancel")
        }}</el-button>
        <el-button size="mini" type="primary" @click="chapterSave(chapter)">{{
          $t("lang.confirm")
        }}</el-button>
      </span>
    </el-dialog>
    <!-- 编辑器 -->
    <div :style="{ display: editorDisplay }" class="editor">
      <div class="editor-header">
        <div style="margin-left: 20px">{{ currentChapter.name }}</div>
        <div style="flex: 1; margin-right: 20px">
          <el-button
            size="mini"
            type="primary"
            style="float: right; margin-left: 5px"
            @click="chapterContentSave"
            icon="el-icon-folder-checked"
            >{{ $t("lang.save") }}
          </el-button>
          <el-button
            size="mini"
            style="float: right"
            @click="
              editorDisplay = 'none';
              isEditing = false;
            "
            >{{ $t("lang.cancel") }}
          </el-button>
        </div>
      </div>
      <div class="editor-body">
        <!-- 文本编辑器 -->
        <div
          :style="{ display: currentChapter.doc_type == 0 ? 'block' : 'none' }"
        ></div>
        <!-- markdown编辑器 -->
        <div
          :style="{ display: currentChapter.doc_type == 1 ? 'block' : 'none' }"
          style="width: 100%; height: 100%"
        >
          <div :id="vditorID" style="height: 98%"></div>
        </div>
        <!-- 流程图编辑器 -->
        <div
          :style="{ display: currentChapter.doc_type == 2 ? 'block' : 'none' }"
        >
          <GraphEditor el="graphEditor" ref="graphEditor"></GraphEditor>
        </div>
      </div>
    </div>
    <!-- 章节历史版本 -->
    <div class="contentList" v-if="contentListVisible">
      <el-card
        class="box-card"
        style="flex: 1; height: 0; display: flex; flex-direction: column"
        body-style="flex:1;overflow-y:auto;padding:20px;"
      >
        <div slot="header">
          <span>{{ $t("lang.history") }}</span>
          <el-button
            style="float: right; padding: 3px 0; color: grey"
            @click="contentListVisible = false"
            icon="el-icon-close"
            type="text"
            >{{ $t("lang.close") }}</el-button
          >
        </div>
        <div>
          <el-timeline style="margin: 5px; margin-left: -50px">
            <el-timeline-item
              v-for="item in contentList"
              :key="item.version"
              :timestamp="item.name + '提交于' + item.create_time_str"
              placement="top"
            >
              <el-card body-style="padding:10px;" style="cursor: pointer">
                <div style="font-size: 12px">
                  {{ item.comment ? item.comment : "没有备注~" }}
                </div>
                <div style="float: right; margin-top: 5px">
                  <el-tooltip
                    class="item"
                    effect="dark"
                    content="回滚至该版本"
                    placement="top-start"
                  >
                    <i
                      class="iconfont icon-rollback"
                      @click="rollbackChapterRevision(item.version)"
                      style="cursor: pointer; color: #25b864"
                    ></i>
                  </el-tooltip>
                  <el-tooltip
                    class="item"
                    effect="dark"
                    content="查看"
                    placement="top-start"
                  >
                    <i
                      class="iconfont icon-read"
                      @click="viewChapterRevision(item.version)"
                      style="cursor: pointer; color: #25b864"
                    ></i>
                  </el-tooltip>
                </div>
              </el-card>
            </el-timeline-item>
          </el-timeline>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script>
//markdow编辑器
import { Loading } from "element-ui";
import Vditor from "vditor";
import "~/vditor/src/assets/scss/classic.scss";
import GraphEditor from "../editor/GraphEditor";

export default {
  name: "DocView",
  components: { GraphEditor },
  props: {
    docID: {
      type: String,
      required: true,
    },
    docName: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      isEditing: false,
      addChapterDialog: true,
      graphPreviewID: "graphPreviewID" + new Date().getTime(),
      vditorID: "vditor" + new Date().getTime(),
      vditorPreviewID: "vditorPreview" + new Date().getTime(),
      vditor: null, //markdow编辑器
      vditorPreview: null,
      editorDisplay: "none",
      graph: null, //流程图显示
      graphFullScreen: false,
      graphCursor: "auto",
      centerCenterBodyStyle: "position:static",
      searchChapterKeyword: "",
      chapterList: [],
      chapterListTemp: [],
      treeProps: {
        label: "name",
        children: "children",
        isLeaf: "leaf",
      },
      node: null,
      closeOnClickModal: false,
      editChapterDialogVisible: false,
      // 章节编辑
      chapter: {
        id: "",
        doc_id: this.docID,
        doc_type: 1,
        sort: 0,
        pid: "0",
        name: "",
      },
      // 当前正在编辑的章节
      currentChapter: {
        id: "",
        doc_id: this.docID,
        doc_type: 1,
        name: "",
        content: "",
      },
      rules: {
        name: [
          {
            type: "string",
            required: true,
            message: "章节名称不能为空",
            trigger: "blur",
          },
        ],
      },
      contentList: [],
      contentListVisible: false,
      // 定时器
      timer_5s: "",
    };
  },
  watch: {
    // 监听章节关键字输入框值的变化
    searchChapterKeyword(val) {
      this.$refs.tree.filter(val);
    },
  },
  methods: {
    // 是否允许章节拖拽
    allowChapterDrag(draggingNode) {
      return draggingNode.id != "0";
    },
    // 是否允许章节放置
    allowChapterDrop(draggingNode, dropNode, type) {
      if (dropNode.data.id == "0") {
        return type == "inner";
      } else {
        return true;
      }
    },
    // 拖拽结束
    handleChapterDragEnd(draggingNode, dropNode, dropType, ev) {
      if (dropType == "inner") {
        // 修改章节
        this.chapter.id = draggingNode.data.id;
        this.chapter.name = draggingNode.data.name;
        this.chapter.doc_type = draggingNode.data.doc_type;
        this.chapter.sort = draggingNode.data.sort;
        this.chapter.pid = dropNode.data.id;
        this.chapterSave(this.chapter);
      } else {
        let chapter = {
          id: draggingNode.data.id,
          name: draggingNode.data.name,
          doc_type: draggingNode.data.doc_type,
          pid: dropNode.data.pid,
        };
        this.chapterSave(chapter);
        /*批量修改排序*/
        let idSorts = "";
        let childNodes = dropNode.parent.childNodes;
        for (var i = 0; i < childNodes.length; i++) {
          if (childNodes[i].data.sort != i) {
            if (idSorts != "") {
              idSorts = idSorts + "," + childNodes[i].data.id + "-" + i;
            } else {
              idSorts = childNodes[i].data.id + "-" + i;
            }
          }
        }
        if (idSorts != "") {
          //保存新的排序
          let data = {
            id_sorts: idSorts,
          };
          this.Request.chapterBatchSort(data)
            .then((response) => {
              if (response.data.code != 0) {
                this.$message(response.data.msg);
              }
            })
            .catch((err) => {});
        }
      }
    },
    //流程图预览移动开始
    graphPreviewMoveStart(e) {
      if (e.which == 3) {
        this.graphCursor = "move";
      }
    },
    //流程图预览移动结束
    graphPreviewMoveEnd(e) {
      if (e.which == 3) {
        this.graphCursor = "auto";
      }
    },
    //流程图预览全屏控制
    graphFullScreenChange() {
      this.graphFullScreen = !this.graphFullScreen;
      if (this.graphFullScreen) {
        this.centerCenterBodyStyle =
          "position:absolute;top:0;left:0;width:100%;height:100%;background:#fff;";
      } else {
        this.centerCenterBodyStyle = "position:static;";
      }
      this.$nextTick(() => {
        this.setGraphPreview(this.currentChapter.content);
      });
    },
    // 获取章节内容
    async getChapterContent(chapter) {
      let loadingInstance = Loading.service({
        target: document.getElementById("editor_preview"),
      });
      this.isEditing = false;
      if (chapter.id == "0") return;
      this.currentChapter.id = chapter.id;
      this.currentChapter.name = chapter.name;
      this.currentChapter.doc_type = chapter.doc_type;
      let data = {
        chapter_id: chapter.id,
      };
      this.$nextTick(() => {
        //获取章节内容并根据章节内容类型进行显示
        this.Request.chapterContent(data)
          .then((response) => {
            loadingInstance.close();
            if (response.data.code == 0) {
              let content = response.data.data;
              this.currentChapter.content = content;
              let type = parseInt(this.currentChapter.doc_type);
              switch (type) {
                //markdown
                case 1:
                  this.setMarkdownPreview(content);
                  break;
                //流程图
                case 2:
                  /*流程图预览-初始化*/
                  this.setGraphPreview(content);
                  break;
                default:
                  console.log("类型有误");
                  break;
              }
            } else {
              this.$message(response.data.msg);
            }
          })
          .catch((err) => {});
      });
    },
    //返回
    back() {
      this.$emit("docViewHide");
    },
    //获取章节列表
    getChapterList() {
      let data = {
        doc_id: this.docID,
        pid: 0,
      };
      this.Request.chapterList(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.chapterList = response.data.data;
            this.editChapterDialogVisible = false;
            this.$nextTick(() => {
              //设置当前的要选中的章节
              let chapterID = sessionStorage.getItem("chapterID");
              let docType = sessionStorage.getItem("docType");
              let chapterName = sessionStorage.getItem("chapterName");
              if (chapterID) {
                this.chapter.id = chapterID;
                this.chapter.name = chapterName;
                this.chapter.doc_type = parseInt(docType);
                this.$refs.tree.setCurrentKey(chapterID);
                this.getChapterContent(this.chapter);
                sessionStorage.removeItem("chapterID");
                sessionStorage.removeItem("docType");
                sessionStorage.removeItem("chapterName");
              }
            });
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    //新增章节
    addChapter(data, node) {
      this.node = node;
      this.chapter.id = "";
      this.chapter.doc_type = 1;
      this.chapter.sort = 0;
      this.chapter.name = "";
      this.chapter.pid = data.id;
      this.editChapterDialogVisible = true;
      this.addChapterDialog = true;
    },
    //编辑章节
    editChapter(data, node) {
      this.node = node;
      this.chapter.id = data.id;
      this.chapter.doc_type = data.doc_type;
      this.chapter.sort = data.sort;
      this.chapter.name = data.name;
      this.chapter.pid = data.pid;
      this.editChapterDialogVisible = true;
      this.addChapterDialog = false;
    },
    // 编辑章节内容
    editChapterContent() {
      this.isEditing = true;
      this.editorDisplay = "block";
      let docType = parseInt(this.currentChapter.doc_type);
      let content = this.currentChapter.content;
      // 从本地缓存读取缓存的值
      let key = "doc_" + this.currentChapter.id;
      let contentCache = window.localStorage.getItem(key);
      console.log(JSON.stringify(contentCache));
      if (contentCache && contentCache != "\n" && content != "\r\n") {
        this.$confirm("检测到未成功保存的内容,是否加载?", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "放弃",
          type: "warning",
        })
          .then(() => {
            this.setEditorContent(docType, contentCache);
          })
          .catch(() => {
            let key = "doc_" + this.currentChapter.id;
            let contentCache = window.localStorage.removeItem(key);
            this.setEditorContent(docType, content);
          });
      } else {
        this.setEditorContent(docType, content);
      }
    },
    // 设置编辑器内容
    setEditorContent(docType, content) {
      switch (docType) {
        case 1:
          this.vditor.focus();
          this.vditor.setValue(content);
          break;
        case 2:
          this.$nextTick(() => {
            this.$refs.graphEditor.refresh();
            this.$refs.graphEditor.refresh();
            let initData =
              '<mxGraphModel dx="946" dy="254" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169"><root><mxCell id="0"/><mxCell id="1" parent="0"/></root></mxGraphModel>';
            //设置内容
            try {
              this.$refs.graphEditor.setValue(content);
            } catch (err) {
              this.$refs.graphEditor.setValue(initData);
            }
          });
          break;
        default:
          break;
      }
    },
    //保存章节
    chapterSave(data) {
      this.Request.chapterSave(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.currentChapter.doc_type = this.chapter.doc_type;
            this.editChapterDialogVisible = false;
            //  保存最近的章节，以便在刷新加载章节列表时候自动显示
            sessionStorage.setItem("chapterID", this.chapter.id);
            sessionStorage.setItem("docType", this.chapter.doc_type);
            sessionStorage.setItem("chapterName", this.chapter.name);
            this.getChapterList();
          } else {
            // 保存失败，提示失败信息
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    //节点过滤
    filterChapter(value, data) {
      if (!value) return true;
      return data.name.indexOf(value) !== -1;
    },
    // 获取编辑器内容
    getEditorContent() {
      let content = "";
      let docType = parseInt(this.currentChapter.doc_type);
      switch (docType) {
        case 1:
          content = this.vditor.getValue();
          break;
        case 2:
          content = this.$refs.graphEditor.getValue();
          break;
        default:
          break;
      }
      return content;
    },
    // 保存
    chapterContentSave() {
      this.$prompt("", "提交备注", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
      })
        .then(({ value }) => {
          let content = "";
          let docType = parseInt(this.currentChapter.doc_type);
          switch (docType) {
            case 1:
              content = this.vditor.getValue();
              break;
            case 2:
              content = this.$refs.graphEditor.getValue();
              break;
            default:
              break;
          }

          let data = {
            chapter_id: this.currentChapter.id,
            content: content,
            comment: value,
          };
          this.Request.chapterContentSave(data)
            .then((response) => {
              this.isEditing = false; //退出编辑状态
              if (response.data.code == 0) {
                this.editorDisplay = "none";
                this.currentChapter.content = content;
                this.getChapterContent(this.currentChapter);
                //如果保存成功，则删除localstorage的缓存值
                let key = "doc_" + this.currentChapter.id;
                window.localStorage.removeItem(key);
              } else {
                // 保存失败，提示失败信息
                this.$message(response.data.msg);
              }
            })
            .catch((err) => {});
        })
        .catch(() => {});
    },
    //设置markdown预览
    setMarkdownPreview(content) {
      this.vditorPreview.setValue(content);
    },
    //设置流程图预览
    setGraphPreview(content) {
      let initData =
        '<mxGraphModel dx="946" dy="254" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169"><root><mxCell id="0"/><mxCell id="1" parent="0"/></root></mxGraphModel>';
      this.graph.resizeContainer = false;
      this.graph.setEnabled(false);
      // 先进行预览初始化
      let xmlDoc0 = mxUtils.parseXml(initData);
      let codec0 = new mxCodec(xmlDoc0);
      codec0.decode(xmlDoc0.documentElement, this.graph.getModel());

      let xmlDoc = mxUtils.parseXml(content);
      let codec = new mxCodec(xmlDoc);
      codec.decode(xmlDoc.documentElement, this.graph.getModel());
    },
    showChapterContentList() {
      this.contentListVisible = !this.contentListVisible;
      this.getChapterContentList();
    },
    //获取章节内容历史版本
    getChapterContentList() {
      let data = {
        chapter_id: this.currentChapter.id,
      };
      this.Request.chapterContentList(data)
        .then((response) => {
          if (response.data.code == 0) {
            this.contentList = response.data.data;
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    // 查看版本内容
    viewChapterRevision(version) {
      let data = {
        id: this.currentChapter.id,
        version: version,
      };
      this.Request.getChapterRevision(data)
        .then((response) => {
          if (response.data.code == 0) {
            let revision = response.data.data;
            let docType = parseInt(this.currentChapter.doc_type);
            switch (docType) {
              case 1:
                this.setMarkdownPreview(revision);
                break;
              case 2:
                this.setGraphPreview(revision);
                break;
              default:
                break;
            }
          } else {
            this.$message(response.data.msg);
          }
        })
        .catch((err) => {});
    },
    // 章节回滚
    rollbackChapterRevision(version) {
      this.$confirm("此操作将会回滚版本, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          let data = {
            id: this.currentChapter.id,
            version: version,
          };
          this.Request.setChapterRevision(data)
            .then((response) => {
              if (response.data.code == 0) {
                this.getChapterContent(this.currentChapter);
              } else {
                this.$message(response.data.msg);
              }
            })
            .catch((err) => {});
        })
        .catch(() => {});
    },
    // 复位章节表单
    resetChapterForm() {
      this.$refs.chapterForm.resetFields();
    },
    // 定时保存编辑器内容到缓存中
    autoSaveToCache() {
      if (this.isEditing) {
        //判断处于编辑状态
        // 获取编辑的内容，保存到localstorage中
        let content = this.getEditorContent();
        let key = "doc_" + this.currentChapter.id;
        // key: doc_chapterID
        if (content != this.currentChapter.content) {
          window.localStorage.setItem(key, content);
        } else {
          window.localStorage.removeItem(key);
        }
      }
    },
  },
  created() {
    //加载章节列表
    this.getChapterList();
  },
  mounted() {
    /*markdown预览-初始化*/
    let md_preview_options = {
      typewriterMode: true,
      width: "100%",
      toolbar: ["preview", "fullscreen"],
      preview: {
        delay: 100,
        mode: "preview",
      },
      //禁用默认的本地缓存
      cache: {
        enable: false,
      },
    };
    this.vditorPreview = new Vditor(this.vditorPreviewID, md_preview_options);
    this.vditorPreview.disabledCache(); //禁用本地缓存
    this.vditorPreview.clearCache();
    /*markdown编辑器*/
    let md_options = {
      typewriterMode: true,
      width: "100%",
      toolbar: [
        "headings",
        "bold",
        "italic",
        "strike",
        "|",
        "line",
        "quote",
        "|",
        "list",
        "ordered-list",
        "check",
        "|",
        "code",
        "inline-code",
        "|",
        "undo",
        "redo",
        "|",
        "upload",
        "link",
        "table",
        "|",
        "both",
        "preview",
        "format",
        "|",
        "fullscreen",
        "help",
      ],
      upload: {
        url: domain + "/upload/vditorUpload",
        // filename:'file',
        //文件最大200M
        max: 200 * 1024 * 1024,
      },
      preview: {
        delay: 200,
      },
    };
    this.vditor = new Vditor(this.vditorID, md_options);
    this.vditor.disabledCache(); //禁用本地缓存
    this.vditor.clearCache();
    //流程图预览初始化
    this.graph = new Graph(document.getElementById(this.graphPreviewID));

    // 增加一个定时器
    this.timer_5s = setInterval(this.autoSaveToCache, 5000);
  },
  beforeDestroy() {
    clearInterval(this.timer_5s);
  },
};
</script>

<style scoped>
.doc-view-main {
  height: 100%;
  width: 100%;
  overflow-x: auto;
  display: flex;
  display: -webkit-flex;
  flex-direction: column;
  align-items: stretch;
}
.doc-view-header {
  height: 60px;
  border-bottom: 1px solid #dcdfe6;
  display: flex;
  display: -webkit-flex;
  flex-direction: row;
  align-items: center;
}
.doc-view-body {
  flex: 1;
  height: 0;
  display: flex;
  display: -webkit-flex;
  flex-direction: row;
  align-items: stretch;
}
.left {
  min-width: 300px;
  overflow-y: auto;
  border-right: 1px solid #dcdfe6;
}
.center-left {
  flex: 1;
}
.center-right {
  flex: 1;
  display: flex;
  display: -webkit-flex;
  flex-direction: column;
  align-items: flex-start;
  padding-top: 40px;
}
.center-center {
  width: 875px;
  display: flex;
  display: -webkit-flex;
  flex-direction: column;
  align-items: stretch;
}
.center-center-header {
  padding-top: 10px;
  height: 30px;
  font-size: 18px;
  font-weight: 700;
}
.center-center-body {
  flex: 1;
  height: 0;
  align-items: stretch;
}
.custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 14px;
  padding-right: 8px;
}
.tree-button {
  display: none;
}
.custom-tree-node:hover .tree-button {
  display: inline;
}
.editor {
  position: fixed;
  z-index: 200;
  background: #fff;
  height: 100%;
  width: 100%;
  left: 0;
  top: 0;
}
.editor-header {
  width: 100%;
  height: 40px;
  font-size: 18px;
  font-weight: 700;
  display: flex;
  display: -webkit-flex;
  flex-direction: row;
  align-items: center;
}
.editor-body {
  position: absolute;
  width: 100%;
  top: 40px;
  bottom: 0px;
  left: 0px;
  right: 0px;
}
.graphTool:hover {
  cursor: pointer;
}
/* 章节历史版本 */
.contentList {
  position: absolute;
  top: 60px;
  left: 0;
  height: 100%;
  width: 300px;
  background: white;
  border: 1px solid #dcdfe6;
  display: flex;
  display: -webkit-flex;
  flex-direction: column;
}
</style>
