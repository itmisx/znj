// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue";
import App from "./App";
import router from "./router";
import Element from "element-ui";
// 通过scss自定义样式
import "../element-variables.scss";
// 加载iconfont矢量图
import "./assets/iconfont/iconfont.css";
// 加载请求定义的请求
import * as request from "./common/request";

Vue.use(Element);

Vue.config.productionTip = false;

//引入接口定义
Vue.prototype.Request = request;

import VueI18n from "vue-i18n";

Vue.use(VueI18n); // 通过插件的形式挂载

const i18n = new VueI18n({
  locale: "zh", // 语言标识
  //this.$i18n.locale // 通过切换locale的值来实现语言切换
  messages: {
    zh: require("./config/lang/zh"), // 中文语言包
    en: require("./config/lang/en"), // 英文语言包
  },
});

/* eslint-disable no-new */
new Vue({
  el: "#app",
  i18n,
  router,
  components: { App },
  template: "<App/>",
});
