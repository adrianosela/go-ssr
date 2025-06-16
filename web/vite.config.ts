import { defineConfig } from 'vite';
import { svelte } from '@sveltejs/vite-plugin-svelte';
import { resolve } from 'path';

// https://vite.dev/config/
export default defineConfig({
  plugins: [svelte()],
  build: {
    // We need to produce a manifest to communicate built output paths
    // for individual modules to Go.
    manifest: true,

    // By default, Vite treats index.html as the app entrypoint (which we 
    // deleted), and only bundles what it thinks is needed based on that.
    // We are using Svelte purely as a widget system inside a Go SSR app,
    // so we need make Vite aware of all entrypoints explicitly.
    rollupOptions: {
      input: {
        demo: resolve(__dirname, 'src/entrypoints/demo.ts')
      }
    }
  }
});
