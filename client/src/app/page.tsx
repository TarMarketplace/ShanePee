import { Icon } from '@iconify/react'
import Image from 'next/image'

import { Text } from '@/components/text'

import { RecommendedStore } from './_components/recommended-store'
import { ProductListContainer } from './_containers/product-list-container'

export default function Home() {
  return (
    <main className='flex w-full flex-col items-center gap-8 bg-background'>
      <section className='flex w-full items-center justify-center gap-10 bg-background-gradient px-4 pb-6 pt-8'>
        <article className='grid h-max w-full max-w-screen-lg grid-cols-1 items-center justify-center gap-10 md:grid-cols-[2fr,1fr]'>
          <div className='relative hidden aspect-video size-full sm:block md:aspect-auto'>
            <Image
              src='https://placehold.co/740x420.png'
              alt=''
              fill
              className='object-cover'
            />
          </div>
          <div className='flex w-full flex-col gap-3.5 rounded-xl bg-white px-5 pb-4 pt-5 shadow-md'>
            <div className='flex items-center gap-3'>
              <Icon icon='ri:store-3-fill' className='size-9 text-primary' />
              <Text variant='heading-sm' desktopVariant='heading-md'>
                ร้านค้าแนะนำ
              </Text>
            </div>
            <div className='grid grid-cols-2 gap-3'>
              {Array.from({ length: 4 }).map((_, index) => (
                <RecommendedStore
                  key={index}
                  name={`ร้านค้า ${index}`}
                  average_rating={index + 2}
                  review_count={index * 10 + 1}
                  image_src='https://placehold.co/80x80.png'
                />
              ))}
            </div>
          </div>
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
