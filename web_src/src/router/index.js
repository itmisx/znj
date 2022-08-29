import Vue from "vue";
import Router from "vue-router";
// import Main from '@/components/Main'
// import Login from '@/components/Login'
// import Todolist from '@/components/todo/Todolist'
// import List from "@/components/project/List";
// import WikiList from "../components/doc/WikiList";
// import Project from "../components/project/Project";
// import ProjectBoard from "@/components/project/ProjectBoard";
// import ProjectDoc from "../components/project/ProjectDoc";
// import ProjectMember from "../components/project/ProjectMember";
// import Team from "../components/team/Team"
// import TeamList from "../components/team/TeamList"
// import TeamMember from "../components/team/TeamMember"
// import DocView  from "../components/doc/DocView"
// import Help from "../components/help/Help"

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/main",
      component: (resolve) => require(["@/components/Main"], resolve),
      children: [
        {
          path: "/main/todo-list",
          component: (resolve) =>
            require(["@/components/todo/Todolist"], resolve),
        },
        {
          path: "/main/project-list",
          component: (resolve) =>
            require(["@/components/project/List"], resolve),
        },
        {
          path: "/main/project",
          component: (resolve) =>
            require(["@/components/project/Project"], resolve),
          children: [
            {
              path: "/main/project/board",
              component: (resolve) =>
                require(["@/components/project/ProjectBoard"], resolve),
            },
            {
              path: "/main/project/doc",
              component: (resolve) =>
                require(["@/components/project/ProjectDoc"], resolve),
            },
            {
              path: "/main/project/member",
              component: (resolve) =>
                require(["@/components/project/ProjectMember"], resolve),
            },
          ],
        },
        {
          path: "/main/wiki",
          component: (resolve) =>
            require(["@/components/doc/WikiList"], resolve),
        },
        {
          path: "/main/team",
          component: (resolve) => require(["@/components/team/Team"], resolve),
          children: [
            {
              path: "/main/team/team-list",
              component: (resolve) =>
                require(["@/components/team/TeamList"], resolve),
            },
            {
              path: "/main/team/team-member",
              component: (resolve) =>
                require(["@/components/team/TeamMember"], resolve),
            },
          ],
        },
      ],
    },
    {
      path: "/login",
      component: (resolve) => require(["@/components/Login"], resolve),
    },
    {
      path: "/docview",
      component: (resolve) => require(["@/components/doc/DocView"], resolve),
    },
  ],
});
