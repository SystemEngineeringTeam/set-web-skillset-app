<template>
        <v-layout wrap > 
                <v-col class="d-flex" cols="12" xs='2' sm="2" md='2'>
                        <v-select
                        v-model="$store.state.filter"
                        :items="$store.state.filterlist"
                        label="選択"
                        solo
                        ></v-select>
                </v-col>
                <v-col class="d-flex" cols="12" xs='2' sm="2" md='2' v-for="(filter_item,i) in filter_items" :key="i">
                        <v-select
                        :items="filter_item.filter"
                        :label="filter_item.title"
                        solo
                        ></v-select>
                </v-col>
        </v-layout>
</template>
<script>
export default {
  data: () => ({
          cnt:0,
          filter_items:[]
  }),
  created(){
          this.filter_add()
  },
  methods: {
  filter_add(){
        this.filter_items.push({title:"学年選択",filter:this.$store.state.grade},
                                {title:"専攻選択",filter:this.$store.state.major},
                                {title:"所属選択",filter:this.$store.state.group},
                                {title:"言語選択",filter:this.$store.state.technology},
                                {title:"技術選択",filter:this.$store.state.technical_area});
        },
//       filter(filters) {
//         for(var i = 0; i < filters.length; i++){
//                 if(filters[i] === '学年')
//                 {
//                         this.$store.state.show_filter.grade = true
//                 }
//                 else if(filters[i] === '専攻')
//                 {
//                         this.$store.state.show_filter.major = true
//                 }
//                 else if(filters[i] === '役員')
//                 {
//                         this.$store.state.show_filter.position = true
//                 }
//                 else if(filters[i] === '所属')
//                 {
//                         this.$store.state.show_filter.group = true       
//                 }
//         }
// if(filters.indexOf("学年") == -1)
// {
//         this.$store.state.show_filter.grade = false

// }
// if(filters.indexOf("専攻") == -1)
// {
//         this.$store.state.show_filter.major = false
// }
//       }
    },
};
</script>