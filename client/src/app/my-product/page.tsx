import { Icon } from '@iconify/react'
import Link from 'next/link'

import { Button } from '@/components/button'
import { Text } from '@/components/text'

import { ProductContainer } from './_containers/product-container'

export default function MyProductPage() {
  return (
    <main className='grid size-full grid-cols-1 place-items-center p-4 md:p-12'>
      <div className='flex w-full flex-col sm:w-fit sm:min-w-[60%] sm:divide-y-4 sm:divide-primary'>
        <div className='flex justify-between py-3'>
          <div className='flex items-center'>
            <Icon icon='material-symbols:list' className='size-9' />
            <Text variant='heading-md' desktopVariant='heading-lg'>
              สินค้าของคุณ
            </Text>
          </div>
          <Link href='/product/create'>
            <Button>
              <Icon icon='ic:round-plus' />
              <Text variant='sm-semibold' desktopVariant='lg-semibold'>
                เพิ่มสินค้า
              </Text>
            </Button>
          </Link>
        </div>
        <ProductContainer />
      </div>
    </main>
  )
}
