'use client'

import { useUser } from '@/providers/user-provider'

import { SellerStatsCard } from '../_components/seller-stats-card'

type SellerStatsContainerProps = {
  sellerId: string
}

export function SellerStatsContainer({ sellerId }: SellerStatsContainerProps) {
  const { user } = useUser()
  const isSeller = user?.id.toString() == sellerId

  // useEffect(() => {
  //   getMyStats()
  //   .then((response) => {
  //     if (response?.data) {
  //       setProducts(response.data)
  //     } else {
  //       toast.error('Something went wrong')
  //     }
  //   })
  //   .catch(() => {
  //     toast.error('Something went wrong')
  //   })
  // }, [])
  // TODO get seller stat api
  return (
    <div className='w-full max-w-5xl md:my-4'>
      <SellerStatsCard
        photo='asdf'
        sellerName='John Doe'
        store_rating={4.5}
        total_art_toys={7}
        sold_art_toys={69}
        total_review={120}
        joined_for={70}
        isSeller={isSeller}
      />
    </div>
  )
}
