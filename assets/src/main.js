import Vue from "vue";
import App from "./App.vue";
import axios from "axios";
import router from "./router";
import Vuetify from "vuetify"; //追加
import store from "./store";
import "vuetify/dist/vuetify.min.css"; //追加

Vue.use(Vuetify); //追加

const vuetify = new Vuetify(); //追加

Vue.prototype.$axios = axios;

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  vuetify, //追加
  render: h => h(App)
}).$mount("#app");
