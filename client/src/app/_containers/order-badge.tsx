import { Badge } from '@/components/badge'

import { type Order } from '@/generated/api'

interface OrderBadgeProps {
  order: Order
  className?: string
}

export default function OrderBadge({ order, className }: OrderBadgeProps) {
  return (
    <Badge
      variant={
        order.status === 'COMPLETED'
          ? 'success'
          : order.status === 'DELIVERING'
            ? 'info'
            : 'warning'
      }
      className={className}
    >
      {order.status === 'COMPLETED'
        ? 'จัดส่งสำเร็จ'
        : order.status === 'DELIVERING'
          ? 'กำลังจัดส่ง'
          : 'รอจัดส่ง'}
    </Badge>
  )
}
