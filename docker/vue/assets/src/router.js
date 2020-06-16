import Vue from "vue";
import Router from "vue-router";
import Home from "./views/Home";
import Signup from "./views/Signup";

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
    }
  ]
});
