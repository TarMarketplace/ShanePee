'use client'

import { Icon } from '@iconify/react'
import { useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'
import { toast } from 'sonner'

import { Button } from '@/components/button'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/dialog'
import { Skeleton } from '@/components/skeleton'
import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import type { CartItem } from '@/generated/api'
import { checkout, getCart, removeItemFromCart } from '@/generated/api'

import { CartBox } from '../_components/cart-box'

export function CartContainer() {
  const [cart, setCart] = useState<CartItem[]>([])
  const { user } = useUser()
  const [loading, setLoading] = useState(true)
  const [checkingOut, setCheckingOut] = useState(false)
  const [checkoutSuccess, setCheckoutSuccess] = useState(false)
  const [dialogOpen, setDialogOpen] = useState(false)
  const router = useRouter()

  useEffect(() => {
    const fetchCart = async () => {
      const { data, error } = await getCart()

      if (error) {
        toast.error(error.title ?? 'Not found')
      }

      setCart(data?.data ?? [])
      setLoading(false)
    }

    fetchCart()
  }, [])

  const handleDeleteItem = (id: number) => {
    toast.promise(
      async () => {
        const { error } = await removeItemFromCart({
          path: { id },
        })

        if (error) {
          throw new Error(error.title ?? 'Error deleting item')
        }
      },
      {
        loading: 'Deleting item...',
        success: () => {
          setCart((prev) => prev.filter((item) => item.id !== id))
          return 'Item deleted!'
        },
        error: 'Error deleting item',
      }
    )
  }

  const handleCheckout = async () => {
    setCheckingOut(true)
    const { data, error, response } = await checkout()
    if (!response.ok || !data) {
      if (error) {
        toast.error(error.title ?? 'Error checking out')
      } else {
        toast.error('Error checking out')
      }
      setCheckingOut(false)
      return
    }

    router.push(data.url)

    setCheckingOut(false)
    setCheckoutSuccess(true)
  }

  if (loading) {
    return <Skeleton className='h-96' />
  }

  return (
    <>
      <CartBox items={cart} onDeleteItem={handleDeleteItem} />
      <Dialog
        open={dialogOpen}
        onOpenChange={(isOpen) => {
          if (!checkingOut) {
            setDialogOpen(isOpen)
          }
        }}
      >
        <DialogTrigger asChild>
          <Button
            variant='filled'
            className='ml-auto'
            disabled={cart.length === 0}
          >
            ชำระเงิน
          </Button>
        </DialogTrigger>
        <DialogContent
          className='w-11/12 rounded-lg'
          aria-describedby={undefined}
        >
          <DialogHeader className='text-left'>
            <DialogTitle asChild>
              <Text variant='heading-md'>
                {checkingOut
                  ? 'กำลังชำระเงิน'
                  : checkoutSuccess
                    ? 'ชำระเงินสำเร็จ'
                    : 'กรุณาตรวจสอบข้อมูล'}
              </Text>
            </DialogTitle>
          </DialogHeader>

          {checkingOut ? (
            <div className='flex flex-col items-center justify-center'>
              <Icon
                icon='eos-icons:loading'
                className='animate-spin text-4xl text-primary-500'
              />
              <Text className='mt-4'>กำลังชำระเงิน...</Text>
            </div>
          ) : checkoutSuccess ? (
            <div className='flex flex-col items-center justify-center'>
              <Icon
                icon='mdi:check-circle'
                className='text-6xl text-success-500'
              />
              <Text className='mt-4 text-lg'>การชำระเงินเสร็จสิ้น</Text>
            </div>
          ) : (
            <>
              <div className='flex flex-col'>
                <Text variant='md-semibold'>ที่อยู่สำหรับการจัดส่ง</Text>
                <Text>
                  {user?.first_name} {user?.last_name}
                </Text>
                <Text>
                  {user?.address.house_no} {user?.address.district}{' '}
                  {user?.address.province}
                  {user?.address.postcode}
                </Text>
                <Text>{user?.tel}</Text>
              </div>
              <div className='flex w-full items-center justify-between'>
                <Text variant='md-semibold'>รายการสินค้า</Text>
                <Text variant='xl-semibold' className='text-primary-500'>
                  ฿{' '}
                  {cart
                    .reduce((acc, item) => acc + (item.art_toy?.price ?? 0), 0)
                    .toLocaleString()}
                </Text>
              </div>
              <Button
                variant='filled'
                className='ml-auto'
                onClick={handleCheckout}
                disabled={checkingOut}
              >
                ชำระเงิน
              </Button>
            </>
          )}
        </DialogContent>
      </Dialog>
    </>
  )
}
