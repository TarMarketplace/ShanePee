import { env } from '@/env'

export async function POST(request: Request) {
  try {
    const req = await request.json()
    const res = await fetch(`${env.BASE_API_URL}/auth/register`, {
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
  } catch (error) {
    return Response.json({ error })
  }
}
