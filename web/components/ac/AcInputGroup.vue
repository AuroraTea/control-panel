<script setup lang="ts">
const props = defineProps<{
  modelValue:(string | number)[]
}>()

type EmitType = {
  (e: 'update:modelValue'): void
}
defineEmits<EmitType>()

props.modelValue.forEach((item, index, arr) => {
  if (typeof item === 'number') arr[index] = item.toString()
})
</script>

<template>
  <n-input-group style="gap: 2px">
    <template v-for="(item, index) in modelValue" :key="index">
      <span>{{ index === 0 ? '' : '.' }}</span>
      <n-input
        :value="modelValue[index].toString()"
        @update:value="(v: string) => modelValue[index] = v"
        style="text-align: center"
        placeholder=""
        :allow-input="(v: string) => !v || /^\d+$/.test(v)"
      />
    </template>
  </n-input-group>
</template>

<style scoped lang="scss"></style>
