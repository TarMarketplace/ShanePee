import { Icon } from '@iconify/react'
import Image from 'next/image'

import { Text } from '@/components/text'

import type { ArtToy } from '@/generated/api'

import { Badge } from '../badge'

export interface ProductCardProps {
  // TODO: remove any
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  product: ArtToy | any
}

const ProductCard = ({ product }: ProductCardProps) => {
  return (
    <div className='aspect-[170/240] size-full max-w-64 overflow-hidden rounded-xl bg-card shadow-cardbox md:aspect-[250/320]'>
      <div className='relative h-2/5 w-full'>
        <Image src={product.photo as string} alt={product.name} fill />
      </div>
      <div className='flex h-3/5 w-full flex-col justify-between p-3'>
        <div className='flex flex-col gap-2'>
          <Text variant='sm-semibold' desktopVariant='md-semibold'>
            {product.name}
          </Text>
          {product.discount ? (
            <div className='flex flex-col'>
              <div className='flex items-center gap-2'>
                <Text
                  variant='sm-regular'
                  desktopVariant='md-regular'
                  className='italic text-primary'
                >
                  ฿{product.price - product.discount}
                </Text>
                <Badge variant='error'>
                  {((product.discount * 100) / product.price).toFixed(0)}%
                </Badge>
              </div>
              <Text
                variant='xs-regular'
                desktopVariant='sm-regular'
                className='line-through'
              >
                ฿{product.price}
              </Text>
            </div>
          ) : (
            <Text variant='sm-regular' desktopVariant='md-regular'>
              ฿{product.price}
            </Text>
          )}
        </div>
        <div className='flex items-center gap-4'>
          <div className='flex items-center'>
            <Icon
              icon='material-symbols:star-rounded'
              className='size-5 text-warning md:size-6'
            />
            <Text variant='xs-regular' desktopVariant='sm-regular'>
              1
            </Text>
          </div>
          <div className='flex items-center text-grey-500'>
            <Icon icon='typcn:location' className='size-4 md:size-5' />
            <Text variant='xs-regular' desktopVariant='sm-regular'>
              {product.location}
            </Text>
          </div>
        </div>
      </div>
    </div>
  )
}

export { ProductCard }
