<script lang="ts" setup>
import { ref, onBeforeUnmount } from 'vue'
import { CloseConnection } from '~/go/main/App'
import Connect from '@/components/Connect.vue'
import MainPage from '@/components/MainPage.vue'

const isConnected = ref(false)
onBeforeUnmount(CloseConnection)
</script>

<template>
  <div class="header" style="--wails-draggable: drag" />
  <connect v-if="!isConnected" @connect="isConnected = $event" />
  <main-page v-else />
</template>
<style scoped lang="scss">
.header {
  height: 2rem;
  background: cadetblue;
  display: none;
}
</style>
