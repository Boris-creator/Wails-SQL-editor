<script setup lang="ts">
import { ref } from 'vue'
import TablesList from '@/components/TablesList.vue'
import TableView from '@/components/TableView.vue'

const selectedTable = ref<string | null>(null)
</script>
<template>
  <div class="flex">
    <div class="tablesList">
      <tables-list v-model="selectedTable" />
    </div>
    <div class="tableView">
      <table-view v-if="selectedTable" :table="selectedTable" />
    </div>
  </div>
</template>
<style scoped lang="scss">
.flex {
  display: flex;
  height: 100vh;
  .tablesList {
    width: 25%;
    overflow-x: auto;
  }
  .tableView {
    width: 75%;
    background: white;
  }
}
</style>
