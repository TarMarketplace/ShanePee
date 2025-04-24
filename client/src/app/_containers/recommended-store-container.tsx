'use client'

import { Icon } from '@iconify/react'
import { useEffect, useState } from 'react'

import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import { type UserWithReview, getSellers } from '@/generated/api'

import { RecommendedStore } from '../_components/recommended-store'

export function RecommendedStoreContainer() {
  const { user } = useUser()
  const [recommendedSellers, setRecommendedSellers] = useState<
    UserWithReview[]
  >([])

  useEffect(() => {
    getSellers().then((response) => {
      const sellers = response?.data?.data
      if (sellers) {
        setRecommendedSellers(
          sellers
            .filter((seller: UserWithReview) => seller.id !== user?.id)
            .slice(0, 4)
        )
      }
    })
  }, [user])

  return (
    <div className='flex w-full flex-col gap-3.5 rounded-xl bg-white px-5 pb-4 pt-5 shadow-md'>
      <div className='flex items-center gap-3'>
        <Icon icon='ri:store-3-fill' className='size-9 text-primary' />
        <Text variant='heading-sm' desktopVariant='heading-md'>
          ร้านค้าแนะนำ
        </Text>
      </div>
      <div className='grid grid-cols-2 gap-3'>
        {recommendedSellers.map((seller, index) => (
          <RecommendedStore
            key={index}
            name={`${seller.first_name} ${seller.last_name}`}
            average_rating={seller.rating}
            review_count={seller.number_of_reviews}
            image_src={seller.photo ?? ''}
            seller_id={seller.id}
          />
        ))}
      </div>
    </div>
  )
}
