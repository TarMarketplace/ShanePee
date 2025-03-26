import { Button } from '@/components/button'
import { Text } from '@/components/text'

import { type Order } from '@/generated/api'

import AddReviewContainer from './add-review-container'
import CompleteOrderContainer from './complete-order-container'

interface OrderCardFooterProps {
  order: Order
}

export default function OrderCardFooter({ order }: OrderCardFooterProps) {
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
        {/* TODO: check using new review field of order */}
        {order.status === 'PREPARING' ? (
          <AddReviewContainer />
        ) : order.status === 'DELIVERING' ? (
          <CompleteOrderContainer order={order} />
        ) : null}
        <Button variant='outline' colorVariant='outline'>
          รายละเอียด
        </Button>
      </div>
    </div>
  )
}
