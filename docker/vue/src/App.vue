<template>
  <v-app>
    <v-content>
       <router-link to="/">
        <v-btn>
          homeページに遷移する
        </v-btn>
      </router-link>
       <router-link to="/test">
        <v-btn>
          testページに遷移する
        </v-btn>
      </router-link>
      <router-view/>
    </v-content>
  </v-app>
</template>

<script>

export default {
  name: 'App',

  // data: () => ({
  //   //
  // }),
};
</script>
