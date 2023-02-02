module.exports = {
  env: {
    browser: true,
    es2021: true
  },
  extends: ['@nuxtjs/eslint-config-typescript', 'plugin:vue/vue3-recommended'],
  parserOptions: {
    ecmaVersion: 13,
    sourceType: 'module'
  },
  plugins: [],
  rules: {},
  overrides: [
    {
      files: [
        '**/apis/**/*.{js,ts,vue}',
        '**/components/**/*.{js,ts,vue}',
        '**/composables/**/*.{js,ts,vue}',
        '**/pages/**/*.{js,ts,vue}',
        '**/app.{js,ts,vue}',
      ],
      rules: {
        'vue/multi-word-component-names': 'off'
      }
    }
  ]
}