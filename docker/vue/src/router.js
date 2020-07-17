import Vue from "vue";
import Router from "vue-router";
import Home from "./pages/Home";
import Test from "./pages/Test";
import Registration_Form from "./pages/Registration_Form";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [{
      path: "/",
      name: "Home",
      component: Home
    },
    {
      path: "/Test",
      name: "Test",
      component: Test
    },
    {
      path: "/Registration_Form",
      name: "Registration_Form",
      component: Registration_Form
    }
  ]
});