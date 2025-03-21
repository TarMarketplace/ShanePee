'use client'

import Link from 'next/link'
import { useEffect, useState } from 'react'
import { toast } from 'sonner'

import { type ArtToy, getMyArtToys } from '@/generated/api'

import { SellerProductCard } from '../_components/seller-product-card'

const getMyProducts = async () => {
  const { data, response } = await getMyArtToys({
    cache: 'no-cache',
  })

  if (response.ok) {
    return data
  } else {
    return null
  }
}

export function ProductContainer() {
  const [products, setProducts] = useState<ArtToy[] | null>(null)

  useEffect(() => {
    getMyProducts()
      .then((response) => {
        if (response?.data) {
          setProducts(response.data)
        } else {
          toast.error('Something went wrong')
        }
      })
      .catch(() => {
        toast.error('Something went wrong')
      })
  }, [])

  return (
    <div className='flex flex-col gap-3 sm:grid sm:grid-cols-[repeat(2,minmax(0,1fr))] sm:p-3 md:grid-cols-[repeat(3,minmax(0,1fr))] lg:grid-cols-[repeat(4,minmax(0,1fr))]'>
      {products?.map((product) => {
        return (
          <Link key={product.id} href={`/product/edit/${product.id}`}>
            <SellerProductCard key={product.id} product={product} />
          </Link>
        )
      })}
    </div>
  )
}
