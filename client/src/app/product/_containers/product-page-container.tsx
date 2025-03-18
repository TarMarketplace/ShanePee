'use client'

import { Icon } from '@iconify/react/dist/iconify.js'
import { useRouter } from 'next/navigation'
import { useState } from 'react'
import { toast } from 'sonner'

import { Button } from '@/components/button'
import { Text } from '@/components/text'

import { useCart } from '@/providers/cart-provider'
import { useUser } from '@/providers/user-provider'

import type { ArtToy } from '@/generated/api'
import { addItemToCart, removeItemFromCart } from '@/generated/api'

import { ProductPage } from '../_components/product-page'

interface ProductPageProps {
  product: ArtToy | null | undefined
}

export function ProductPageContainer({ product }: ProductPageProps) {
  const router = useRouter()
  const { user } = useUser()
  const { cartItems, fetchCart } = useCart()
  const [showPopover, setShowPopover] = useState(false)

  if (!product) return <p>Product not found.</p>

  const isInCart = cartItems.some((item) => item.art_toy_id === product.id)
  const isDifferentSeller =
    cartItems.length > 0 && cartItems[0].owner_id !== product.owner_id

  const getOwnerName = (ownerId: number) => {
    // TODO: get owner name from owner id
    return ownerId
  }

  const handleCartButton = async () => {
    if (!user) {
      router.push('/login')
      toast.error('Please login to continue')
      return
    }
    if (isInCart) {
      const cartItem = cartItems.find((item) => item.art_toy_id === product.id)
      if (cartItem?.id) {
        await removeItemFromCart({
          path: { id: cartItem.id },
        })
        toast.success('Removed from cart')
      }
    } else {
      if (isDifferentSeller) {
        setShowPopover(true)
        return
      }
      await addItemToCart({ body: { art_toy_id: product.id } })
      toast.success('Added to cart successfully')
    }
    fetchCart()
  }

  const handleConfirmDifferentSeller = async () => {
    // TODO: BE API to clear cart
    for (const item of cartItems) {
      await removeItemFromCart({ path: { id: item.id } })
    }
    await addItemToCart({ body: { art_toy_id: product.id } })
    toast.success('Cart updated with new seller')
    setShowPopover(false)
    fetchCart()
  }

  return (
    <>
      <ProductPage
        product={product}
        handleCartButton={handleCartButton}
        isInCart={isInCart}
      />
      {showPopover && (
        <>
          <div
            className='fixed inset-0 bg-black/20'
            onClick={() => setShowPopover(false)}
          ></div>
          <div className='fixed left-1/2 top-1/2 flex w-[354px] -translate-x-1/2 -translate-y-1/2 flex-col rounded bg-white p-4 shadow-lg md:w-[500px]'>
            <div className='mb-4 flex items-center'>
              <Text variant='heading-md' className='w-full'>
                ต้องการจะเปลี่ยนร้านหรือไม่
              </Text>
              <Icon
                icon='maki:cross'
                className='size-5 cursor-pointer text-grey-500'
                onClick={() => setShowPopover(false)}
              />
            </div>
            <div className='flex flex-wrap gap-1'>
              <Text variant='lg-regular' className=''>
                คุณกำลังจะเปลี่ยนร้านจาก
                <span className='mx-1 font-bold'>
                  {getOwnerName(cartItems[0]?.owner_id)}
                </span>
                เป็น
                <span className='mx-1 font-bold'>
                  {getOwnerName(product.owner_id)}
                </span>
                <br />
                หากคุณเปลี่ยนร้าน <br className='md:hidden' />
                <span className='text-primary underline'>
                  สินค้าในตะกร้าทั้งหมดของคุณจะหายไป!!
                </span>
                กรุณาตรวจสอบสินค้าในตะกร้าก่อนยืนยัน
              </Text>
            </div>
            <div className='flex justify-end'>
              <Button variant='filled' onClick={handleConfirmDifferentSeller}>
                เปลี่ยนร้าน
              </Button>
            </div>
          </div>
        </>
      )}
    </>
  )
}
