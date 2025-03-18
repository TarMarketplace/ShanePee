'use client'

import { Icon } from '@iconify/react'
import { useState } from 'react'
import { toast } from 'sonner'

import { Button } from '@/components/button'

import { cn } from '@/lib/utils'

import type { ArtToy} from '@/generated/api';
import { addItemToCart } from '@/generated/api'

interface AddToCartButtonProps {
  product: ArtToy
}

export function AddToCartButton({ product }: AddToCartButtonProps) {
  const [loading, setLoading] = useState(false)

  const handleAddToCart = () => {
    setLoading(true)
    toast.promise(
      addItemToCart({
        body: {
          art_toy_id: product.id,
        },
      }),
      {
        loading: `Adding ${product.name} to cart...`,
        success: `Added ${product.name} to cart!`,
        error: `Failed to add ${product.name} to cart`,
        finally: () => setLoading(false),
      }
    )
  }

  return (
    <Button
      variant='filled'
      disabled={!product.availability || loading}
      onClick={handleAddToCart}
    >
      <Icon
        icon={loading ? 'fa6-solid:spinner' : 'tdesign:cart-filled'}
        className={cn('size-4', loading && 'animate-spin')}
      />
      {product.availability ? 'เพิ่มในตะกร้า' : 'ขายแล้ว'}
    </Button>
  )
}
