import Components from 'unplugin-vue-components/vite'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  telemetry: false,
  ssr: false,

  // components: {
  //   global: true,
  //   // dirs: ['~/components/cp'] //无效?
  // },

  vite: {
    plugins: [
      Components({
        resolvers: [NaiveUiResolver()]
      }),
    ],
    optimizeDeps: {
      include:
        process.env.NODE_ENV === 'development'
          ? ['naive-ui', 'vueuc', 'date-fns-tz/esm/formatInTimeZone']
          : [],
    },
  },

  experimental: {
    reactivityTransform: true, // macro
  },
})
