import Vue from "vue";
import Router from "vue-router";
import Home from "./views/Home";
import Signup from "./views/Signup";
import Login from "./views/Login";
import Hometest from "./views/Home";
import Articles from "./views/Articles";
import Chats from "./views/Chats";

import Article from "./views/Article";
import Chat from "./views/Chat";
import JavaChat from "./views/JavaChat";
import PersonalChat from "./views/PersonalChat";
import IndividualChat from "./views/IndividualChat";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    //ルーティングの設定
    {
      path: "/",
      name: "home",
      component: Home
    },
    {
      path: "/signup",
      name: "signup",
      component: Signup
    },
    {
      path: "/login",
      name: "login",
      component: Login
    },
    {
      path: "/home",
      name: "home",
      component: Hometest
    },
    {
      path: "/articles",
      name: "articles",
      component: Articles
    },
    {
      path: "/chats",
      name: "chats",
      component: Chats
    },
    {
      path: "/article",
      naame: "article",
      component: Article
    },
    {
      path: "/chat",
      name: "chat",
      component: Chat
    },
    {
      path: "/javachat",
      name: "javachat",
      component: JavaChat
    },
    {
      path: "/personalchat",
      name: "personalchat",
      component: PersonalChat
    },
    {
      path: "/individualchat",
      name: "individualchat",
      component: IndividualChat
    }
  ]
});
