<template>
  <v-app>
    <v-content>
      <router-link to="/">
        <v-btn>homeページに遷移する</v-btn>
      </router-link>
      <router-link to="/test">
        <v-btn>testページに遷移する</v-btn>
      </router-link>
      <router-link to="/Registration_Form">
        <v-btn>登録ページに遷移する</v-btn>
      </router-link>
      <router-view />
    </v-content>
    <Footer />
  </v-app>
</template>

<script>
import Footer from "./components/footer";
export default {
  name: "App",
  components: {
    Footer
  },

  data: () => ({})
};
</script>
