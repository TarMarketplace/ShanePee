'use client'

import { useState } from 'react'
import { toast } from 'sonner'

import { Button } from '@/components/button'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/popover'

import { type Order, completeOrder } from '@/generated/api'

import CompleteOrderModal from '../../components/complete-order-modal'

interface CompleteOrderContainerProps {
  order: Order
}

export default function CompleteOrderContainer({
  order,
}: CompleteOrderContainerProps) {
  const [open, setOpen] = useState(false)
  const handleClose = () => {
    setOpen(false)
  }

  const handleSubmit = async () => {
    const { response } = await completeOrder({ path: { id: order.id } })

    if (response.ok) {
      toast.success('Order is completed')
      location.reload()
    } else {
      toast.error('Failed to complete order')
    }
  }

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button variant='filled'>ยืนยันการรับสินค้า</Button>
      </PopoverTrigger>
      <PopoverContent className='size-fit'>
        <CompleteOrderModal
          handleClose={handleClose}
          handleSubmit={handleSubmit}
          order={order}
        />
      </PopoverContent>
    </Popover>
  )
}
