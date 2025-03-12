import Image from 'next/image'
import { useState } from 'react'

import { Text } from '@/components/text'

import type { ArtToy, Order } from '@/generated/api'
import { formatPrice } from '@/utils/format-price'

import { Button } from '../button'
import { OrderStatus } from './order-status'

export interface OrderCardProps {
  buyer_name: string
  order: Order
  order_items: ArtToy[]
}

const OrderCard = ({ buyer_name, order, order_items }: OrderCardProps) => {
  const [expanded, setExpanded] = useState(false)

  const order_date = new Date(order.created_at)
  const order_day = String(order_date.getUTCDate()).padStart(2, '0')
  const order_month = String(order_date.getUTCMonth() + 1).padStart(2, '0')
  const order_year = order_date.getUTCFullYear() + 543

  const numberOfItems = window.innerWidth < 640 ? 1 : 3
  const itemsToShow = expanded
    ? order_items
    : order_items.slice(0, numberOfItems)
  const totalPrice = order_items.reduce(
    (sum, item) => sum + (item.price || 0),
    0
  )

  return (
    <div className='w-full divide-y rounded-lg bg-white p-3 shadow'>
      <div className='flex pb-2'>
        <div>
          <Text
            variant='sm-regular'
            className='inline-block min-w-max text-grey-500'
          >
            SHANE-{order.id} (ขายเมื่อ {order_day}/{order_month}/{order_year})
          </Text>
          <Text variant='md-semibold'>ผู้ซื้อ: {buyer_name}</Text>
        </div>
        <div className='flex w-full items-center justify-end'>
          <OrderStatus status={order.status} />
        </div>
      </div>

      <div>
        <div className='divide-y'>
          {itemsToShow.map((item, index) => (
            <div key={index} className='flex h-[104px] gap-2 py-2'>
              <Image
                src={item.photo || 'https://placehold.co/150x96.png'}
                alt={item.name}
                width={150}
                height={96}
              />
              <div className='flex w-full flex-col truncate'>
                <Text
                  variant='lg-regular'
                  className='h-full truncate md:h-auto'
                >
                  {item.name}
                </Text>
                <Text variant='lg-regular' className='justify-end'>
                  x1
                </Text>
              </div>
              <div className='flex h-full items-end md:items-center'>
                <Text variant='lg-regular' className='inline-block'>
                  {formatPrice(item.price)}
                </Text>
              </div>
            </div>
          ))}
        </div>
        {order_items.length > numberOfItems && (
          <div className='flex items-center justify-center border-t'>
            {!expanded ? (
              <Text
                variant='md-regular'
                className='cursor-pointer p-2 text-grey-500'
                onClick={() => setExpanded(true)}
              >
                + แสดงรายการอื่นๆ อีก {order_items.length - numberOfItems}{' '}
                รายการ
              </Text>
            ) : (
              <Text
                variant='md-regular'
                className='cursor-pointer p-2 text-grey-500'
                onClick={() => setExpanded(false)}
              >
                ซ่อน
              </Text>
            )}
          </div>
        )}
      </div>

      <div className='flex flex-col pb-2 pt-4 md:flex-row md:items-center'>
        <div className='flex w-full justify-end md:justify-normal'>
          <Text variant='lg-regular' className='inline-block pr-2'>
            รวมการสั่งซื้อ
          </Text>
          <Text variant='xl-semibold' className='inline-block text-primary'>
            {formatPrice(totalPrice)}
          </Text>
        </div>
        <div className='mt-4 flex justify-end gap-2 md:mt-0 md:justify-normal'>
          {order.status == 'PENDING' && (
            <Button variant='filled'>เพิ่มรหัสติดตาม</Button>
          )}
          <Button variant='outline' colorVariant='outline'>
            รายละเอียด
          </Button>
        </div>
      </div>
    </div>
  )
}

export { OrderCard }
