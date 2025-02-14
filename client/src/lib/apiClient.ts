import { env } from '@/env'
import { client } from '@/generated/api/client.gen'

client.setConfig({
  baseUrl: env.NEXT_PUBLIC_BASE_API_URL,
  credentials: 'include',
})
