<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
import * as productConfig from "./config/product";
export default {
  name: "App",
  data() {
    return {};
  },
  watch: {
    $route(to, from) {
      let path = to.path;
      if (path.indexOf("/main/todo-list") != -1) this.activeIndex = "todolist";
      else if (path.indexOf("/main/project") != -1)
        this.activeIndex = "project";
      else if (path.indexOf("/main/wiki") != -1) this.activeIndex = "wiki";
      else if (path.indexOf("/main/team") != -1) this.activeIndex = "team";
    },
  },
  mounted() {
    document.title = productConfig.name;
    this.Request.isSignedIn().then((response) => {
      if (response.data.code === 0) {
        //默认为待办页面
        if (this.$route.path == "/main" || this.$route.path == "/") {
          this.$router.push({
            path: "/main/todo-list",
          });
        }
      } else {
        this.$router.push({
          path: "/login",
        });
      }
    });
  },
};
</script>
