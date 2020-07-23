<template>
  <v-layout wrap>
    <!-- wrap属性をつけるのがポイント -->
    <v-flex class="d-flex" xs="12" sm="6" md="4">
      <v-row>
        <div v-for="(group, i) in $store.state.groups" :key="i">
          <p>{{ group }}</p>
          <v-sheet class="d-flex" color="teal lighten-3" height="300" width="2000">
            <div v-for="user in $store.state.users" :key="user.id">
              <template v-if="user.group === group">
                <UserCard :user="user"></UserCard>
              </template>
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