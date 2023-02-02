<script setup lang="ts">
import { defineProps } from 'vue'

const props = defineProps<{
  modelValue: Array<string | number>
}>()

defineEmits(['update:modelValue'])

props.modelValue.forEach((item, index, arr) => {
  if (typeof item === 'number') arr[index] = item.toString()
})
</script>

<template>
  <n-input-group style="gap: 2px">
    <template v-for="(item, index) in modelValue" :key="index">
      {{ index == 0 ? '' : '.' }}
      <n-input
        :value="modelValue[index]"
        @update:value="(v:string) => modelValue[index] = v"
        style="text-align: center"
        placeholder=""
        :allow-input="(v: string) => !v || /^\d+$/.test(v)"
      />
    </template>
  </n-input-group>
</template>

<style scoped lang="scss"></style>
