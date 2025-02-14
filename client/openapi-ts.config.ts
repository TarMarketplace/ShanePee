import { defineConfig } from '@hey-api/openapi-ts'

export default defineConfig({
  input: '../api/docs/openapi.json',
  output: 'src/generated/api',
  plugins: ['@hey-api/client-next'],
})
