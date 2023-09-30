<script lang="ts" setup>
import { reactive, ref, computed } from 'vue'
import { Connect } from '~/go/main/App'
import { main } from '~/go/models'

const emits = defineEmits<{ (e: 'connect', value: boolean): void }>()

const loginData = reactive({
  dbPassword: '',
  dbPort: '127.0.0.1:3306',
  dbName: ''
})
const isValid = computed(() => loginData.dbName && loginData.dbPassword && loginData.dbPort)
const connectionError = ref<main.DBConnection['err'] | null>(null)

async function connect() {
  if (!isValid.value) {
    return
  }
  const { success, err } = await Connect(loginData.dbPassword, loginData.dbPort, loginData.dbName)
  connectionError.value = err
  emits('connect', success)
}
</script>

<template>
  <main>
    <div class="form">
      <input v-model="loginData.dbPort" type="text" placeholder="port" />
      <input v-model="loginData.dbName" type="text" placeholder="DB" />
      <input
        v-model="loginData.dbPassword"
        autocomplete="off"
        type="password"
        placeholder="password"
      />
      <p v-if="connectionError" class="form__errorMessage">{{ connectionError }}</p>
      <button :disabled="!isValid" @click="connect">Connect</button>
    </div>
  </main>
</template>
<style scoped lang="scss">
main {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
.form {
  width: 40vw;
  max-width: 400px;
  &__errorMessage {
    color: darkred;
  }
}
input,
button {
  display: block;
  width: 100%;
  margin-top: 1rem;
  padding: 1rem 0.5rem;
}
</style>
