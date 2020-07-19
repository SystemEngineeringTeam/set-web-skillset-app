<template>
  <v-hover v-slot:default="{ hover }" open-delay="200">
    <v-card
      v-if="
        $store.state.filter_grade === user.grade ||
          $store.state.filter_major === user.major ||
          $store.state.filter_group === user.group ||
          $store.state.filter_technology === user.technology ||
          $store.state.filter_technical_area === user.technical_area ||
          (!$store.state.filter_grade &&
            !$store.state.filter_major &&
            !$store.state.filter_group &&
            !$store.state.filter_technology &&
            !$store.state.filter_technical_area)
      "
      :elevation="hover ? 16 : 2"
      class="d-flex mx-auto card"
      max-width="400"
    >
      <v-responsive
        class="text-center grey lighten-2 rounded-circle d-inline-flex align-center justify-center ma-3"
        height="108"
        width="108"
      >
        <v-img v-bind:src="user.img"></v-img>
      </v-responsive>
      <v-card-text>
        <p class="display-1 text--primary">{{ user.name }}</p>
        <p>{{ user.grade }}年 {{ user.major }} {{ user.group }}</p>
        <div class="text--primary">
          最近はサーバー構築やってるよ!
          <br />気軽に話しかけてね
        </div>
        <v-card-actions></v-card-actions>
      </v-card-text>
    </v-card>
  </v-hover>
</template>
<script>
export default {
  props: ["user"],
  data: () => ({}),
  methods: {},
};
</script>
<style>
.card {
  animation: anim 1s;
}

@keyframes anim {
  0% {
    transform: translateX(600px);
  }

  100% {
    transform: translateX(0px);
  }
}
</style>
