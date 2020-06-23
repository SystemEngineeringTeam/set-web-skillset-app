
<template>
    <v-layout wrap> <!-- wrap属性をつけるのがポイント -->
        <v-col class="d-flex" cols="12" xs='2' sm="2" md='2'>
        <v-select
          :items="filterlist"
          label="選択"
          solo
        ></v-select>
        </v-col>
        <v-col class="d-flex" cols="12" xs='2' sm="2" md='2'>
        <v-select
          :items="filterlistdetails.grade"
          label="選択"
          solo
        ></v-select>
        </v-col>
        <v-flex class="d-flex" xs='12' sm='6' md='4'>
            <v-row>
                <div v-for="user in users" :key="user.id">
                    <v-col cols="12">
                        <v-hover
                        v-slot:default="{ hover }"
                        open-delay="200"
                        >
                            <v-card
                                :elevation="hover ? 16 : 2"
                                class="d-flex  mx-auto card"
                                max-width="400"
                            >
                                <v-responsive
                                    class="text-center grey lighten-2 rounded-circle d-inline-flex align-center justify-center ma-3"
                                    height="108"
                                    width="108"
                                    >
                                    
                                    <v-img v-bind:src='user.img'>
                                    </v-img>

                                </v-responsive>
                                <v-card-text>
                                    <p class="display-1 text--primary">
                                        {{user.name}}
                                    </p>
                                    <p>{{user.grade}}年 {{user.major}} {{user.position}}</p>
                                    <div class="text--primary">
                                        最近はサーバー構築やってるよ!<br>
                                        気軽に話しかけてね
                                    </div>
                                    <v-card-actions>                         
                                    </v-card-actions>
                                </v-card-text>
                            </v-card>
                        </v-hover>
                    </v-col>
                </div>
            </v-row>
        </v-flex>
        <!--<v-flex xs='12' sm='6' md='4'>コンテンツ</v-flex>
        <v-flex xs='12' sm='6' md='4'>コンテンツ</v-flex>
        <v-flex xs='12' sm='6' md='4'>コンテンツ</v-flex>!-->
    </v-layout>
</template>
<script>

  export default {
    data: () => ({
        
        users:[
            {id:1,img:require("../assets/logo.png"),name:"鈴木健介",grade:"3",major:"kk",position:"インフラ班"},
            {id:2,img:require("../assets/logo.png"),name:"山田太郎",grade:"3",major:"kk",position:"クリエイティブ班"},
            {id:3,img:require("../assets/logo.png"),name:"田中綾子",grade:"2",major:"kk",position:"インフラ班"},
            {id:4,img:require("../assets/logo.png"),name:"山下恵子",grade:"1",major:"kx",position:"クリエイティブ班"}
        ],
        // b:[
        //     {id:1,filter:"学年"},
        //     {id:2,filter:"専攻"},
        //     {id:3,filter:"役員"},
        //     {id:4,filter:"所属"},
        //     {id:5,filter:"役員"},
        //     {id:6,filter:"言語"},
        //     {id:7,filter:"技術"},
        // ],
        filterlist: ['選択','学年','専攻','役員','所属','言語','技術'],
        filterlistdetails:[
            {
                grade:['1','2','3','4'],
                major: ['kk','kx'],
                technology: ['C','Java'],
                position: ['クリエイティブ班','インフラ班'],
            }
        ],
    }),
  }
</script>
<style>

.card{
  animation: anim 1s;
}

@keyframes anim {
  0% {
  transform: translateX(600px);
  }

  100% {
  transform: translateX(0px);
  }
}
</style>
