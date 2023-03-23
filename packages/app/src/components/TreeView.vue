<script setup>
import VueTree from "@ssthouse/vue3-tree-chart";
import "@ssthouse/vue3-tree-chart/dist/vue3-tree-chart.css";
import { ref, onMounted } from "vue";

const config = {
  nodeWidth: 120,
  nodeHeight: 80,
  levelHeight: 200
}

const tree = ref(null)

const props = defineProps({
  data: Object | Array,
});

const zoomIn = () => {
  tree.value.zoomIn()
}
const zoomOut = () => {
  tree.value.zoomOut()
}

onMounted(() => {
  const el = document.querySelector(".tree-container")
  el.addEventListener("mousewheel", (event) => {
    console.log(event)
  })

})

</script>
<template>
  <div>
    <div>
      <button @click="zoomIn"> + </button>
      <button @click="zoomOut"> - </button>
    </div>
    <vue-tree ref="tree"  :dataset="props.data" :config="config" linkStyle="straight">
      <template #node="{ node, collapsed }">
        <div style="background-color: gray; border: 1px solid black; padding: 10px;">
          <p> Key: {{ node.key }} </p>
          <p> Value: {{ node.value }} </p>
        </div>
      </template>
    </vue-tree>
  </div>
</template>
<style lang="css">
.tree-container {
  height: 80vh;
  border: 1px solid gray;
}
</style>
