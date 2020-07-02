import Vue from "vue";
import Vuex from "vuex";
// import axios from "axios";

// axios.defaults.headers.common["Access-Control-Allow-Origin"] = "*";

Vue.use(Vuex);
export default new Vuex.Store({
  state: {
    users: [
      {
        id: 1,
        img: require("./assets/logo.png"),
        name: "鈴木健介",
        grade: "3",
        major: "kk",
        group: "インフラ班"
      },
      {
        id: 2,
        img: require("./assets/logo.png"),
        name: "山田太郎",
        grade: "3",
        major: "kk",
        group: "クリエイティブ班"
      },
      {
        id: 3,
        img: require("./assets/logo.png"),
        name: "田中綾子",
        grade: "2",
        major: "kk",
        group: "インフラ班"
      },
      {
        id: 4,
        img: require("./assets/logo.png"),
        name: "山下恵子",
        grade: "1",
        major: "kx",
        group: "クリエイティブ班"
      },
      {
        id: 5,
        img: require("./assets/logo.png"),
        name: "松本康太",
        grade: "3",
        major: "kx",
        group: "クリエイティブ班"
      }
    ],
    filter: "",
    // show_filter: [
    //   {
    //     grade: false,
    //     major: false,
    //     position: false,
    //     group: false
    //   }
    // ],
    filter_grade: "",
    filter_major: "",
    filter_group: "",
    filter_technology: "",
    filter_technical_area: "",
    filterlist: ["選択", "学年", "専攻", "所属", "言語", "技術"],
    grade: ["1", "2", "3", "4"],
    major: [
      "EE",
      "EV",
      "CC",
      "CB",
      "MM",
      "MP",
      "DD",
      "DS",
      "FA",
      "FL",
      "HT",
      "HH",
      "KK",
      "KX"
    ],
    group: ["クリエイティブ班", "インフラ班"],
    technology: ["C", "Java"],
    technical_area: ["フロントエンド", "バックエンド"]
  },
  mutations: {
    update_filter_grade(store, grade) {
      store.filter_grade = grade;
    },
    update_filter_major(store, major) {
      store.filter_major = major;
    },
    update_filter_group(store, group) {
      store.filter_group = group;
    },
    update_filter_technology(store, technology) {
      store.filter_technology = technology;
    },
    update_filter_technical_area(store, technical_area) {
      store.filter_technical_area = technical_area;
    }
  },
  actions: {}
});
