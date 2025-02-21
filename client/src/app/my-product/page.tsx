import { Suspense } from 'react'

import { ListProductContainer } from './_containers/list-product-container'

export default function Login() {
  return (
    <main className='grid size-full grid-cols-1 place-items-center p-4 md:p-12'>
      <Suspense fallback={<div>Loading...</div>}>
        <ListProductContainer />
      </Suspense>
    </main>
  )
}
