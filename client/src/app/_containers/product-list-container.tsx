'use client'

import { useEffect, useState } from 'react'
import { toast } from 'sonner'

import { ProductCard } from '@/components/product-card'
import { Skeleton } from '@/components/skeleton'

import type { ArtToy } from '@/generated/api'
import { getArtToys } from '@/generated/api'

export function ProductListContainer() {
  const [products, setProducts] = useState<ArtToy[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchProducts = async () => {
      const { data, error } = await getArtToys()

      if (error) {
        toast.error(error.title ?? 'Cannot get products')
      }

      setProducts(
        data?.data?.filter((product) => product.availability).slice(0, 8) ?? []
      )
      setLoading(false)
    }

    fetchProducts()
  }, [])

  return (
    <div className='grid size-full grid-cols-2 place-items-center gap-4 md:grid-cols-4'>
      {loading
        ? Array.from({ length: 12 }).map((_, index) => (
            <Skeleton
              key={index}
              className='aspect-[170/240] size-full max-w-64'
            />
          ))
        : products.map((product) => (
            <ProductCard key={product.id} product={product} />
          ))}
    </div>
  )
}
