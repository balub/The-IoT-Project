import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      "/protected": {
        target: "http://localhost:9090",
        changeOrigin: true,
        secure: false,
      },
    },
  },
})
