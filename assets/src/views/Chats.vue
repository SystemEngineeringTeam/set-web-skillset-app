<template>
    <v-card
      color="light-blue darken-3"
      class="mx-auto"
      max-width="1500"
    >
        <v-container>

            <v-row
            v-for="n in 2"
            :key="n"
            :class="n === 1 ? 'mb-6' : ''"
            no-gutters
            >
                <v-col
                v-for="k in 4"
                :key="k"
                >
                    <v-card
                        class="text-center"
                        width="250"
                        height="300"
                    >
                        <div v-if="n==1&&k==1">
                            <router-link to="/chat">
                                全体チャット
                            </router-link>
                        </div>
                        <div v-else-if="n==1&&k==2">
                            <router-link to="/javachat">
                                Javaチャット
                            </router-link>
                        </div>
                        <div v-else-if="n==1&&k==3">
                            <router-link to="/individualchat">
                                C言語チャット
                            </router-link>
                        </div>
                        <div v-else-if="n==1&&k==4">
                            <router-link to="/individualchat">
                                C++チャット
                            </router-link>
                        </div>
                        <div v-else-if="n==2&&k==1">
                            <router-link to="/individualchat">
                                Pythonチャット
                            </router-link>
                        </div>
                        <div v-else-if="n==2&&k==2">
                            <router-link to="/individualchat">
                                JavaScriptチャット
                            </router-link>
                        </div>
                        <div v-else-if="n==2&&k==3">
                            <router-link to="/individualchat">
                                C#チャット
                            </router-link>
                        </div>
                        <div v-else-if="n==2&&k==4">
                            <router-link to="/individualchat">
                                Goチャット
                            </router-link>
                        </div>
                    </v-card>
                </v-col>
            </v-row>
                
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