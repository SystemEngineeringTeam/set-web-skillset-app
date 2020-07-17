<template>
  <v-container>
    <!-- fluidを付けることで左右いっぱいにコンテナが広がります -->
    <!-- fill-heightを付けるとここに書いた要素は上下中央寄せで表示されます -->
    <Filters />
    <UserCards />
  </v-container>
</template>

<script>
import UserCards from "../components/user_cards";
import Filters from "../components/filter";
export default {
  components: {
    UserCards,
    Filters,
  },
  data: () => ({}),
};
</script>
