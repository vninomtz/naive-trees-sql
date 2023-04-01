<script setup lang="ts">
import TreeView from './components/TreeView.vue'
import TreeCanvas from './components/TreeCanvas.vue'
import { onMounted, ref } from 'vue';

const dataTree = ref({})
const loading = ref(false)


const fetchNodes = async () => {
  try {
    loading.value = true
    const req = await fetch("http://localhost:8000/api/nodes")
    const res = await req.json()
    dataTree.value = res
    loading.value = false

  } catch (error) {
    console.log(error)
  }
}

onMounted(() => {
  //fetchNodes()
})

</script>

<template>
  <header>
    <h1> Tree Viewer </h1>
  </header>
  <main>
    <div>
      <TreeCanvas />
    </div>
    <div>
      <!-- <TreeView :data="dataTree"> -->
      <!-- </TreeView> -->
    </div>
  </main>
</template>

<style scoped>
</style>
