// @ts-check
const { defineConfig } = require('eslint-define-config')

module.exports = defineConfig({
  root: true,
  env: {
    browser: true,
    es2021: true
  },
  extends: [
    '@nuxtjs/eslint-config-typescript',
    'plugin:vue/vue3-recommended',
    'prettier',
    'plugin:prettier/recommended',
  ],
  parser: 'vue-eslint-parser',
  parserOptions: {
    parser: '@typescript-eslint/parser',
    ecmaVersion: 'latest',
    sourceType: 'module',
    jsxPragma: 'React',
    ecmaFeatures: {
      jsx: true
    }
  },
  plugins: [
    'vue', '@typescript-eslint'
  ],
  rules: {
    'vue/no-multiple-template-root': 'off', // 允许 vue3 之后 template 下有多个组件
    'vue/script-setup-uses-vars': 'error',
    'vue/no-reserved-component-names': 'off',
    '@typescript-eslint/ban-ts-ignore': 'off',
    '@typescript-eslint/explicit-function-return-type': 'off',
    '@typescript-eslint/no-explicit-any': 'off',
    '@typescript-eslint/no-var-requires': 'off',
    '@typescript-eslint/no-empty-function': 'off',
    'vue/custom-event-name-casing': 'off',
    'no-use-before-define': 'off',
    '@typescript-eslint/no-use-before-define': 'off',
    '@typescript-eslint/ban-ts-comment': 'off',
    '@typescript-eslint/ban-types': 'off',
    '@typescript-eslint/no-non-null-assertion': 'off',
    '@typescript-eslint/explicit-module-boundary-types': 'off',
    '@typescript-eslint/no-unused-vars': 'off',
    'no-unused-vars': 'off',
    'space-before-function-paren': 'off',
    'comma-dangle': ['error', 'always-multiline'],
    quotes: ['warn', 'single'],
    curly: ['error', 'multi-line', 'consistent'],
    semi: ['error', 'never'],
    'space-before-blocks': ['error', 'always'],
    'spaced-comment': [
      'warn',
      'always',
      {
        exceptions: ['-'],
      },
    ],
    'no-var': 'error',
    indent: [
      'error',
      2,
      {
        ArrayExpression: 1,
        ObjectExpression: 1,
        ImportDeclaration: 1,
        SwitchCase: 1,
        ignoreComments: true,
      },
    ],
    'prefer-promise-reject-errors': 'off',
    'vue/attributes-order': 'off',
    'vue/one-component-per-file': 'off',
    'vue/html-closing-bracket-newline': 'off',
    'vue/max-attributes-per-line': 'off',
    'vue/multiline-html-element-content-newline': 'off',
    'vue/singleline-html-element-content-newline': 'off',
    'vue/attribute-hyphenation': 'off',
    'vue/require-default-prop': 'off',
    'vue/require-explicit-emits': 'off',
    'vue/no-mutating-props': 'off',
    'vue/html-self-closing': [
      'error',
      {
        html: {
          void: 'always',
          normal: 'never',
          component: 'always',
        },
        svg: 'always',
        math: 'always',
      },
    ],
    'vue/multi-word-component-names': 'off',
    'vue/no-v-html': 'off',
    '@typescript-eslint/no-this-alias': 'off',
  },
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
})