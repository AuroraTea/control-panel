<script setup lang="ts">
import { putConfig} from '../apis'
const toolsResolved = {
  IPv4: resolveComponent('IPv4'),
  DeviceName: resolveComponent('DeviceName'),
}

const config = useConfig()
// @ts-ignore
const toolsSelected = $ref<Array<keyof typeof toolsResolved>>(config.value.toolsSelected || [])

watch(
  () => toolsSelected,
  async () => {
    console.log(toolsSelected)
    await putConfig({ config: JSON.stringify({ toolsSelected: toolsSelected }) })
  })
</script>

<template>
  <NuxtLink to="network/ipv4">修改IP(IPv4)</NuxtLink>
  <n-dynamic-input v-model:value="toolsSelected" show-sort-button placeholder="Choose components" />
  <template v-for="i in toolsSelected">
    <component :is="(toolsResolved[i])" />
  </template>
</template>

<style lang="scss" scoped>

</style>