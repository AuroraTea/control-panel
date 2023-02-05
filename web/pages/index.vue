<script setup lang="ts">
import { setConfig } from '../apis'
const toolsResolved = {
  IPv4: resolveComponent('IPv4'),
  DeviceName: resolveComponent('DeviceName'),
}

const config = useConfig()
const toolsSelected = $ref<(keyof typeof toolsResolved)[]>(
  // @ts-ignore
  config.value.toolsSelected || [],
)

watch(
  () => toolsSelected,
  async () => {
    await setConfig({ config: JSON.stringify({ toolsSelected }) })
  },
)
</script>

<template>
  <NuxtLink to="network/ipv4">修改IP(IPv4)</NuxtLink>
  <n-dynamic-input
    v-model:value="toolsSelected"
    show-sort-button
    placeholder="Choose components"
  />
  <component
    v-for="(i, index) in toolsSelected"
    :key="index"
    :is="toolsResolved[i]"
  />
</template>

<style lang="scss" scoped></style>
