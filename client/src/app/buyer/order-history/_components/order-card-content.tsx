import Image from 'next/image'

import { Text } from '@/components/text'

import { type Order } from '@/generated/api'

interface OrderCardContentProps {
  order: Order
}

export default function OrderCardContent({ order }: OrderCardContentProps) {
  return (
    <div className='flex w-full flex-col divide-y'>
      {order.order_items?.slice(0, 3).map((item) => (
        <div key={item.id} className='flex gap-2.5 py-2.5'>
          <div className='relative aspect-video h-16 sm:h-24'>
            <Image
              src={item.art_toy?.photo ?? ''}
              alt={item.art_toy?.name ?? ''}
              fill
              className='object-cover'
            />
          </div>
          <div className='flex w-full min-w-0 flex-col justify-between gap-2.5 sm:justify-start'>
            <Text variant='lg-regular' className='truncate'>
              {item.art_toy?.name}
            </Text>
            <div className='flex justify-between'>
              <Text variant='md-regular'>x1</Text>
              <Text variant='md-regular'>฿ {item.art_toy?.price}</Text>
            </div>
          </div>
        </div>
      ))}
      {(order.order_items?.length ?? 0) > 3 ? (
        <Text variant='md-regular' className='w-full text-center text-grey-500'>
          + รายการอื่น ๆ อีก {(order.order_items?.length ?? 3) - 3} รายการ
        </Text>
      ) : null}
    </div>
  )
}
