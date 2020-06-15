<template>
  <div id="Login">
    <v-sheet 
      :elevation="10"
      class="mx-auto pt-5"
      height="300"
      width="500"
      >
       
        <v-form 
            ref="form"
            v-model="valid"
            lazy-validation
        >
            
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
                label="Password"
                required
            ></v-text-field>

            <v-btn block
                :disabled="!valid"
                color="success"
                @click="login"
            >ログイン</v-btn>


        </v-form>
    </v-sheet>
  </div>
</template>

<script>
  export default {
    data: () => ({
      valid: true,
      email: '',
      emailRules: [
        v => !!v || 'E-mailを入力してね',
        v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
      ],
      password: '',
      passwordRules: [
        v => !!v || 'passwordを入力してね',
      ]
    }),

    methods: {
      login () {
          this.$refs.form.login()
      },
    },
  }
</script>