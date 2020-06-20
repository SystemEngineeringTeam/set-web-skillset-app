<template>
    <v-container > <!-- fluidを付けることで左右いっぱいにコンテナが広がります -->
        <!-- fill-heightを付けるとここに書いた要素は上下中央寄せで表示されます -->
        <UserCard/>
    </v-container>
</template>

<script>

import UserCard from '../components/user_card';
  export default {
    components: {
        UserCard,
    },
    data: () => ({

    }),
  }
</script>