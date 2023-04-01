<script setup>
import VueTree from "@ssthouse/vue3-tree-chart";
import "@ssthouse/vue3-tree-chart/dist/vue3-tree-chart.css";
import Dialog from "./Dialog.vue"
import { ref, onMounted, computed } from "vue";

const config = {
  nodeWidth: 120,
  nodeHeight: 80,
  levelHeight: 200
}

const tree = ref(null)
const dialog = ref(null)

const nodeSelected = ref({})

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
    const { deltaY } = event

    if (deltaY === undefined) {
      return
    }

    if (deltaY > 0 ) {
      zoomIn()
    } else {
      zoomOut()
    }
  })
})

const onViewMore = (node) => {
  console.log(dialog.value.setIsOpen)
  nodeSelected.value = node
  dialog.value.setIsOpen(true)
}
</script>
<template>
  <div>
    <Dialog ref="dialog" :obj="nodeSelected" />
    <div>
      <button @click="zoomIn"> + </button>
      <button @click="zoomOut"> - </button>
    </div>
    <vue-tree ref="tree"  :dataset="props.data" :config="config" linkStyle="straight">
      <template #node="{ node, collapsed }">
        <div style="background-color: gray; border: 1px solid black; padding: 10px;">
          <p> Key: {{ node.key }} </p>
          <p> Value: {{ node.value }} </p>
          <button @click.stop="onViewMore(node)"> View more</button>
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
