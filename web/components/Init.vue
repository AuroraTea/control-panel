<script setup lang="ts">
import { useMessage } from 'naive-ui'

window.nMessage = useMessage()

// TODO:
// get config here and block rendering, now in app.vue, but maybe also ok... naive message🤷🏻‍♂️
// const config = useConfig()
// const res = await getConfig()
// if(res.isExist) config.value = JSON.parse(res.config)
</script>

<template>
  <tempalte style="display: none" />
</template>
