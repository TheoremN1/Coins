import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  base: '/',
   server: {
    host: 'localhost',
    port: 9005
  //   proxy: {
  //     "/": {
  //       target: "http://127.0.0.1:9009",
  //       changeOrigin: true,
  //       rewrite: (path) => path.replace(/^\/api/, "/"),
  //     },
  //   },
   },
})
