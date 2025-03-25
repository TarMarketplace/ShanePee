import { Icon } from '@iconify/react'

import { Text } from '@/components/text'

import { AddressContainer } from '../_containers/address-container'
import { CartContainer } from './_containers/cart-container'

export default async function CartPage() {
  return (
    <main className='grid size-full grid-cols-1 place-items-center p-4 md:p-12'>
      <div className='flex w-full flex-col sm:w-fit sm:min-w-[60%] sm:divide-y-4 sm:divide-primary'>
        <div className='flex items-center gap-2 py-3'>
          <Icon icon='tdesign:cart-filled' className='size-6 md:size-9' />
          <Text variant='heading-md' desktopVariant='heading-lg'>
            สรุปรายการคำสั่งซื้อ
          </Text>
        </div>
        <div className='flex flex-col gap-5 py-2 md:gap-6 md:p-6'>
          <AddressContainer enablePencil />
          <CartContainer />
        </div>
      </div>
    </main>
  )
}
