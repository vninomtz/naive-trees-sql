<script setup lang="ts">
import VueTree from "@ssthouse/vue3-tree-chart";
import "@ssthouse/vue3-tree-chart/dist/vue3-tree-chart.css";
import { onMounted, ref } from 'vue';

const treeProps = {
  label: 'value',
  children: 'children',
  key: 'key'
}


const treeConfig = {
  nodeWidth: 120,
  nodeHeight: 80,
  levelHeight: 200
}

const dataTree = ref({})
const loading = ref(true)

const tree = ref(null)

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

const getChildrenOf = (n: any) => {
  console.log(n)
}
onMounted(() => {
  fetchNodes()
})

const zoomIn = () => {
  tree.value?.zoomIn()
}
const zoomOut = () => {
  tree.value?.zoomOut()
}

</script>

<template>
  <header>
    <h1> Tree Viewer </h1>
  </header>
  <main>
    <div>
      <blocks-tree v-if="!loading" :data="dataTree" :collapsable="false" :props="treeProps">
        <template #node="{ data, context }">
          <span style="color: black">
            {{ data.value }}
          </span>
          <br />
          <div v-if="data.children && data.children.length == 0">
            <a href="#" @click="getChildrenOf(data)"> Show more </a>
          </div>
        </template>
      </blocks-tree>

      <p v-if="loading">
        Cargando ...
      </p>
    </div>
    <div class="container">
      <button @click="zoomIn"> zoom in</button>
      <button @click="zoomOut"> zoom out</button>
      <vue-tree v-if="!loading" ref="tree" style="width: 800px; height: 600px; border: 1px solid gray;"
        :dataset="dataTree" :config="treeConfig" />
    </div>
  </main>
</template>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
