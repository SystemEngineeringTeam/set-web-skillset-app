import Vue from "vue";
import Router from "vue-router";
import Home from "./pages/home";
import Test from "./pages/test";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "home",
      component: Home
    },
    {
      path: "/test",
      name: "test",
      component: Test
    }
  ]
});
