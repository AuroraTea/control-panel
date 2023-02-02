<script setup lang="ts">
import { getNetAdapters, setIPv4 } from '@/apis'

// defineProps(['adapter', 'ip', 'mask', 'gateway'])

const options = ref([{ value: '0', label: '无法获取网卡列表' }])
const netAdapter = $ref('')

const config = useConfig()
// TODO: default value

const ip = $ref(['', '', '', ''])
const mask = $ref(['', '', '', ''])
const gateway = $ref(['', '', '', ''])

const edit = async () => {
  const res = await setIPv4({
    netAdapter,
    ip: ip.join('.'),
    mask: mask.join('.'),
    gateway: gateway.join('.'),
  })
  window.nMessage.success(res)
}

const refreshOptions = async () => {
  console.log('refreshOptions')
  const res = await getNetAdapters()
  options.value = res.map(item => ({
    value: item,
    label: item,
  }))
}
</script>

<template>
  <div class="ipv4-layout">
    <div class="label">网卡名称</div>
    <n-select
v-model:value="netAdapter"
:options="options"
@click="refreshOptions" />
    <div class="label">IP地址</div>
    <AcInputGroup v-model="ip" />
    <div class="label">子网掩码</div>
    <AcInputGroup v-model="mask" />
    <div class="label">网关地址</div>
    <AcInputGroup v-model="gateway" />
    <div class="footer">
      <n-button type="primary" @click="edit">修改</n-button>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.ipv4-layout {
  padding: 16px;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  flex-direction: column;

  .label {
    padding: 8px 0 0 0;
  }

  .footer {
    padding: 16px 0;
    align-self: end;
  }
}
</style>
