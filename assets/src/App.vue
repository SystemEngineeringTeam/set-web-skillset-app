<template>
  <v-app>
  <div id="app">

    <v-app-bar color="#6200EA" dense dark >

      <v-btn @click.stop="drawer = !drawer" icon>
        <v-app-bar-nav-icon></v-app-bar-nav-icon>
      </v-btn>
      
      <v-toolbar-title>ISS</v-toolbar-title>

      
      
      <v-spacer></v-spacer>

      <v-menu left bottom>
        <template v-slot:activator="{ on }">
          <v-btn icon v-on="on">
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>

        <v-list>
          <v-list-item
            v-for="n in 5"
            :key="n"
            @click="() => {}"
          >
            <v-list-item-title>Option {{ n }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>

    
      <v-btn color="success" @click="logout">ログアウト</v-btn>

    </v-app-bar>

    

    <v-navigation-drawer v-model="drawer" absolute temporary>
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title>ようこそ</v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-list dense>
        <v-list-item v-for="item in items" :key="item.title" link>
          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
          
            
          
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <!-- <router-link to="/signup">
        signup test
    </router-link>
    <br>
    <router-link to="/login">
      login test
    </router-link>
    <br>
    <router-link to="/home">
      home test
    </router-link> -->


    <v-content>
      <!-- Provides the application the proper gutter -->
      <v-container fluid>
        <!-- If using vue-router -->
        <router-view></router-view>
      </v-container>
    </v-content>
  </div>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  data () {
      return {
        drawer: null,
        items: [
          { title: 'Home', icon: 'dashboard' },
          { title: 'About', icon: 'question_answer' },
          { title: 'Create', icon: 'create' }
        ],
      }
  },

  methods: {
      logout () {
          // this.$refs.form.logout()
        this.$router.push("/Login")
      },
    },

  // components: {
  //   HelloWorld
  // }
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  // created () {
  //   this.$axios.get('http://localhost:8080')
  //     .then(response => {
  //       console.log(response)
  //     })
  // }
}
</script>

<style>
</style>
