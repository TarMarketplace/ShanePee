import { defineConfig } from '@hey-api/openapi-ts'

export default defineConfig({
  input: '../api/docs/openapi.json',
  output: 'src/generated/api',
  plugins: [
    {
      name: '@hey-api/client-next',
      runtimeConfigPath: './src/lib/api-client.ts',
    },
  ],
})
