<script setup>
import { ref } from 'vue'
import {
  Dialog,
  DialogPanel,
  DialogTitle,
  DialogDescription
} from '@headlessui/vue'

const props = defineProps({
  obj: Object
});

const isOpen = ref(false)

function setIsOpen(value) {
  isOpen.value = value;
}

defineExpose({
  setIsOpen
})
</script>
<template>
  <Dialog :open="isOpen" @close="setIsOpen" class="dialog-container">
    <div class="dialog-panel-container">
    <DialogPanel class="dialog-panel">
      <DialogTitle> Test {{obj.value}} </DialogTitle>
      <DialogDescription>
        Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.
      </DialogDescription>

      <p>
        Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem
        pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud
        ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut
        officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris
        cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia
        eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat
        ullamco ut ea consectetur et est culpa et culpa duis.
      </p>

      <button @click="setIsOpen(false)">Close</button>
    </DialogPanel>
    </div>
  </Dialog>
</template>
<style lang="css">
.dialog-container {
  position: relative;
  z-index: 50;
}
.dialog-panel-container {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  overflow-y: auto;
  display: flex;
  justify-items: center;
  padding: 4px;
}
.dialog-panel {
  width: 100%;
  border-radius: 15px;
  background-color: gray;
  max-width: 28rem;
  margin: auto auto;
}
</style>
