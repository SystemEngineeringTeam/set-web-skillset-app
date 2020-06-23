<template>
        <v-layout wrap > 
                <v-col class="d-flex" cols="12" xs='3' sm="3" md='3'>
                        <v-select
                        v-model="$store.state.filter"
                        :items="$store.state.filterlist"
                        label="選択"
                        multiple
                        chips
                        >{{filter($store.state.filter)}}</v-select>
                </v-col>
                <v-col class="d-flex" cols="12" xs='2' sm="2" md='2'>
                        <v-select
                        v-if="this.$store.state.show_filter.grade"
                        v-model="$store.state.filter_grade"
                        :items="$store.state.grade"
                        label="学年選択"
                        solo
                        ></v-select>
                </v-col>
                <v-col class="d-flex" cols="12" xs='2' sm="2" md='2'>
                        <v-select
                        v-if="this.$store.state.show_filter.major"
                        v-model="$store.state.filter_major"
                        :items="$store.state.major"
                        label="専攻選択"
                        solo
                        ></v-select>
                </v-col>
                <!--{{this.$store.state.filter}}
                {{this.$store.state.show_filter.grade}}
                {{this.$store.state.filter_grade}}-->
        </v-layout>
</template>
<script>
export default {
  data: () => ({
          cnt:0
  }),
    methods: {
      filter(filters) {
        for(var i = 0; i < filters.length; i++){
                if(filters[i] === '学年')
                {
                        this.$store.state.show_filter.grade = true
                }
                else if(filters[i] === '専攻')
                {
                        this.$store.state.show_filter.major = true
                }
                else if(filters[i] === '役員')
                {
                        this.$store.state.show_filter.position = true
                }
                else if(filters[i] === '所属')
                {
                        this.$store.state.show_filter.group = true       
                }
        }
        if(filters.indexOf("学年") == -1)
        {
                this.$store.state.show_filter.grade = false
                this.$store.state.filter_grade = ''
        }
        if(filters.indexOf("専攻") == -1)
        {
                this.$store.state.show_filter.major = false
                this.$store.state.filter_major = ''
        }
      }
    },
};
</script>