<template>
    <div class="test">
        test
        <div class="box"></div>
    </div>
</template>

<script>

  export default {
    data: () => ({

    }),
  }
</script>
<style>

.box{
  width:100px;
  height:100px;
  background:#0091EA;
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