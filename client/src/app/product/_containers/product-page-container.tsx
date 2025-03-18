'use client'

import { useRouter } from 'next/navigation'
import { toast } from 'sonner'

import { useCart } from '@/providers/cart-provider'
import { useUser } from '@/providers/user-provider'

import type { ArtToy} from '@/generated/api';
import { addItemToCart, removeItemFromCart } from '@/generated/api'

import { ProductPage } from '../_components/product-page'

interface ProductPageProps {
  product: ArtToy | null | undefined
}

export function ProductPageContainer({ product }: ProductPageProps) {
  const router = useRouter()
  const { user } = useUser()
  const { cartItems, fetchCart } = useCart()

  if (!product) return <p>Product not found.</p>

  const isInCart = cartItems.some((item) => item.art_toy_id === product.id)
  const isDifferentSeller =
    cartItems.length > 0 && cartItems[0].owner_id !== product.owner_id

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
        // TODO Handle different seller
        console.log('Handle different seller')
      }
      await addItemToCart({ body: { art_toy_id: product.id } })
      toast.success('Added to cart successfully')
    }
    fetchCart()
  }

  return (
    <ProductPage
      product={product}
      handleCartButton={handleCartButton}
      isInCart={isInCart}
    />
  )
}
