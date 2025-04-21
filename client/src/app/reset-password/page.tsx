import { Suspense } from 'react'

import { AuthBanner } from '../login/_components/banner'
import { ForgotPasswordContainer } from './_containers/forgot-password-container'

export default function Login() {
  return (
    <main className='grid size-full grid-cols-1 place-items-center bg-background-gradient p-4 md:p-12'>
      <div className='flex w-full justify-center gap-12 rounded-xl bg-white shadow-lg md:max-w-5xl md:p-12 lg:gap-24 lg:px-24'>
        <AuthBanner />
        <Suspense fallback={<div>Loading...</div>}>
          <ForgotPasswordContainer />
        </Suspense>
      </div>
    </main>
  )
}
