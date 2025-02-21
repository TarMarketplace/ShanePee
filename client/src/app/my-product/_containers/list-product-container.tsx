'use client'

import { Icon } from '@iconify/react'
import Link from 'next/link'

import { Button } from '@/components/button'
import { Text } from '@/components/text'

import { ProductContainer } from './product-container'

export async function ListProductContainer() {
  return (
    <div className='flex min-w-[60%] flex-col divide-y-4 divide-primary'>
      <div className='flex justify-between py-3'>
        <div className='flex items-center'>
          <Icon icon='material-symbols:list' className='size-9' />
          <Text variant='heading-lg'>สินค้าของคุณ</Text>
        </div>
        <Link href='/product/create'>
          <Button>
            <Icon icon='ic:round-plus' />
            <Text variant='lg-semibold'>เพิ่มสินค้า</Text>
          </Button>
        </Link>
      </div>
      <div>
        <ProductContainer />
      </div>
    </div>
  )
}
