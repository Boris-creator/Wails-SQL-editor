<script setup lang="ts">
import { ref, onBeforeMount, watchEffect } from 'vue'
import { Table } from '~/go/main/App'
import { main } from '~/go/models'

const props = defineProps({
  table: {
    type: String,
    required: true
  }
})

const tableRows = ref<Record<string, string>[]>([])
const tableColumns = ref<string[]>([])
const tableColumnTypes = ref<main.ColumnType[]>([])

const displayMode = ref<'data' | 'meta'>('data')

const isLoading = ref(false)

async function selectAll() {
  isLoading.value = true
  try {
    const rows = await Table(props.table)
    tableColumns.value = rows.columns
    tableRows.value = (rows.data ?? []).map((row: string[]) =>
      Object.fromEntries(row.map((item, i) => [tableColumns.value[i], item]))
    )
    tableColumnTypes.value = rows.columnTypes
  } finally {
    isLoading.value = false
  }
}

onBeforeMount(selectAll)
watchEffect(selectAll)
</script>
<template>
  <div class="relative">
    <div class="tools">
      <button :disabled="isLoading" @click="selectAll">refresh</button>
      <select v-model="displayMode">
        <option value="data">data</option>
        <option value="meta">column types</option>
      </select>
    </div>
    <div class="wrapper">
      <table>
        <thead>
          <tr>
            <th v-for="col of tableColumns" :key="col">{{ col }}</th>
          </tr>
        </thead>
        <tbody v-if="displayMode === 'data'">
          <tr v-for="(row, i) of tableRows" :key="i">
            <td v-for="(value, col) in row" :key="col">
              {{ value }}
            </td>
          </tr>
        </tbody>
        <tbody v-else>
          <tr>
            <td v-for="(col, i) of tableColumnTypes" :key="i">
              {{ col.type }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<style scoped lang="scss">
.tools {
  display: flex;
  gap: 2rem;
}
.wrapper {
  width: 100%;
  overflow-x: auto;
}
</style>
