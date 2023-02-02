import { useMessage } from 'naive-ui'

declare global {
  interface Window {
    nMessage: ReturnType<typeof useMessage>
  }
}
