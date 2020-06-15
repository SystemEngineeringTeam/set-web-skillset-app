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
            >
              <v-card>Javaチャット</v-card>
            </v-card>
              
            
            <br>
            <v-btn
                @click="goChats"
                >チャット一覧に戻る
            </v-btn>
        </v-container>
        
        
    
    
    </v-card>
</template>

<script>
  export default {
    data: () => ({
      
    }),
    methods: {
      
      goChats () {
        this.$router.push("/chats")
      }
    },

    
  }
</script>