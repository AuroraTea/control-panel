<script setup>
const toolsResolved = {
  IPv4 : resolveComponent('IPv4'),
  DeviceName : resolveComponent('DeviceName'),
}
const toolsSelected = $ref([
  'IPv4',
  'DeviceName'
])
</script>

<template>
  <NuxtLink to="network/ipv4">修改IP(IPv4)</NuxtLink>
  <n-dynamic-input
    v-model:value="toolsSelected"
    show-sort-button
    placeholder="Choose components"
  />
  <template v-for="i in toolsSelected">
    <component :is="toolsResolved[i]"/>
  </template>
</template>

<style lang="scss" scoped>

</style>