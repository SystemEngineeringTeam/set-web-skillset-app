import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

axios.defaults.headers.common["Access-Control-Allow-Origin"] = "*";

Vue.use(Vuex);

export default new Vuex.Store({
  strict: true,
  state: {
    email: "",
    status: false
  },
  getters: {
    email(state) {
      return state.email;
    },
    isSignedIn(state) {
      return state.status;
    }
  },
  mutations: {
    onAuthEmailChanged(state, email) {
      state.email = email; //firebase user情報
    },
    onUserStatusChanged(state, status) {
      state.status = status;
    }
  }
});
