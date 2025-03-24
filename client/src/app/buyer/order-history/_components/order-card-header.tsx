import { Badge } from '@/components/badge'
import { Text } from '@/components/text'

import { type Order } from '@/generated/api'

interface OrderCardHeaderProps {
  order: Order
}

export default function OrderCardHeader({ order }: OrderCardHeaderProps) {
  return (
    <div className='flex items-center justify-between gap-3 pb-2'>
      <div className='flex flex-col'>
        <Text variant='xs-regular' className='text-grey-500'>
          {order.id} (ซื้อเมื่อ{' '}
          {new Date(order.created_at).toLocaleDateString('th', {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
          })}
          )
        </Text>
        <Text variant='md-semibold'>ร้านค้า: {order.seller_id}</Text>
      </div>
      <Badge
        variant={
          order.status === 'COMPLETED'
            ? 'success'
            : order.status === 'DELIVERING'
              ? 'info'
              : 'warning'
        }
      >
        {order.status === 'COMPLETED'
          ? 'จัดส่งสำเร็จ'
          : order.status === 'DELIVERING'
            ? 'กำลังจัดส่ง'
            : 'รอจัดส่ง'}
      </Badge>
    </div>
  )
}
