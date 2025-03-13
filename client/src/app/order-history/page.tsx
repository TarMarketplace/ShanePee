'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { Icon } from '@iconify/react'
import Image from 'next/image'
import { useEffect, useState } from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'

import { Badge } from '@/components/badge'
import { Button } from '@/components/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/dropdown-menu'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/popover'
import { Text } from '@/components/text'
import { TrackingInputCard } from '@/components/tracking-input-card'
import type { TrackingInputCardSchema } from '@/components/tracking-input-card/index.stories'

import { type Order, getOrdersOfSeller } from '@/generated/api'

const trackingInputCardSchema = z.object({
  trackingNumberValue: z.string().min(1, 'Tracking number is required'),
  deliveryCompanyValue: z.enum(['Shopee express', 'Kerry', 'Flash']),
})

const STATUS_MAP_TO_LABEL: Record<Order['status'] | 'ALL', string> = {
  ALL: 'ทั้งหมด',
  PREPARING: 'รอจัดส่ง',
  DELIVERING: 'กำลังจัดส่ง',
  COMPLETED: 'จัดส่งสำเร็จ',
}

export default function OrderHistoryPage() {
  const [orders, setOrders] = useState<Order[]>([])
  const [filter, setFilter] = useState<Order['status'] | 'ALL'>('ALL')

  useEffect(() => {
    const fetchOrders = async () => {
      const { data, error } = await getOrdersOfSeller()
      if (error) {
        console.error(error)
      }

      console.log(data)
      setOrders(data?.data ?? [])
    }

    fetchOrders()
  }, [])

  const form = useForm<TrackingInputCardSchema>({
    resolver: zodResolver(trackingInputCardSchema),
    defaultValues: {
      trackingNumberValue: '',
      deliveryCompanyValue: 'Shopee express',
    },
  })

  return (
    <main className='grid size-full grid-cols-1 place-items-center p-4 md:p-12'>
      <div className='flex w-full flex-col sm:w-fit sm:min-w-[60%] sm:divide-y-4 sm:divide-primary'>
        <div className='flex justify-between py-3'>
          <div className='flex items-center'>
            <Icon icon='material-symbols:list' className='size-9' />
            <Text variant='heading-md' desktopVariant='heading-lg'>
              สินค้าของคุณ
            </Text>
          </div>

          <div className='flex items-center gap-3'>
            <Text variant='sm-regular' desktopVariant='md-regular'>
              เรียงลำดับตาม
            </Text>
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <Button>
                  <Text variant='sm-semibold' desktopVariant='md-semibold'>
                    {STATUS_MAP_TO_LABEL[filter]}
                  </Text>
                  <Icon icon='fa-solid:filter' className='size-3 sm:size-4' />
                </Button>
              </DropdownMenuTrigger>
              <DropdownMenuContent className='divide-y divide-grey-200 shadow-md'>
                {Object.entries(STATUS_MAP_TO_LABEL).map(([value, label]) => (
                  <DropdownMenuItem
                    key={value}
                    onClick={() => setFilter(value as Order['status'] | 'ALL')}
                  >
                    {label}
                  </DropdownMenuItem>
                ))}
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        </div>
        <section className='flex w-full flex-col'>
          {orders
            .filter((order) => order.order_items)
            .filter((order) => order.status === filter || filter === 'ALL')
            .map((order) => (
              <div
                className='flex w-full flex-col gap-3 divide-y p-2.5'
                key={order.id}
              >
                <div className='flex items-center gap-3 pb-2.5'>
                  <div className='flex flex-col'>
                    <Text variant='xs-regular' className='text-grey-500'>
                      {order.id} (ขายเมื่อ {order.created_at})
                    </Text>
                    <Text variant='md-semibold'>ผู้ซื้อ</Text>
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
                <div className='flex w-full flex-col divide-y'>
                  {order.order_items?.slice(0, 3).map((item) => (
                    <div key={item.id} className='flex gap-2.5 py-2.5'>
                      <div className='relative aspect-video h-24'>
                        <Image
                          src={item.art_toy?.photo ?? ''}
                          alt={item.art_toy?.name ?? ''}
                          fill
                          className='object-cover'
                        />
                      </div>
                      <div className='flex w-full flex-col gap-2.5'>
                        <Text variant='lg-regular'>{item.art_toy?.name}</Text>
                        <Text variant='md-regular'>x1</Text>
                      </div>
                      <Text
                        variant='md-regular'
                        className='my-auto h-fit text-nowrap'
                      >
                        ฿ {item.art_toy?.price}
                      </Text>
                    </div>
                  ))}
                  {(order.order_items?.length ?? 0) > 3 ? (
                    <Text variant='md-regular' className='text-grey-500'>
                      + รายการอื่น ๆ อีก {(order.order_items?.length ?? 3) - 3}{' '}
                      รายการ
                    </Text>
                  ) : null}
                </div>
                <div className='flex items-center justify-between py-2.5'>
                  <Text variant='md-regular'>
                    รวมการสั่งซื้อ{' '}
                    <span className='text-primary'>
                      ฿{' '}
                      {orders.reduce(
                        (acc, order) =>
                          acc +
                          (order.order_items ?? []).reduce(
                            (acc, item) => acc + (item.art_toy?.price ?? 0),
                            0
                          ),
                        0
                      )}
                    </span>
                  </Text>
                  <div className='flex items-center justify-center gap-3'>
                    {order.status === 'PREPARING' ? (
                      <>
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
                      </>
                    ) : null}
                    <Button variant='filled' colorVariant='outline'>
                      รายละเอียด
                    </Button>
                  </div>
                </div>
              </div>
            ))}
        </section>
      </div>
    </main>
  )
}
