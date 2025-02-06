import { Icon } from '@iconify/react'
import Image from 'next/image'

import { Text } from '@/components/text'

import type { Product } from '@/types/products'

import { Badge } from '../badge'

export interface ProductCardProps {
  product: Product
}

const ProductCard = ({ product }: ProductCardProps) => {
  return (
    <div className='aspect-[250/320] size-full max-w-64 overflow-hidden rounded-xl bg-card'>
      <div className='relative h-2/5 w-full'>
        <Image src={product.image} alt={product.name} fill />
      </div>
      <div className='flex h-3/5 w-full flex-col justify-between p-3'>
        <div className='flex flex-col gap-2'>
          <Text variant='md-semibold'>{product.name}</Text>
          {product.discount ? (
            <div className='flex flex-col'>
              <div className='flex items-center gap-2'>
                <Text variant='md-regular' className='italic text-primary'>
                  ฿{product.price - product.discount}
                </Text>
                <Badge variant='error'>
                  {((product.discount * 100) / product.price).toFixed(0)}%
                </Badge>
              </div>
              <Text variant='sm-regular' className='line-through'>
                ฿{product.price}
              </Text>
            </div>
          ) : (
            <Text variant='md-regular'>฿{product.price}</Text>
          )}
        </div>
        <div className='flex items-center gap-4'>
          <div className='flex items-center'>
            <Icon
              icon='material-symbols:star-rounded'
              className='size-6 text-warning'
            />
            <Text variant='sm-regular'>{product.rating.toFixed(1)}</Text>
          </div>
          <div className='flex items-center text-grey-500'>
            <Icon icon='typcn:location' className='size-5' />
            <Text variant='sm-regular'>{product.location}</Text>
          </div>
        </div>
      </div>
    </div>
  )
}

export { ProductCard }
