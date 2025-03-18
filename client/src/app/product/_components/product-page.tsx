import { Icon } from '@iconify/react'
import Image from 'next/image'

import { Avatar, AvatarFallback, AvatarImage } from '@/components/avatar'
import { Button } from '@/components/button'
import { Text } from '@/components/text'

import type { ArtToy } from '@/generated/api'

type ProductPageProps = {
  product: ArtToy
  handleCartButton: () => void
  isInCart: boolean
}

export function ProductPage({
  product,
  handleCartButton,
  isInCart,
}: ProductPageProps) {
  return (
    <div className='mx-auto flex w-full max-w-screen-lg flex-col py-8'>
      <div className='grid grid-cols-[2fr_3fr] gap-4'>
        <div className='flex flex-col gap-2.5'>
          <div className='relative aspect-video w-full overflow-hidden rounded-lg'>
            <Image
              src={product.photo ?? ''}
              alt={product.name + "'s photo"}
              fill
              className='sticky object-cover'
            />
          </div>
          <div className='flex items-center justify-between rounded-lg bg-grey-50 p-2 shadow-sm'>
            <div className='flex items-center gap-2'>
              <Avatar className='size-20'>
                <AvatarImage
                  src='https://placehold.co/80x80.png'
                  width={80}
                  height={80}
                />
                <AvatarFallback>{product.owner_id}</AvatarFallback>
              </Avatar>
              <div className='flex flex-col'>
                <span className='text-sm font-semibold'>
                  {product.owner_id}
                </span>
                <div className='flex items-center gap-2'>
                  <div className='flex items-center'>
                    <Icon
                      icon='material-symbols:star-rounded'
                      className='size-5 text-warning md:size-6'
                    />
                    <Text variant='xs-regular' desktopVariant='sm-regular'>
                      4.6
                    </Text>
                  </div>
                  <div className='flex items-center text-grey-500'>
                    <Icon icon='typcn:location' className='size-4 md:size-5' />
                    <Text variant='xs-regular' desktopVariant='sm-regular'>
                      Bangkok
                    </Text>
                  </div>
                </div>
              </div>
            </div>
            <Button variant='filled' colorVariant='outline'>
              <Icon icon='ri:store-3-fill' className='size-4' />
              ดูร้านค้า
            </Button>
          </div>
        </div>
        <div className='flex flex-col gap-6'>
          <div className='flex flex-col gap-3'>
            <Text variant='heading-lg'>{product.name}</Text>
            <Text variant='heading-md' className='text-primary'>
              ฿ {product.price}
            </Text>
            <div className='flex items-center gap-4'>
              <Button
                variant='filled'
                colorVariant={isInCart ? 'outline' : 'default'}
                disabled={!product.availability}
                onClick={() => handleCartButton()}
              >
                <Icon
                  icon={
                    isInCart ? 'mingcute:check-fill' : 'tdesign:cart-filled'
                  }
                  className='size-4'
                />
                {isInCart ? 'เพิ่มในตะกร้าแล้ว' : 'เพิ่มในตะกร้า'}
              </Button>
              <Button variant='filled' colorVariant='outline'>
                <Icon icon='material-symbols:chat' className='size-4' />
                สอบถามข้อมูล
              </Button>
            </div>
          </div>
          <hr className='w-full border border-primary' />
          <Text>{product.description}</Text>
        </div>
      </div>
    </div>
  )
}
