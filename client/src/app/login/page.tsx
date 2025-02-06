import { AuthBanner } from './_components/banner'
import { AuthContainer } from './_containers/auth-container'

export default function Login() {
  return (
    <main className='grid size-full grid-cols-1 place-items-center bg-background-gradient p-4 md:p-12'>
      <div className='flex w-full justify-center gap-12 rounded-xl bg-white shadow-lg md:max-w-5xl md:p-12 lg:gap-24 lg:px-24'>
        <AuthBanner />
        <AuthContainer />
      </div>
    </main>
  )
}
