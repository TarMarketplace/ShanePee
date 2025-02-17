import { env } from '@/env'
import type { CreateClientConfig } from '@/generated/api/client.gen'

export const createClientConfig: CreateClientConfig = (config) => ({
  ...config,
  baseUrl: env.NEXT_PUBLIC_BASE_API_URL,
  credentials: 'include',
})
