import { env } from '@/env'

export async function POST(request: Request) {
  try {
    const req = await request.json()
    const res = await fetch(`${env.BASE_API_URL}/auth/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        email: req.email,
        password: req.password,
      }),
    })

    const data = await res.json()
    return Response.json(data, { status: res.status })
  } catch {
    return new Response('Internal Server Error', { status: 500 })
  }
}
