<template>
  <v-container>
    <!-- fluidを付けることで左右いっぱいにコンテナが広がります -->
    <div class="registration">
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-card class="mx-auto" max-width="800" outlined>
          <v-list-item three-line>
            <v-list-item-content>
              <v-list-item-title class="headline mb-1"
                >アイコン</v-list-item-title
              >
              <v-file-input
                :rules="rules"
                accept="image/png, image/jpeg, image/bmp"
                placeholder="Pick an avatar"
                prepend-icon="mdi-camera"
                label="Avatar"
              ></v-file-input>
            </v-list-item-content>
          </v-list-item>
        </v-card>
        <v-card class="mx-auto" max-width="800" outlined>
          <v-list-item three-line>
            <v-list-item-content>
              <v-list-item-title class="headline mb-1">名前</v-list-item-title>
              <v-form ref="form" v-model="valid" lazy-validation>
                <v-text-field
                  v-model="name"
                  :counter="10"
                  :rules="nameRules"
                  label="山田 太郎"
                  required
                ></v-text-field>
              </v-form>
            </v-list-item-content>
          </v-list-item>
        </v-card>
        <v-card class="mx-auto" max-width="800" outlined>
          <v-list-item three-line>
            <v-list-item-content>
              <v-list-item-title class="headline mb-1">学年</v-list-item-title>
              <v-form ref="form" v-model="valid" lazy-validation>
                <v-select
                  v-model="select"
                  :items="grade"
                  :rules="[(v) => !!v || '学年を入力してください']"
                  label="学年"
                  required
                ></v-select>
              </v-form>
            </v-list-item-content>
          </v-list-item>
        </v-card>
        <v-card class="mx-auto" max-width="800" outlined>
          <v-list-item three-line>
            <v-list-item-content>
              <v-list-item-title class="headline mb-1">専攻</v-list-item-title>
              <v-form ref="form" v-model="valid" lazy-validation>
                <v-select
                  v-model="select2"
                  :items="major"
                  :rules="[(v) => !!v || '専攻を入力してください']"
                  label="専攻"
                  required
                ></v-select>
              </v-form>
            </v-list-item-content>
          </v-list-item>
        </v-card>
        <v-card class="mx-auto" max-width="800" outlined>
          <v-list-item three-line>
            <v-list-item-content>
              <v-list-item-title class="headline mb-1">所属</v-list-item-title>
              <v-form ref="form" v-model="valid" lazy-validation>
                <v-select
                  v-model="select3"
                  :items="group"
                  label="所属"
                  required
                ></v-select>
              </v-form>
            </v-list-item-content>
          </v-list-item>
        </v-card>
        <v-card class="mx-auto" max-width="800" outlined>
          <v-list-item three-line>
            <v-list-item-content>
              <v-list-item-title class="headline mb-1"
                >得意な技術領域</v-list-item-title
              >
              <v-form ref="form" v-model="valid" lazy-validation>
                <v-select
                  v-model="select4"
                  :items="technical_area"
                  :rules="[(v) => !!v || 'Item is required']"
                  label="得意な技術領域"
                  required
                ></v-select>
              </v-form>
            </v-list-item-content>
          </v-list-item>
        </v-card>

        <v-card class="mx-auto" max-width="800" outlined>
          <v-list-item three-line>
            <v-list-item-content>
              <v-list-item-title class="headline mb-1">技術</v-list-item-title>
              <v-select
                v-model="e7"
                :items="states"
                label="使える言語など"
                multiple
                chips
              ></v-select>
            </v-list-item-content>
          </v-list-item>
        </v-card>
        <v-card class="mx-auto" max-width="800" outlined>
          <v-list-item three-line>
            <v-list-item-content>
              <v-list-item-title class="headline mb-1"
                >自己紹介</v-list-item-title
              >
              <v-textarea
                counter
                label="簡単に自分を紹介してください"
                :rules="rules"
                :value="value"
              ></v-textarea>
            </v-list-item-content>
          </v-list-item>
        </v-card>

        <v-checkbox
          v-model="checkbox"
          :rules="[(v) => !!v || 'You must agree to continue!']"
          label="Do you agree?"
          required
        ></v-checkbox>

        <v-btn
          :disabled="!valid"
          color="success"
          class="mr-4"
          @click="validate"
        >
          Validate
        </v-btn>

        <v-btn color="error" class="mr-4" @click="reset">
          Reset Form
        </v-btn>

        <v-btn color="warning" @click="resetValidation">
          Reset Validation
        </v-btn>
      </v-form>
    </div>
  </v-container>
</template>
<script>
export default {
  data: () => ({
    valid: true,
    name: "",
    nameRules: [
      (v) => !!v || "名前の入力をしてください",
      (v) => (v && v.length <= 10) || "Name must be less than 10 characters",
    ],
    select: null,
    select2: null,
    select3: null,
    select4: null,
    grade: ["1", "2", "3", "4"],
    major: ["kk", "kx"],
    items: ["item 1", "item 2"],
    group: ["クリエイティブ班", "インフラ班"],
    technical_area: [
      "スマホ開発（IOS）",
      "スマホ開発（Android）",
      "ゲーム開発（Unity）",
      "Webフロント",
      "サーバーサイド",
      "AI",
      "インフラ",
    ],
    e7: [],
    states: [
      "C",
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
    checkbox: false,
  }),

  methods: {
    validate() {
      this.$refs.form.validate();
    },
    reset() {
      this.$refs.form.reset();
    },
    resetValidation() {
      this.$refs.form.resetValidation();
    },
  },
};
</script>
