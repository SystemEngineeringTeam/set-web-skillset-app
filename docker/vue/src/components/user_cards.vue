<template>
  <v-layout wrap>
    <!-- wrap属性をつけるのがポイント -->
    <v-flex class="d-flex" xs="12" sm="6" md="4">
      <v-row>
        <template v-if="this.$store.state.change_display === '学年'">
          <ChangeDisplayGrade />
        </template>
        <template v-else-if="this.$store.state.change_display === '専攻'">
          <ChangeDisplayMajor />
        </template>
        <template v-else-if="this.$store.state.change_display === '所属'">
          <ChangeDisplayGroup />
        </template>
        <template v-else-if="this.$store.state.change_display === '言語'">
          <ChangeDisplayTechnology />
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
import ChangeDisplayGrade from "./change_display/grade";
import ChangeDisplayMajor from "./change_display/major";
import ChangeDisplayGroup from "./change_display/group";
import ChangeDisplayTechnology from "./change_display/technology";
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
    ChangeDisplayMajor,
    ChangeDisplayGroup,
    ChangeDisplayTechnology,
    UserCard,
  },
};
</script>
<style></style>
