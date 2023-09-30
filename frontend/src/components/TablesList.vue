<script setup lang="ts">
import { ref, onBeforeMount, type PropType } from 'vue'
import { Tables } from '~/go/main/App'

defineProps({
  modelValue: {
    type: String as PropType<string | null>,
    default: null
  }
})
const emits = defineEmits<{ (e: 'update:modelValue', value: string): void }>()

const isLoading = ref(false)

const dbTables = ref<string[]>([])
async function getTables() {
  isLoading.value = true
  dbTables.value = await Tables()
  isLoading.value = false
}
onBeforeMount(getTables)
</script>
<template>
  <button :disabled="isLoading" @click="getTables">refresh</button>
  <div class="list">
    <div
      v-for="(table, i) of dbTables"
      :key="i"
      class="tableName"
      :class="{ selectedTableName: table === modelValue }"
      @click="emits('update:modelValue', table)"
    >
      {{ table }}
    </div>
  </div>
</template>
<style scoped lang="scss">
.tableName {
  cursor: pointer;
  padding: 0.2rem 0.5rem 0.2rem 1rem;
  &.selectedTableName {
    background: lightblue;
  }
}
.list {
  height: 90vh;
  overflow-y: auto;
}
</style>
