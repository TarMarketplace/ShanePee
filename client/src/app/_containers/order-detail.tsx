import { Icon } from '@iconify/react'
import { toast } from 'sonner'

import { OrderDetailCard } from '@/components/order-detail-card'
import { Text } from '@/components/text'

import { type Order } from '@/generated/api'

import { AddressContainer } from './address-container'
import OrderBadge from './order-badge'

interface OrderDetailProps {
  order: Order
}

export function OrderDetail({ order }: OrderDetailProps) {
  const copyToClipboard = () => {
    navigator.clipboard
      .writeText(order?.tracking_number ?? '')
      .then(() => {
        toast.success('Copied to clipboard')
      })
      .catch(() => {
        toast.error('Failed to copy')
      })
  }

  return (
    <>
      <div className='flex w-full items-center gap-2 sm:border-b-4 sm:border-primary sm:pb-3'>
        <Icon icon='tdesign:cart-filled' className='size-6 sm:size-8' />
        <Text variant='heading-md' desktopVariant='heading-lg'>
          รายละเอียดคำสั่งซื้อ
        </Text>
      </div>
      <div className='flex w-full flex-col gap-5 sm:gap-6 sm:px-6'>
        <AddressContainer />
        <div className='flex w-full flex-col gap-2'>
          <div className='flex gap-2'>
            <Text variant='md-semibold' desktopVariant='heading-sm'>
              สถานะการจัดส่ง
            </Text>
            <OrderBadge order={order} className='sm:hidden' />
          </div>
          <div className='flex flex-col gap-1'>
            <div className='flex items-center gap-2.5'>
              <OrderBadge order={order} className='hidden sm:block' />
              {order.status === 'PREPARING' ? (
                <Text variant='sm-regular' desktopVariant='md-regular'>
                  ยังไม่ปรากฎเลขติดตามพัสดุ
                </Text>
              ) : (
                <div className='flex gap-1'>
                  <Text variant='sm-regular' desktopVariant='md-regular'>
                    {order.delivery_service ?? ''}
                    {' : '}
                    {order.tracking_number ?? ''}
                  </Text>
                  <button onClick={copyToClipboard}>
                    <Icon
                      icon='fluent:copy-24-regular'
                      className='size-5 cursor-pointer'
                    />
                  </button>
                </div>
              )}
            </div>
            <Text variant='sm-regular' desktopVariant='md-regular'>
              27/01/2568 18:42 - พัสดุถึงสาขาปลายทาง: บ้านโจ้วิจักษ์
            </Text>
          </div>
        </div>
        <OrderDetailCard order={order} />
      </div>
    </>
  )
}
