'use client'

import { useEffect, useState } from 'react'
import { toast } from 'sonner'

import { Button } from '@/components/button'
import { ProductCard } from '@/components/product-card'
import { Skeleton } from '@/components/skeleton'

import type { ArtToy } from '@/generated/api'
import { getArtToys } from '@/generated/api'

export function ProductListContainer() {
  const [products, setProducts] = useState<ArtToy[]>([])
  const [loading, setLoading] = useState(true)
  const [showing, setShowing] = useState(8)

  useEffect(() => {
    const fetchProducts = async () => {
      const { data, error } = await getArtToys()

      if (error) {
        toast.error(error.title ?? 'Cannot get products')
      }

      setProducts(data?.data?.filter((product) => product.availability) ?? [])
      setLoading(false)
    }

    fetchProducts()
  }, [])

  return (
    <div className='grid size-full grid-cols-2 place-items-center gap-4 md:grid-cols-4'>
      {loading
        ? Array.from({ length: 8 }).map((_, index) => (
            <Skeleton
              key={index}
              className='aspect-[170/240] size-full max-w-64'
            />
          ))
        : products
            .slice(0, showing)
            .map((product) => (
              <ProductCard key={product.id} product={product} />
            ))}
      {products.length > showing && (
        <Button onClick={() => setShowing(showing + 8)} className='col-span-4'>
          ดูเพิ่มเติม
        </Button>
      )}
    </div>
  )
}
