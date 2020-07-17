<template>
  <v-layout wrap>
    <v-col class="d-flex" cols="12" xs="2" sm="2" md="2">
      <v-select
        v-model="$store.state.change_display"
        :items="$store.state.filterlist"
        label="選択"
        solo
      ></v-select>
    </v-col>
    <v-col
      class="d-flex"
      cols="12"
      xs="2"
      sm="2"
      md="2"
      v-for="(filter_item, i) in filter_items"
      :key="i"
    >
      <v-select
        v-model="filter_data[i]"
        :items="filter_item.filter"
        :label="filter_item.title"
        solo
      ></v-select>
    </v-col>
  </v-layout>
</template>
<script>
// import { mapMutations } from "vuex";

export default {
  data: () => ({
    filter_data: [],
    filter_items: [],
  }),
  created() {
    this.filter_add();
  },
  methods: {
    filter_add() {
      this.filter_items.push(
        {
          title: "学年選択",
          filter: this.$store.state.grades,
        },
        {
          title: "専攻選択",
          filter: this.$store.state.major,
        },
        {
          title: "所属選択",
          filter: this.$store.state.group,
        },
        {
          title: "言語選択",
          filter: this.$store.state.technology,
        },
        {
          title: "技術選択",
          filter: this.$store.state.technical_area,
        }
      );
    },
    //   filter_add(){
    //         this.filter_items.push({title:"学年選択",filter:this.$store.state.grade,set_filter:this.update_filter_grade(this.filter_data[0])},
    //                                 {title:"専攻選択",filter:this.$store.state.major,set_filter:this.update_filter_major(this.filter_data[1])},
    //                                 {title:"所属選択",filter:this.$store.state.group,set_filter:this.update_filter_group(this.filter_data[2])},
    //                                 {title:"言語選択",filter:this.$store.state.technology,set_filter:this.update_filter_technology(this.filter_data[3])},
    //                                 {title:"技術選択",filter:this.$store.state.technical_area,set_filter:this.update_filter_technical_area(this.filter_data[4])});
    //         },
    //     ...mapMutations([
    //       "update_filter_grade",
    //       "update_filter_major",
    //       "update_filter_group",
    //       "update_filter_technology",
    //       "update_filter_technical_area"
    //     ])
  },
};
</script>
