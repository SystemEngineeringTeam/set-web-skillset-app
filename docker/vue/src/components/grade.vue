<template>
  <v-layout wrap>
    <!-- wrap属性をつけるのがポイント -->
    <v-flex class="d-flex" xs="12" sm="6" md="4">
      <v-row>
        <div v-for="(grade, i) in $store.state.grades" :key="i">
          <p>{{ i + 1 }}年生</p>
          <v-sheet
            class="d-flex"
            color="teal lighten-3"
            height="300"
            width="2000"
          >
            <div v-for="user in $store.state.users" :key="user.id">
              <v-hover v-slot:default="{ hover }" open-delay="200">
                <v-card
                  v-if="user.grade === grade"
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
            </div>
          </v-sheet>
        </div>
      </v-row>
    </v-flex>
  </v-layout>
</template>
<script>
export default {
  data: () => ({}),
  methods: {},
};
</script>
