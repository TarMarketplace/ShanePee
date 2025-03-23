import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

const PROTECTED_ROUTES = [
  '/product/create',
  '/product/edit',
  '/user/account',
  '/cart',
  '/order-history',
]
const AUTH_ROUTES = ['/login', '/register']

export default async function middleware(request: NextRequest) {
  const session = request.cookies.get('session')?.value

  const isProtectedRoute = PROTECTED_ROUTES.includes(request.nextUrl.pathname)
  const isAuthRoute = AUTH_ROUTES.includes(request.nextUrl.pathname)

  if (isProtectedRoute && !session) {
    return NextResponse.redirect(new URL('/login', request.nextUrl))
  }

  if (isAuthRoute && session) {
    return NextResponse.redirect(new URL('/', request.nextUrl))
  }

  return NextResponse.next()
}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|.*\\.png$).*)'],
}
