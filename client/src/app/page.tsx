import Image from 'next/image'

import { Text } from '@/components/text'

import { ProductListContainer } from './_containers/product-list-container'
import { RecommendedStoreContainer } from './_containers/recommended-store-container'

export default function Home() {
  return (
    <main className='flex w-full flex-col items-center gap-8 bg-background'>
      <section className='flex w-full items-center justify-center gap-10 bg-background-gradient px-4 pb-6 pt-8'>
        <article className='grid h-max w-full max-w-screen-lg grid-cols-1 items-center justify-center gap-10 md:grid-cols-[2fr,1fr]'>
          <div className='relative hidden aspect-video size-full sm:block md:aspect-auto'>
            <Image
              src='/tar-marketplace-group.jpg'
              alt=''
              fill
              className='rounded-xl object-cover'
            />
          </div>
          <RecommendedStoreContainer />
        </article>
      </section>
      <section className='flex size-full max-w-screen-lg flex-col items-center gap-4 px-4 pb-4 pt-0 md:gap-6 md:py-8'>
        <Text
          variant='heading-md'
          desktopVariant='heading-xl'
          className='w-full text-center md:text-start'
        >
          Art Toys แนะนำสำหรับคุณ
        </Text>
        <hr className='w-full border border-primary md:border-2' />
        <ProductListContainer />
      </section>
    </main>
  )
}
