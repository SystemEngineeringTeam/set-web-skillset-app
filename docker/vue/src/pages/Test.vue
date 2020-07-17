<template>
  <div class="test">
    test
    <div class="box"></div>
    <v-sheet class="d-flex" color="teal lighten-3" height="300">
      <sheet-footer>#4: (2r x 1c)</sheet-footer> </v-sheet
    >aa
  </div>
</template>

<script>
export default {
  data: () => ({}),
};
</script>
<style>
.box {
  width: 100px;
  height: 100px;
  background: #0091ea;
  animation: anim 1s;
}

@keyframes anim {
  0% {
    transform: translateX(300px);
  }

  100% {
    transform: translateX(0px);
  }
}
</style>
