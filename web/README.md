# Control-Panel-frontend

## Setup

```shell
pnpm i
```

## Development

Start the development server on http://localhost:3000

```shell
pnpm dev
```

## Production

Generate static web and move it to backend static embed:

```shell
pnpm gen
```

## ToDo

前端现在的核心问题是纠结动态工具页如何实现:

- components + 动态组件
- pages + iframe

这决定了整个项目的结构和代码的复用方式, 以及默认值和配置项如何传递.
