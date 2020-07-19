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
        :label="filter_item.title"
        :items="filter_item.filter"
        @change="onChange(filter_item.title, $event)"
        solo
      ></v-select>
    </v-col>
  </v-layout>
</template>
<script>
import { mapMutations } from "vuex";

export default {
  data: () => ({
    filter_data: [],
    filter_items: [],
  }),
  created() {
    this.filter_add();
  },
  methods: {
    onChange(title, event) {
      if (title === "学年選択") {
        this.update_filter_grade(event);
      } else if (title === "専攻選択") {
        this.update_filter_major(event);
      } else if (title === "所属選択") {
        this.update_filter_group(event);
      } else if (title === "言語選択") {
        this.update_filter_technology(event);
      } else if (title === "技術選択") {
        this.update_filter_technical_area(event);
      }
    },
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
    ...mapMutations([
      "update_filter_grade",
      "update_filter_major",
      "update_filter_group",
      "update_filter_technology",
      "update_filter_technical_area",
    ]),
  },
};
</script>
