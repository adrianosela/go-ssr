# go-ssr

Demo of a Go server-side rendered, Svelte-hydrated, web application.

### Decision-Making

- Go is explicit, fast, and type safe. Its html/template library escapes JavaScript automatically.
- Gin is simple, pretty, and has a lot of community support (many libraries and middlewares).
- Yarn is faster than npm. Bun is not ready.
- Svelte is simpler than React and has strong library and community support. It compiles at build time, and we use it to enrich Go-rendered pages — not to manage routing or state.

### Set-Up Steps

Pre-Requisites:
- go
- yarn (requires node)

(1) Initialize go module

```
go mod init
```

(2) Add main.go

(3) Add go libraries to module

```
go get -u "github.com/gin-gonic/gin"
```

(4) Create (and change into) empty /web directory

```
mkdir web && cd web
```

(4) Initialize Vite + Svelte (+ TypeScript) project

```
yarn create vite . --template svelte-ts
```

(5) Add Vite to project (not sure if needed -- I had to)

```
yarn add vite
```

(5) Update Vite config (`vite.config.ts`) to write build output to ../static/vite/dist

```
import { defineConfig } from 'vite';
import { svelte } from '@sveltejs/vite-plugin-svelte';
import path from 'path';

export default defineConfig({
  plugins: [svelte()],
  build: {
    outDir: path.resolve(__dirname, '../static/vite/dist'),
    emptyOutDir: true,
    sourcemap: true
  }
});
```
