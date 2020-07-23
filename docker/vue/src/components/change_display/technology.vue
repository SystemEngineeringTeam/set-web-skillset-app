<template>
  <v-layout wrap>
    <!-- wrap属性をつけるのがポイント -->
    <v-flex class="d-flex" xs="12" sm="6" md="4">
      <v-row>
        <div v-for="(technology, i) in $store.state.technologys" :key="i">
          <p>{{ technology }}</p>
          <v-sheet class="d-flex" color="teal lighten-3" height="300" width="2000">
            <div v-for="user in $store.state.users" :key="user.id">
              <div v-for="(user_technology,i) in user.technologys" :key="i">
                <template v-if="user_technology === technology">
                  <UserCard :user="user"></UserCard>
                </template>
              </div>
            </div>
          </v-sheet>
        </div>
      </v-row>
    </v-flex>
  </v-layout>
</template>

<script>
import UserCard from ".././user_card";
export default {
  data: () => ({}),
  methods: {},
  components: {
    UserCard,
  },
};
</script>