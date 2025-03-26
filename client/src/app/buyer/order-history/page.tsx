'use client'

import { useEffect, useState } from 'react'

import { type Order, getOrdersOfBuyer } from '@/generated/api'

import Heading from './_components/heading'
import OrderCardContent from './_components/order-card-content'
import OrderCardHeader from './_components/order-card-header'
import OrderCardFooter from './_containers/order-card-footer'

export default function BuyerOrderHistoryPage() {
  const [orders, setOrders] = useState<Order[]>([])
  const [filter, setFilter] = useState<Order['status'] | 'ALL'>('ALL')

  useEffect(() => {
    const fetchOrders = async () => {
      const { data, error } = await getOrdersOfBuyer()
      if (error) {
        console.error(error)
      }

      console.log(data)
      setOrders(data?.data ?? [])
    }

    fetchOrders()
  }, [])

  return (
    <main className='grid size-full grid-cols-1 place-items-center p-4 md:p-12'>
      <div className='flex w-full flex-col sm:w-fit sm:min-w-[60%] sm:divide-y-4 sm:divide-primary'>
        <Heading filter={filter} setFilter={setFilter} />
        <section className='flex w-full flex-col gap-6 py-6 md:px-5'>
          {orders
            .filter((order) => order.order_items)
            .filter((order) => order.status === filter || filter === 'ALL')
            .map((order) => (
              <div
                className='flex w-full flex-col gap-3 divide-y rounded-lg p-2.5 shadow-cardbox'
                key={order.id}
              >
                <OrderCardHeader order={order} />
                <OrderCardContent order={order} />
                <OrderCardFooter order={order} />
              </div>
            ))}
        </section>
      </div>
    </main>
  )
}
