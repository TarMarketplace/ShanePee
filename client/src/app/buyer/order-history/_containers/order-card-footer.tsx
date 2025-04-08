import Link from 'next/link'
import { useEffect, useState } from 'react'

import { Button } from '@/components/button'
import { Text } from '@/components/text'

import AddReviewContainer from '@/app/_containers/add-review-container'
import { type Order, getSellerById } from '@/generated/api'

import CompleteOrderContainer from '../../../_containers/complete-order-container'

interface OrderCardFooterProps {
  order: Order
}

export default function OrderCardFooter({ order }: OrderCardFooterProps) {
  const [seller, setSeller] = useState<
    { name: string; photo: string } | undefined
  >(undefined)

  useEffect(() => {
    const fetchSeller = async () => {
      const sellerData = await getSellerById({ path: { id: order.seller_id } })
      if (sellerData.data) {
        setSeller({
          name: `${sellerData.data.first_name} ${sellerData.data.last_name}`,
          photo: sellerData.data.photo || '',
        })
      }
    }

    fetchSeller()
  }, [order.seller_id])

  return (
    <div className='flex w-full flex-col items-end gap-2.5 py-2.5 sm:flex-row sm:items-center sm:justify-between'>
      <Text variant='sm-regular' desktopVariant='md-regular'>
        รวมการสั่งซื้อ{' '}
        <span className='font-semibold text-primary'>
          ฿{' '}
          {(order.order_items ?? [])
            .reduce((acc, item) => acc + (item.art_toy?.price ?? 0), 0)
            .toFixed(2)}
        </span>
      </Text>
      <div className='flex items-center justify-center gap-3'>
        {order.status === 'PREPARING' && seller ? (
          <AddReviewContainer
            name={seller.name}
            photo={seller.photo}
            orderID={order.id}
          />
        ) : order.status === 'DELIVERING' ? (
          <CompleteOrderContainer order={order} />
        ) : null}
        <Link href={`/buyer/order-detail/${order.id}`}>
          <Button variant='outline' colorVariant='outline'>
            รายละเอียด
          </Button>
        </Link>
      </div>
    </div>
  )
}
