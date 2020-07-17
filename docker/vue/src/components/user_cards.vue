<template>
  <v-layout wrap>
    <!-- wrap属性をつけるのがポイント -->
    <v-flex class="d-flex" xs="12" sm="6" md="4">
      <v-row>
        <template v-if="this.$store.state.change_display === '学年'">
          <ChangeDisplayGrade />
        </template>
        <template v-else>
          <div v-for="user in this.$store.state.users" :key="user.id">
            <v-col cols="12">
              <UserCard :user="user"></UserCard>
            </v-col>
          </div>
        </template>
      </v-row>
    </v-flex>
    <!--<v-flex xs='12' sm='6' md='4'>コンテンツ</v-flex>
        <v-flex xs='12' sm='6' md='4'>コンテンツ</v-flex>
    <v-flex xs='12' sm='6' md='4'>コンテンツ</v-flex>!-->
  </v-layout>
</template>
<script>
import ChangeDisplayGrade from "./change_display_grade";
import UserCard from "./user_card";

export default {
  data: () => ({}),
  methods: {
    reset_filter(filter) {
      if (filter === "選択") {
        this.filter_grade = "";
      }
    },
  },
  components: {
    ChangeDisplayGrade,
    UserCard,
  },
};
</script>
<style></style>
