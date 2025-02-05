'use client'

import { LoginPageContainer } from '../_containers/login-page-container'

export default function Login() {
  return (
    <main className='flex w-full flex-col items-center gap-8 bg-background'>
      <section className='flex w-full items-center justify-center gap-10 bg-background-gradient p-12'>
        <LoginPageContainer />
      </section>
    </main>
  )
}
