import type { ArtToy } from '@/generated/api'

import { SellerProductCardDesktop } from './seller-product-card-desktop'
import { SellerProductCardMobile } from './seller-product-card-mobile'

export interface SellerProductCardProps {
  product: ArtToy
}

export function SellerProductCard(props: SellerProductCardProps) {
  return (
    <>
      <div className='hidden sm:flex'>
        <SellerProductCardDesktop product={props.product} />
      </div>
      <div className='flex sm:hidden'>
        <SellerProductCardMobile product={props.product} />
      </div>
    </>
  )
}
