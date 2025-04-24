import { Icon } from '@iconify/react'
import { AnimatePresence, motion } from 'motion/react'
import Image from 'next/image'

import { Skeleton } from '@/components/skeleton'
import { Text } from '@/components/text'

import type { CartItem } from '@/generated/api'

interface CartBoxProps {
  items: CartItem[]
  onDeleteItem: (id: number) => void
}

export function CartBox({ items, onDeleteItem }: CartBoxProps) {
  if (items.length === 0) {
    return (
      <div className='flex aspect-[3/4] w-full flex-col items-center justify-center rounded-lg p-3 text-center shadow-cardbox md:aspect-video'>
        <Text
          desktopVariant='heading-lg'
          variant='heading-md'
          className='text-grey-500'
        >
          คุณยังไม่มีรายการสินค้า
        </Text>
        <Text
          desktopVariant='heading-md'
          variant='md-semibold'
          className='text-grey-500'
        >
          กรุณาเพิ่มสินค้าลงตะกร้าก่อนสั่งซื้อ
        </Text>
      </div>
    )
  }

  return (
    <div className='flex w-full flex-col divide-y divide-grey-200 rounded-lg p-3 shadow-cardbox'>
      <Text
        desktopVariant='md-semibold'
        variant='sm-semibold'
        className='pb-2.5'
      >
        ร้านค้า:{' '}
        {items[0].art_toy?.owner &&
        (items[0].art_toy.owner.first_name || items[0].art_toy.owner.last_name)
          ? items[0].art_toy.owner.first_name +
            ' ' +
            items[0].art_toy.owner.last_name
          : 'ไม่พบชื่อร้านค้า'}
      </Text>
      <AnimatePresence mode='sync' initial={false}>
        {items.map((item) => (
          <motion.div
            className='flex size-full py-2.5'
            animate={{ x: 0, opacity: 1 }}
            exit={{ x: 100, opacity: 0 }}
            transition={{ duration: 0.15, ease: 'easeInOut' }}
            layout
            key={item.id}
          >
            <div className='relative aspect-video h-16 md:h-24'>
              <Image
                src={item.art_toy?.photo ?? ''}
                alt={item.art_toy?.name ?? ''}
                fill
                className='object-cover'
              />
            </div>
            <div className='grid w-full grid-rows-2 pl-2.5 md:flex md:flex-row'>
              <Text
                desktopVariant='lg-regular'
                variant='sm-regular'
                className='line-clamp-1 w-full'
              >
                {item.art_toy?.name ?? <Skeleton />}
              </Text>
              <div className='grid grid-cols-2 md:grid-cols-none md:grid-rows-2'>
                <Text
                  className='mt-auto text-nowrap md:mt-0'
                  variant='sm-regular'
                  desktopVariant='md-regular'
                >
                  ฿ {item.art_toy?.price.toLocaleString()}
                </Text>
                <Icon
                  icon='tabler:trash'
                  onClick={() => {
                    onDeleteItem(item.id)
                  }}
                  className='ml-auto mt-auto size-5 cursor-pointer text-grey-500 hover:text-error-500 md:size-6'
                />
              </div>
            </div>
          </motion.div>
        ))}
      </AnimatePresence>
      <Text
        className='w-full py-2.5 text-right'
        variant='sm-regular'
        desktopVariant='md-regular'
      >
        รวมการสั่งซื้อ{' '}
        <Text
          as='span'
          className='text-primary-500'
          desktopVariant='xl-semibold'
          variant='sm-semibold'
        >
          ฿{' '}
          {items
            .reduce((acc, item) => acc + (item.art_toy?.price ?? 0), 0)
            .toLocaleString()}
        </Text>
      </Text>
    </div>
  )
}
