<template>
    <v-card
      color="light-blue darken-3"
      class="mx-auto"
      max-width="1500"
    >
        <v-container>
            <v-card
                class="text-center teal lighten-3"
                width="1200"
                height="300"
            >記事本文
            </v-card>
            <br>
            <v-card
                class="text-center light-blue lighten-3"
                width="1200"
                height="100"
            >記事への質問
            </v-card>
            <br>
            <v-btn
                @click="goHome"
                >ホームに戻る
            </v-btn>
        </v-container>
        
        
    
    
    </v-card>
</template>

<script>
  export default {
    data: () => ({
      
    }),
    methods: {
      
      goHome () {
        this.$router.push("/Home")
      }
    },

    
  }
</script>