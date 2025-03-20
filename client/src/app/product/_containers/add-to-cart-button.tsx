'use client'

import { Icon } from '@iconify/react'

import { Button } from '@/components/button'

import { cn } from '@/lib/utils'

import type { ArtToy } from '@/generated/api'

interface AddToCartButtonProps {
  product: ArtToy
  isInCart: boolean
  handleAddToCart: () => void
  loading: boolean
}

export function AddToCartButton({
  product,
  isInCart,
  handleAddToCart,
  loading,
}: AddToCartButtonProps) {
  return (
    <Button
      variant='filled'
      colorVariant={isInCart ? 'outline' : 'default'}
      disabled={!product.availability || loading}
      onClick={handleAddToCart}
    >
      <Icon
        icon={
          loading
            ? 'fa6-solid:spinner'
            : isInCart
              ? 'mingcute:check-fill'
              : 'tdesign:cart-filled'
        }
        className={cn('size-4', loading && 'animate-spin')}
      />
      {product.availability
        ? isInCart
          ? 'เพิ่มในตะกร้าแล้ว'
          : 'เพิ่มในตะกร้า'
        : 'ขายแล้ว'}
    </Button>
  )
}
