<script setup>
defineProps(['modelValue'])
defineEmits(['update:modelValue'])

</script>

<template>
  <n-input-group style="gap: 2px;">
    <template v-for="(item, index) in modelValue">
      {{ index==0 ? '' : '.' }}<n-input
        :value="modelValue[index]"
        @update:value = "v => modelValue[index] = v"
        style="text-align:center"
        placeholder=""
        :allow-input="v => !v || /^\d+$/.test(v)" />
    </template>

  </n-input-group>
</template>

<style scoped lang="scss">

</style>