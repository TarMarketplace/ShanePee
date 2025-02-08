import { AccountContainer } from './_containers/account-container'

export default function Account() {
  return (
    <main className='grid size-full grid-cols-1 place-items-center p-4 md:p-12'>
      <div className='flex w-full justify-center gap-12 rounded-xl bg-white md:max-w-6xl md:p-12 lg:gap-12 lg:px-12'>
        <AccountContainer />
      </div>
    </main>
  )
}
