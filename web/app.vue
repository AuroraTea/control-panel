<script setup>

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