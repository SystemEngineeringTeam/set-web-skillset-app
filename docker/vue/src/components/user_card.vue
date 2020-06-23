
<template>
    <v-layout wrap> <!-- wrap属性をつけるのがポイント -->
        <v-flex class="d-flex" xs='12' sm='6' md='4'>
            <v-row >
                <div v-for="user in $store.state.users" :key="user.id">
                    <v-col cols="12">
                        <v-hover
                        v-slot:default="{ hover }"
                        open-delay="200"
                        >
                            <v-card
                                v-if="$store.state.filter_grade === user.grade || !$store.state.filter_grade && !$store.state.filter_major || $store.state.filter === '選択' || $store.state.filter_major === user.major "
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
                                    <p>{{user.grade}}年 {{user.major}} {{user.group}}</p>
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
    
        // b:[
        //     {id:1,filter:"学年"},
        //     {id:2,filter:"専攻"},
        //     {id:3,filter:"役員"},
        //     {id:4,filter:"所属"},
        //     {id:5,filter:"役員"},
        //     {id:6,filter:"言語"},
        //     {id:7,filter:"技術"},
        // ],
        
    }),
    methods: {
        reset_filter(filter){
            if(filter === '選択'){
                this.filter_grade = ""
            }
        }
    }
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
