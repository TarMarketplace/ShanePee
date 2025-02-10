import { env } from '@/env'

export async function POST() {
  try {
    await fetch(`${env.BASE_API_URL}/auth/logout`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
    })
  } catch {
    return new Response('Internal Server Error', { status: 500 })
  }
}
