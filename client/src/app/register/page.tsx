'use client'

import { RegisterPageContainer } from '../_containers/register-page-container'

export default function Register() {
  return (
    <main className='flex w-full flex-col items-center gap-8 bg-background'>
      <section className='flex w-full items-center justify-center gap-10 bg-background-gradient p-4 md:p-12'>
        <RegisterPageContainer />
      </section>
    </main>
  )
}
