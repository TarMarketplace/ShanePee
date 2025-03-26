import Image from 'next/image'

import type { Order } from '@/generated/api'

import { Text } from '../text'

interface OrderDetailCardProps {
  order: Order
}

export function OrderDetailCard({ order }: OrderDetailCardProps) {
  return (
    <div className='flex flex-col gap-2.5 divide-y rounded-lg p-4 shadow-cardbox'>
      <Text variant='sm-regular' desktopVariant='md-regular'>
        <span className='block font-bold sm:inline'>
          ร้านค้า: {order.seller_id}
        </span>
        <span className='hidden sm:inline'>{' - '}</span>
        {order.id} (ซื้อเมื่อ{' '}
        {new Date(order.created_at).toLocaleDateString('th', {
          year: 'numeric',
          month: 'long',
          day: 'numeric',
        })}
        )
      </Text>
      <div className='flex w-full flex-col divide-y'>
        {order.order_items?.map((item) => (
          <div key={item.id} className='flex gap-2.5 py-2.5'>
            <div className='relative aspect-video h-16 sm:h-24'>
              <Image
                src={item.art_toy?.photo ?? ''}
                alt={item.art_toy?.name ?? ''}
                fill
                className='object-cover'
              />
            </div>
            <div className='flex w-full min-w-0 flex-col justify-between gap-2.5 sm:flex-row'>
              <div className='flex flex-col gap-2.5'>
                <Text
                  variant='sm-regular'
                  desktopVariant='lg-regular'
                  className='truncate'
                >
                  {item.art_toy?.name ?? ''}
                </Text>
                <Text desktopVariant='lg-regular' className='hidden sm:block'>
                  x1
                </Text>
              </div>
              <Text variant='sm-regular' desktopVariant='lg-regular'>
                ฿ {item.art_toy?.price ?? 0}
              </Text>
            </div>
          </div>
        ))}
      </div>
      <Text
        variant='sm-regular'
        desktopVariant='md-regular'
        className='w-full py-2.5 text-right'
      >
        รวมการสั่งซื้อ{' '}
        <span className='font-semibold text-primary'>
          ฿{' '}
          {(order.order_items ?? [])
            .reduce((acc, item) => acc + (item.art_toy?.price ?? 0), 0)
            .toFixed(2)}
        </span>
      </Text>
    </div>
  )
}
