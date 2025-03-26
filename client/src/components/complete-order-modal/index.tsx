import { Icon } from '@iconify/react'

import { Button } from '@/components/button'
import { Text } from '@/components/text'

import { type Order } from '@/generated/api'

interface CompleteOrderModalProps {
  handleSubmit: () => void
  handleClose: () => void
  order: Order
}

export default function CompleteOrderModal({
  handleClose,
  handleSubmit,
  order,
}: CompleteOrderModalProps) {
  return (
    <div className='flex w-[353px] flex-col gap-5 rounded-lg p-4 sm:w-[500px]'>
      <div className='flex w-full items-center justify-between'>
        <Text variant='heading-md'>ยืนยันการรับสินค้า</Text>
        <button onClick={handleClose}>
          <Icon icon='maki:cross' className='text-grey-500' />
        </button>
      </div>
      <div className='w-full'>
        <Text
          variant='sm-regular'
          desktopVariant='md-regular'
          className='text-grey-500'
        >
          รหัสรายการ: {order.id}
        </Text>
        <ul>
          {order.order_items?.map((item, index) => (
            <li key={item.id}>
              <Text variant='sm-regular' desktopVariant='md-regular'>
                {index + 1}. {item.art_toy?.name ?? ''}
              </Text>
            </li>
          ))}
        </ul>
      </div>
      <Text
        variant='sm-regular'
        desktopVariant='md-regular'
        className='text-error-500'
      >
        *กรุณาตรวจสอบว่าท่านได้รับสินค้าครบตามรายการข้างต้น เมื่อท่านกด
        ยืนยันรับสินค้าแล้ว ระบบจะทำการโอนเงินให้แก่ผู้ขาย โดยไม่สามารถ
        ยกเลิกได้
      </Text>
      <div className='ml-auto flex gap-2.5 pt-2'>
        <Button variant='outline' colorVariant='outline' onClick={handleClose}>
          ยกเลิก
        </Button>
        <Button variant='filled' onClick={handleSubmit}>
          ยืนยันการรับสินค้า
        </Button>
      </div>
    </div>
  )
}
