<template>
  
    <v-card
      color="light-blue darken-3"
      class="mx-auto"
      max-width="1500"
    >

      <v-container fluid>
       
        
  
        <!-- <div class="text-center"> -->
          <v-row dense>
            <v-card
              class="text-center teal lighten-3"
              width="400"
              height="300"
              >
              <div>
                <v-text>投稿記事</v-text>
              </div>
              
              <router-link to="/article">
                aaa
                <br>
                sss
              </router-link>

              <div>
                <v-btn
                  @click="goArticles"
                  >投稿記事一覧
                </v-btn>
              </div>
            </v-card>

            <v-spacer></v-spacer>

            <v-card
              class="text-center amber lighten-3"
              width="400"
              height="300"
            >知っておいてお得なツール紹介</v-card>

            <v-spacer></v-spacer>

            <v-card
              class="text-center light-blue lighten-3"
              width="400"
              height="300"
            >
              <div>
                <v-text>全体チャット</v-text>
              </div>
              <div>
              <br>
                <v-btn
                  @click="goChats"
                  >チャット一覧へ
                </v-btn>
              </div>
            </v-card>
              
            </v-row>
            <br>
          <!-- </div> -->
      <!-- </v-container> -->
          
      <!-- <v-container fluid> -->
        <v-row dense>

          <v-text>ユーザ検索</v-text>
          <v-col cols="12" sm="2">
            <v-select
              :items="language"
              label="言語"
              solo
            >
              <option>Java</option>
            </v-select>
          </v-col>
          <v-col cols="12" sm="2">  
            <v-select
              :items="development"
              label="開発"
              solo
            ></v-select>
          </v-col>
        </v-row>

      <v-row
        v-for="n in 1"
        :key="n"
        :class="n === 1 ? 'mb-6' : ''"
        no-gutters
      >
        <v-col
          v-for="user in user"
          :key="user.name"
        >
          <v-card
            class="ma-2 pa-2 text-center"
            outlined
            tile
          >
            
            <v-btn @click="goPersonalChat" color="warning" fab large dark>
              <v-icon>mdi-account-circle</v-icon>
            </v-btn>
            
            <br>
            <v-text>ユーザ名: {{ user.name }}</v-text>
            <br>
            <v-text>一言コメント</v-text>
            <v-spacer></v-spacer>
            <v-text>{{ user.comment }}</v-text>
          </v-card>
        </v-col>
      </v-row>
    
    </v-container>



    </v-card>
  
</template>

<script>
  export default {
    data: () => ({
      language: ['Java', 'C言語', 'C++', 'Python', 'JavaScript', 'C#', 'Go'],
      development: ['a', 's'],
      user: [
        { name: 'a',comment: 'z' },
        { name: 's', comment: 'x' },
        { name: 'd', comment: 'c' },
        { name: 'f', comment: 'v' },
        { name: 'g', comment: 'b' },
        ],
    }),

    methods: {
      logout () {
          // this.$refs.form.logout()
        this.$router.push("/Login")
      },
      goArticles () {
        this.$router.push("/Articles")
      },
      goChats () {
        this.$router.push("/Chats")
      },
      goPersonalChat () {
        this.$router.push("/PersonalChat")
      }
    },
  }
</script>

