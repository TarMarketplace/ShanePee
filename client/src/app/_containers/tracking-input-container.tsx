import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { z } from 'zod'

import { Button } from '@/components/button'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/popover'
import { TrackingInputCard } from '@/components/tracking-input-card'
import type { TrackingInputCardSchema } from '@/components/tracking-input-card/index.stories'

import { type Order } from '@/generated/api'

const trackingInputCardSchema = z.object({
  trackingNumberValue: z.string().min(1, 'Tracking number is required'),
  deliveryCompanyValue: z.enum(['Shopee express', 'Kerry', 'Flash']),
})

interface TrackingInputCardContainerProps {
  order: Order
}

export function TrackingInputContainer({
  order,
}: TrackingInputCardContainerProps) {
  const form = useForm<TrackingInputCardSchema>({
    resolver: zodResolver(trackingInputCardSchema),
    defaultValues: {
      trackingNumberValue: '',
      deliveryCompanyValue: 'Shopee express',
    },
  })
  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button variant='filled'>เพิ่มรหัสติดตาม</Button>
      </PopoverTrigger>
      <PopoverContent className='size-fit'>
        <TrackingInputCard
          id={order.id.toString()}
          name={order.order_items?.[0].art_toy?.name ?? ''}
          form={form}
          onSubmit={(data) => {
            console.log(data)
          }}
        />
      </PopoverContent>
    </Popover>
  )
}
