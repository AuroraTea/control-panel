<script setup lang="ts">
import { getConfig } from './apis';
const config = useConfig()
try {
  const res = await getConfig()
  if (res.isExist && res.config) config.value = JSON.parse(res.config)
} catch (error) {
  // @ts-ignore
  error.stopPreErr()
  confirm('Backend Error')
}
</script>


<template>
  <n-message-provider>
    <Init/>
    <div class="app">
      <NuxtPage />
    </div>
  </n-message-provider>
</template>

<style lang="scss" scoped>
.app{
  width: 375px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  .label{
    width: 80px;
  }
}
</style>