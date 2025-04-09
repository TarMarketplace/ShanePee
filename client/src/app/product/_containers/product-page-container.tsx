'use client'

import { useRouter } from 'next/navigation'
import { useCallback, useEffect, useMemo, useState } from 'react'
import { toast } from 'sonner'

import { useCart } from '@/providers/cart-provider'
import { useUser } from '@/providers/user-provider'

import type { ArtToy, UserWithReview } from '@/generated/api'
import {
  addItemToCart,
  clearItemsFromCart,
  getSellerById,
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
  const [sellerNameFrom, setSellerNameFrom] = useState('')
  const [seller, setSeller] = useState<UserWithReview>()

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

  const getSellerNameFrom = useCallback(
    async (ownerId: number) => {
      await getSellerById({
        path: {
          id: ownerId,
        },
      })
        .then((response) => {
          if (response.response.ok) {
            setSellerNameFrom(
              response.data?.first_name + ' ' + response.data?.last_name
            )
          } else if (response.response.status == 401) {
            router.push('/login')
          } else {
            toast.error('Something went wrong')
          }
        })
        .catch(() => {
          toast.error('Something went wrong')
        })
    },
    [router]
  )

  const fetchSeller = useCallback(
    async (ownerId: number) => {
      await getSellerById({
        path: {
          id: ownerId,
        },
      })
        .then((response) => {
          if (response.response.ok) {
            setSeller(response?.data)
          } else if (response.response.status == 401) {
            router.push('/login')
          } else {
            toast.error('Something went wrong')
          }
        })
        .catch(() => {
          toast.error('Something went wrong')
        })
    },
    [router]
  )

  useEffect(() => {
    if (product) {
      fetchSeller(product.owner_id)
    }
    getSellerNameFrom(cartItems[0]?.art_toy?.owner_id ?? 0)
  }, [product, cartItems, fetchSeller, getSellerNameFrom])

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

  const handleGotoStore = () => {
    router.push(`/seller/${product.owner_id}`)
  }

  const handleGotoChat = () => {
    router.push(`/chat?id=${product.owner_id}`)
  }

  return (
    <>
      <ProductPage
        product={product}
        seller={seller}
        onClickCartButton={handleCartButton}
        onGotoStore={handleGotoStore}
        onGotoChat={handleGotoChat}
        isInCart={isInCart}
        cartButtonLoading={cartButtonLoading}
      />
      <ChangeSellerDialog
        sellerNameFrom={sellerNameFrom}
        sellerNameTo={seller?.first_name + ' ' + seller?.last_name}
        showDialog={showDialog}
        setShowDialog={setShowDialog}
        handleConfirmDifferentSeller={handleConfirmDifferentSeller}
      />
    </>
  )
}
