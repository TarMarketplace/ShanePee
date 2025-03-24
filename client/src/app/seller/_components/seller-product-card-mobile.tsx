import Image from 'next/image'
import { useMemo } from 'react'

import { Text } from '@/components/text'

import type { ArtToy } from '@/generated/api'

export interface SellerProductCardDesktopProps {
  product: ArtToy
}

function SellerProductCardMobile({ product }: SellerProductCardDesktopProps) {
  const formattedDate = useMemo(() => {
    const date = new Date(product.release_date)
    return `${date.getDate()}-${date.getMonth()}-${date.getFullYear() + 543}`
  }, [product])

  return (
    <div className='flex h-[76px] w-full items-center gap-3 overflow-hidden rounded-lg bg-card px-2 py-1 shadow-sm'>
      <div className='relative h-full w-1/4'>
        <Image
          className='rounded-lg object-cover'
          src={product.photo as string}
          alt={product.name}
          fill
        />
      </div>
      <div className='flex grow gap-2'>
        <div className='flex grow gap-2 self-center'>
          <Text variant='xs-regular'>{product.name}</Text>
        </div>
        <div className='flex flex-col'>
          <Text
            variant='xs-regular'
            className='text-nowrap text-end text-grey-500'
          >
            ราคา
          </Text>
          <Text variant='xs-regular' className='text-nowrap'>
            ฿ {product.price}
          </Text>
        </div>
        <div className='flex flex-col'>
          <Text
            variant='xs-regular'
            className='text-nowrap text-end text-grey-500'
          >
            วันที่วางขาย
          </Text>
          <Text variant='xs-regular'>{formattedDate}</Text>
        </div>
      </div>
    </div>
  )
}

export { SellerProductCardMobile }
