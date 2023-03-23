<script setup>
import { ref } from 'vue'
import {
  Listbox,
  ListboxButton,
  ListboxOptions,
  ListboxOption
} from '@headlessui/vue'

const people = [
  { id: 1, name: 'Victor', unavailable: false},
  { id: 2, name: 'Alan', unavailable: false},
  { id: 3, name: 'Carlos', unavailable: true},
]

const selectedPerson = ref(people[0])
</script>
<template>
  <Listbox v-model="selectedPerson">
    <ListboxButton>{{ selectedPerson.name }}</ListboxButton>
    <ListboxOptions>
      <ListboxOption
        v-for="person in people"
        :key="person.id"
        :value="person"
        :disabled="person.unavailable"
      >
        {{ person.name }}
      </ListboxOption>
    </ListboxOptions>
  </Listbox>
</template>
