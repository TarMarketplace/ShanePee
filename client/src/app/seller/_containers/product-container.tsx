'use client'

import { Icon } from '@iconify/react/dist/iconify.js'
import Link from 'next/link'
import { useEffect, useState } from 'react'
import { toast } from 'sonner'

import { Button } from '@/components/button'
import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import { type ArtToy, getArtToysOfSeller, getMyArtToys } from '@/generated/api'

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

type ProductContainerProps = {
  sellerId: string
}

export function ProductContainer({ sellerId }: ProductContainerProps) {
  const [products, setProducts] = useState<ArtToy[] | null>(null)
  const { user } = useUser()
  const isSeller = user?.id.toString() == sellerId

  useEffect(() => {
    if (isSeller) {
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
    } else {
      getArtToysOfSeller({
        path: {
          id: parseInt(sellerId),
        },
      })
        .then((response) => {
          if (response?.data) {
            setProducts(response.data.data)
          } else {
            toast.error('Something went wrong')
          }
        })
        .catch(() => {
          toast.error('Something went wrong')
        })
    }
  }, [isSeller, sellerId])

  return (
    <div className='flex w-full max-w-5xl flex-col sm:min-w-[60%] sm:divide-y-4 sm:divide-primary'>
      <div className='flex justify-between py-3'>
        <div className='flex items-center'>
          <Icon icon='material-symbols:list' className='mr-2 size-9' />
          <Text variant='heading-md' desktopVariant='heading-lg'>
            {isSeller ? 'สินค้าของคุณ' : 'สินค้าทั้งหมด'}
          </Text>
        </div>
        {isSeller && (
          <Link href='/product/create'>
            <Button>
              <Icon icon='ic:round-plus' />
              <Text variant='sm-semibold' desktopVariant='lg-semibold'>
                เพิ่มสินค้า
              </Text>
            </Button>
          </Link>
        )}
      </div>
      <div className='flex flex-col gap-3 sm:grid sm:grid-cols-[repeat(2,minmax(0,1fr))] sm:p-3 md:grid-cols-[repeat(3,minmax(0,1fr))] lg:grid-cols-[repeat(4,minmax(0,1fr))]'>
        {products?.map((product) => {
          return (
            <Link
              key={product.id}
              href={
                isSeller
                  ? `/product/edit/${product.id}`
                  : `/product/${product.id}`
              }
            >
              <SellerProductCard key={product.id} product={product} />
            </Link>
          )
        })}
      </div>
    </div>
  )
}
