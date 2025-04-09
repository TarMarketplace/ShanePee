'use client'

import { Icon } from '@iconify/react'
import { useEffect, useState } from 'react'

import { Button } from '@/components/button'
import { Text } from '@/components/text'

import AddReviewContainer from '@/app/_containers/add-review-container'
import CompleteOrderContainer from '@/app/_containers/complete-order-container'
import { OrderDetail } from '@/app/_containers/order-detail'
import { type Order, getOrderOfBuyer, getSellerById } from '@/generated/api'

export default function BuyerOrderDetailPage({
  params,
}: {
  params: { id: string }
}) {
  const [order, setOrder] = useState<Order | undefined>(undefined)
  const [seller, setSeller] = useState<
    { name: string; photo: string } | undefined
  >(undefined)

  useEffect(() => {
    const fetchOrder = async () => {
      const { data, error } = await getOrderOfBuyer({
        path: { orderID: parseInt(params.id) },
      })
      if (error) {
        console.error(error)
      }

      setOrder(data)

      if (data) {
        const sellerData = await getSellerById({ path: { id: data.seller_id } })
        if (sellerData.data) {
          setSeller({
            name: `${sellerData.data.first_name} ${sellerData.data.last_name}`,
            photo: sellerData.data.photo || '',
          })
        }
      }
    }

    fetchOrder()
  }, [params.id])

  return (
    <main className='grid size-full grid-cols-1 place-items-center p-4 md:p-12'>
      <div className='flex w-full flex-col gap-5 sm:w-fit sm:min-w-[60%] sm:gap-6'>
        {order ? (
          <>
            <OrderDetail order={order} />
            <div className='ml-auto flex gap-2.5 sm:px-6'>
              {order.status === 'COMPLETED' && seller ? (
                <AddReviewContainer
                  name={seller.name}
                  photo={seller.photo}
                  orderID={order.id}
                />
              ) : order.status === 'DELIVERING' ? (
                <CompleteOrderContainer order={order} />
              ) : null}
              <Button variant='outline' colorVariant='outline'>
                <Icon icon='material-symbols:chat' />
                <Text variant='sm-semibold' desktopVariant='md-semibold'>
                  ติดต่อผู้ขาย
                </Text>
              </Button>
            </div>
          </>
        ) : null}
      </div>
    </main>
  )
}
