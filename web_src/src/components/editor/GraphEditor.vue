<template>
  <div class="graph-editor-main geEditor" :id="graphEditorID"></div>
</template>

<script>
export default {
  name: "GraphEditor",
  data() {
    return {
      graphEditorID: "graphPreviewID" + new Date().getTime(),
      editorUi: null,
    };
  },
  methods: {
    //流程图编辑器初始化
    initEditorUi() {
      var that = this;
      var editorUiInit = EditorUi.prototype.init;
      EditorUi.prototype.init = function () {
        editorUiInit.apply(this, arguments);
        this.actions.get("export").setEnabled(false);
        // Updates action states which require a backend
        if (!Editor.useLocalStorage) {
          var enabled = true;
          this.actions.get("open").setEnabled(enabled || Graph.fileSupport);
          this.actions.get("import").setEnabled(enabled || Graph.fileSupport);
          this.actions.get("save").setEnabled(enabled);
          this.actions.get("saveAs").setEnabled(enabled);
          this.actions.get("export").setEnabled(enabled);
        }
      };
      // Adds required resources (disables loading of fallback properties, this can only
      // be used if we know that all keys are defined in the language specific file)
      mxResources.loadDefaultBundle = false;
      var bundle =
        mxResources.getDefaultBundle(RESOURCE_BASE, mxLanguage) ||
        mxResources.getSpecialBundle(RESOURCE_BASE, mxLanguage);

      // Fixes possible asynchronous requests
      mxUtils.getAll(
        [bundle, STYLE_PATH + "/default.xml"],
        function (xhr) {
          // Adds bundle text to resources
          mxResources.parse(xhr[0].getText());
          // Configures the default graph theme
          var themes = new Object();
          themes[Graph.prototype.defaultThemeName] =
            xhr[1].getDocumentElement();
          Graph.resizeContainer = true;
          // Main
          that.editorUi = new EditorUi(
            new Editor(urlParams["chrome"] == "0", themes),
            document.getElementById(that.graphEditorID)
          );
        },
        function () {
          document.body.innerHTML =
            '<center style="margin-top:10%;">Error loading resource files. Please check browser console.</center>';
        }
      );
    },
    //获取内容
    getValue() {
      let xml = mxUtils.getXml(this.editorUi.editor.getGraphXml());
      return xml;
    },
    //设置内容
    setValue(xml) {
      let node = mxUtils.parseXml(xml);
      this.editorUi.editor.setGraphXml(node.documentElement);
    },
    //页面刷新
    refresh() {
      this.editorUi.refresh();
    },
  },
  mounted() {
    this.initEditorUi();
  },
};
</script>
<style scoped>
.graph-editor-main {
  height: 100%;
  width: 100%;
}
</style>
>
