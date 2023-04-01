<script setup lang="ts">
import { onMounted } from 'vue';

onMounted(() => {
  const canvas = document.getElementById("scene") as HTMLCanvasElement | null;
  const ctx = canvas.getContext("2d")

  console.log(ctx)
  //draw(ctx)
  sqr(ctx)

})

function draw(ctx: CanvasRenderingContext2D) {
  ctx.fillStyle = "rgb(200, 0, 0)";
  ctx.fillRect(10, 10, 50, 50);

  ctx.fillStyle = "rgba(0, 0, 200, 0.5)";
  ctx.fillRect(30, 30, 50, 50);
}

function sqr(ctx: CanvasRenderingContext2D) {
  ctx.fillRect(25, 25, 100, 100);
  ctx.clearRect(45, 45, 60, 60)
  ctx.strokeRect(50, 50, 50, 50);
}

</script>
<template>
  <canvas width="500" height="500" id="scene">
    Tree view
  </canvas>
</template>
<style lang="css">
canvas {
  border: 1px solid gray;
}
</style>
