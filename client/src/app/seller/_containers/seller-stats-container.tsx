'use client'

import { useEffect, useState } from 'react'
import { toast } from 'sonner'

import { useUser } from '@/providers/user-provider'

import type { UserWithReview } from '@/generated/api'
import { getSellerById } from '@/generated/api'

import { SellerStatsCard } from '../_components/seller-stats-card'

type SellerStatsContainerProps = {
  sellerId: string
}

export function SellerStatsContainer({ sellerId }: SellerStatsContainerProps) {
  const { user } = useUser()
  const isSeller = user?.id.toString() == sellerId
  const [stats, setStats] = useState<UserWithReview>({
    created_at: '',
    first_name: '',
    id: 0,
    last_name: '',
    number_of_art_toys_sold: 0,
    number_of_reviews: 0,
    photo: '',
    rating: 0,
  })

  useEffect(() => {
    getSellerById({
      path: {
        id: parseInt(sellerId),
      },
    })
      .then((response) => {
        if (response?.data) {
          setStats(response.data)
        } else {
          toast.error('Something went wrong')
        }
      })
      .catch(() => {
        toast.error('Something went wrong')
      })
  }, [sellerId])

  return (
    <div className='w-full max-w-5xl md:my-4'>
      <SellerStatsCard stats={stats} isSeller={isSeller} />
    </div>
  )
}
