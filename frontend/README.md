# Link Tree ~ Frontend Section

## Directory structure:
```ts
.
├── components.json
├── cypress // e2e test using cypress
├── cypress.config.ts
├── Dockerfile.vue
├── index.html
├── nginx // nginx server for docker build
├── package.json
├── package-lock.json
├── public
├── README.md
├── src // source code
├── tailwind.config.js
├── tsconfig.app.json
├── tsconfig.json
├── tsconfig.node.json
├── tsconfig.vitest.json
├── vite.config.ts
└── vitest.config.ts
```

## Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) to make the TypeScript language service aware of `.vue` types.

### Run End-to-End Tests with [Cypress](https://www.cypress.io/)
- to open in GUI
```sh
npm cypress open
```
- to run via terminal
```sh
npm cypress run
```
