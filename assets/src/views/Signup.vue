<template>
  <div id="Login">
    <v-sheet 
      :elevation="10"
      class="mx-auto pt-5"
      height="500"
      width="500"
      >
       
        <v-form 
            ref="form"
            v-model="valid"
            lazy-validation
        >
         
            <v-text-field 
              class="pa-4"
              v-model="name"
              :counter="10"
              :rules="nameRules"
              label="名前"
              required
            ></v-text-field>

            <v-text-field
                class="pa-4"
                v-model="email"
                :rules="emailRules"
                label="E-mail"
                required
            ></v-text-field>

            <v-text-field
                class="pa-4"
                v-model="password"
                :rules="passwordRules"
                label="パスワード"
                required
            ></v-text-field>

            <v-text-field
                class="pa-4"
                v-model="password"
                :rules="passwordRules"
                label="パスワード再入力"
                required
            ></v-text-field>

            <v-btn block
                :disabled="!valid"
                color="success"
                @click="signup"
            >登録</v-btn>

        </v-form>
    </v-sheet>
  </div>
</template>

<script>
  export default {
    data: () => ({
      valid: true,
      name: '',
      nameRules: [
        v => !!v || '名前を入力してね',
        v => (v && v.length <= 10) || 'Name must be less than 10 characters',
      ],
      email: '',
      emailRules: [
        v => !!v || 'E-mailを入力してね',
        v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
      ],
      password: '',
      passwordRules: [
        v => !!v || 'パスワードを入力してね',
      ],
    }),

    methods: {
      signup () {
        this.$refs.form.signup()
      },
    },
  }
</script>