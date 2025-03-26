'use client'

import { useRouter } from 'next/navigation'
import { useMemo, useState } from 'react'
import { toast } from 'sonner'

import { useCart } from '@/providers/cart-provider'
import { useUser } from '@/providers/user-provider'

import type { ArtToy } from '@/generated/api'
import {
  addItemToCart,
  clearItemsFromCart,
  removeItemFromCart,
} from '@/generated/api'

import { ChangeSellerDialog } from '../_components/change-seller-dialog'
import { ProductPage } from '../_components/product-page'

interface ProductPageProps {
  product: ArtToy | null | undefined
}

export function ProductPageContainer({ product }: ProductPageProps) {
  const router = useRouter()
  const { user } = useUser()
  const { cartItems, fetchCart } = useCart()
  const [showDialog, setShowDialog] = useState(false)
  const [cartButtonLoading, setCartButtonLoading] = useState(false)

  const isInCart = useMemo(() => {
    return product
      ? cartItems.some((item) => item.art_toy_id === product.id)
      : false
  }, [cartItems, product])
  const isDifferentSeller = useMemo(() => {
    return (
      product &&
      cartItems.length > 0 &&
      cartItems[0].art_toy?.owner_id !== product.owner_id
    )
  }, [cartItems, product])

  const getOwnerName = (ownerId: number | undefined) => {
    // TODO: get owner name from owner id
    return ownerId?.toString() || ''
  }

  if (!product) return <p>Product not found.</p>

  const handleCartButton = async () => {
    if (!user) {
      router.push('/login')
      toast.error('Please login to continue')
      return
    }
    setCartButtonLoading(true)
    if (isInCart) {
      const cartItem = cartItems.find((item) => item.art_toy_id === product.id)
      if (cartItem?.id) {
        const { response } = await removeItemFromCart({
          path: { id: cartItem.id },
        })
        if (response.ok) toast.success('Removed from cart')
        else toast.error('Failed to remove from cart')
      }
    } else {
      if (isDifferentSeller) {
        setCartButtonLoading(false)
        setShowDialog(true)
        return
      }
      const { response } = await addItemToCart({
        body: { art_toy_id: product.id },
      })
      console.log(response)
      if (response.ok) toast.success('Added to cart successfully')
      else toast.error('Failed to add to cart')
    }
    setCartButtonLoading(false)
    fetchCart()
  }

  const handleConfirmDifferentSeller = async () => {
    await clearItemsFromCart()
    const { response } = await addItemToCart({
      body: { art_toy_id: product.id },
    })
    if (response.ok) toast.success('Cart updated with new seller')
    else toast.error('Failed to update cart')
    setShowDialog(false)
    fetchCart()
    setCartButtonLoading(false)
  }

  return (
    <>
      <ProductPage
        product={product}
        handleCartButton={handleCartButton}
        isInCart={isInCart}
        cartButtonLoading={cartButtonLoading}
      />
      <ChangeSellerDialog
        sellerNameFrom={getOwnerName(cartItems[0]?.art_toy?.owner_id)}
        sellerNameTo={getOwnerName(product.owner_id)}
        showDialog={showDialog}
        setShowDialog={setShowDialog}
        handleConfirmDifferentSeller={handleConfirmDifferentSeller}
      />
    </>
  )
}
