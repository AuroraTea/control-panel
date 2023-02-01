<script setup lang="ts">
import { setComputerName } from '@/apis'

const computerName = ref('')

const edit = async () => {
  const res = await setComputerName({
    newName: computerName.value,
  })
  window.nMessage.success(res)
}

</script>

<template>
  <div class="label">计算机名称</div>
  <n-input v-model:value="computerName" />
  <n-button @click="edit">修改</n-button>
</template>

<style lang="scss" scoped>
</style>