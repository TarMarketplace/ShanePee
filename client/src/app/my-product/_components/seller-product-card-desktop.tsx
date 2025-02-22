import Image from 'next/image'
import { useMemo } from 'react'

import { Text } from '@/components/text'

import type { ArtToy } from '@/generated/api'

export interface SellerProductCardDesktopProps {
  product: ArtToy
}

function SellerProductCardDesktop({ product }: SellerProductCardDesktopProps) {
  const formattedDate = useMemo(() => {
    const date = new Date(product.release_date)
    return `${date.getDate()}-${date.getMonth()}-${date.getFullYear() + 543}`
  }, [product])

  return (
    <div className='aspect-[250/320] size-full max-w-64 overflow-hidden rounded-xl bg-card shadow-sm'>
      <div className='relative h-2/5 w-full'>
        <Image src={product.photo as string} alt={product.name} fill />
      </div>
      <div className='flex h-3/5 w-full flex-col justify-between p-3'>
        <div className='flex flex-col gap-2'>
          <Text variant='xs-regular' desktopVariant='md-regular'>
            {product.name}
          </Text>
        </div>
        <div className='flex justify-between'>
          <div className='flex items-center'>
            <div className='flex flex-col'>
              <Text variant='sm-regular' className='text-grey-500'>
                ราคาป้าย
              </Text>
              <Text variant='md-regular'>฿ {product.price}</Text>
            </div>
          </div>
          <div className='flex justify-end'>
            <div className='flex flex-col'>
              <Text variant='sm-regular' className='text-end text-grey-500'>
                วันที่วางขาย
              </Text>
              <Text variant='md-regular'>{formattedDate}</Text>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export { SellerProductCardDesktop }
