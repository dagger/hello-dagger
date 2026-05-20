# hello-dagger

This is an example application for use with the Dagger Quickstart. It uses the Vue 3 + Vite template, with minor modifications.

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
npm run test:unit
```

### Run End-to-End Tests with [Cypress](https://www.cypress.io/)

```sh
npm run test:e2e:dev
```

This runs the end-to-end tests against the Vite development server.
It is much faster than the production build.

But it's still recommended to test the production build with `test:e2e` before deploying (e.g. in CI environments):

```sh
npm run build
npm run test:e2e
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```

## Go Demo Service

This branch also includes a small Go HTTP service that can be used as a backend
or second workspace target in demos.

```sh
cd services/web
go test ./...
go run .
```

The service listens on `:8080` by default and exposes `GET /` and
`GET /healthz`.
