'use client'

import { LoginContainer } from '../_containers/login-container'

export default function Login() {
  return (
    <main className='grid size-full grid-cols-1 place-items-center bg-background-gradient p-4 md:p-12'>
      <LoginContainer />
    </main>
  )
}
