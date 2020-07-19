import Vue from "vue";
import Vuex from "vuex";
// import axios from "axios";

// axios.defaults.headers.common["Access-Control-Allow-Origin"] = "*";

Vue.use(Vuex);
export default new Vuex.Store({
  state: {
    users: [{
        id: 1,
        img: require("./assets/logo.png"),
        name: "鈴木健介",
        grade: "3",
        major: "KK",
        group: "インフラ班",
        technology: "C",
        technical_area: "Webフロントエンド"
      },
      {
        id: 2,
        img: require("./assets/logo.png"),
        name: "一寸法師",
        grade: "3",
        major: "KK",
        group: "クリエイティブ班",
        technology: "Java",
        technical_area: "Webフロントエンド"
      },
      {
        id: 3,
        img: require("./assets/logo.png"),
        name: "白雪姫",
        grade: "2",
        major: "KK",
        group: "インフラ班",
        technology: "JavaScript",
        technical_area: "Webフロントエンド"
      },
      {
        id: 4,
        img: require("./assets/logo.png"),
        name: "ピノキオ",
        grade: "1",
        major: "KX",
        group: "クリエイティブ班",
        technology: "C++",
        technical_area: "Webバックエンド"
      },
      {
        id: 5,
        img: require("./assets/logo.png"),
        name: "赤ずきん",
        grade: "3",
        major: "KX",
        group: "クリエイティブ班",
        technology: "C#",
        technical_area: "Webバックエンド"
      },
      {
        id: 6,
        img: require("./assets/logo.png"),
        name: "桃太郎",
        grade: "4",
        major: "KX",
        group: "クリエイティブ班",
        technology: "PHP",
        technical_area: "Webバックエンド"
      }
    ],
    filter: "",
    filter_grade: "",
    filter_major: "",
    filter_group: "",
    filter_technology: "",
    filter_technical_area: "",
    change_display: "",
    filterlist: ["選択", "学年", "専攻", "所属", "言語", "技術"],
    grades: ["1", "2", "3", "4"],
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
    technology: ["C",
      "C++",
      "C#",
      "Java",
      "Python2",
      "Python3",
      "PHP",
      "Ruby",
      "Perl",
      "Objective-C",
      "Swift",
      "VB",
      "ActionScript",
      "R言語",
      "Groovy",
      "Haskell",
      "Scala",
      "Erlang",
      "Go言語",
      "JavaScript",
      "HTML5",
      "CSS3",
      "CoffeeScript",
      "Haml",
      "Sass",
      "TypeScript",
      "SQL",
      "PL/SQL",
      "COBOL",
      "ABAP",
      "Kotlin",
      "Rust",
      "AHDL",
      "Bash",
      "Clojure",
      "ClojureScript",
      "Common Lisp",
      "CUDA C/C++",
      "ECMAScript",
      "Elixir",
      "F#",
      "Oklahoma",
      "Oregon",
      "Palau",
      "Pennsylvania",
      "Puerto Rico",
      "Lua",
      "MATLAB",
      "Nim",
      "Objective Caml (OCaml)",
      "Prolog",
      "Scheme",
      "Verilog HDL",
      "VHDL",
    ],
    technical_area: ["Webフロントエンド", "Webバックエンド"]
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