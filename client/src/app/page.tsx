import { Icon } from '@iconify/react'
import Image from 'next/image'

import { ButtonCapsule } from '@/components/button-capsule'
import { ProductCard } from '@/components/product-card'
import { Text } from '@/components/text'

export default function Home() {
  return (
    <main className='flex w-full flex-col items-center gap-8 bg-background'>
      <section className='flex w-full items-center justify-center gap-10 bg-background-gradient pb-6 pt-8'>
        <Image
          src='https://placehold.co/740x420.png'
          alt=''
          width={740}
          height={420}
        />
        <div className='flex flex-col gap-3.5 rounded-xl bg-white px-5 py-[1.375rem]'>
          <div className='flex items-center gap-3'>
            <Icon icon='ri:store-3-fill' className='size-9 text-primary' />
            <Text variant='heading-md'>ร้านค้าแนะนำ</Text>
          </div>
        </div>
      </section>
      <section className='flex size-full max-w-screen-lg flex-col items-center gap-6 py-8'>
        <Text variant='heading-xl' className='w-full items-start text-start'>
          Art Toys แนะนำสำหรับคุณ
        </Text>
        <hr className='w-full border-[5px] border-primary' />
        <div className='grid size-full grid-cols-4 place-items-center gap-4'>
          {Array.from({ length: 12 }).map((_, index) => (
            <ProductCard
              key={index}
              product={{
                id: `${index}`,
                name: `Product ${index}`,
                price: 330 + index * 10,
                discount: index % 2 === 0 ? 31 + index * 10 : 0,
                rating: index % 5,
                location: 'Location 1',
                image: 'https://placehold.co/250x140.png',
              }}
            />
          ))}
        </div>
        <ButtonCapsule>View More</ButtonCapsule>
      </section>
    </main>
  )
}
