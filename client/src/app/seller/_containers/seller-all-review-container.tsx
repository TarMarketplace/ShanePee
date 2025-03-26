'use client'

import { Icon } from '@iconify/react/dist/iconify.js'
import { useEffect, useState } from 'react'
import { toast } from 'sonner'

import { Text } from '@/components/text'

import type { Review } from '@/generated/api'
import { getReviewsOfSeller } from '@/generated/api'

import { SellerReviewCard } from '../_components/seller-review-card'

type SellerAllReviewContainerProps = {
  sellerId: string
}

export function SellerAllReviewContainer({
  sellerId,
}: SellerAllReviewContainerProps) {
  const [reviews, setReviews] = useState<Review[]>([])

  useEffect(() => {
    getReviewsOfSeller({
      path: {
        sellerID: parseInt(sellerId),
      },
    })
      .then((response) => {
        if (Array.isArray(response.data?.data)) {
          setReviews(response.data.data)
        } else {
          setReviews([])
          toast.error('No reviews found')
        }
      })
      .catch(() => {
        toast.error('Something went wrong')
      })
  }, [sellerId])

  return (
    <div className='flex w-full max-w-5xl flex-col sm:min-w-[60%] sm:divide-y-4 sm:divide-primary'>
      <div className='flex justify-between py-3'>
        <div className='flex items-center'>
          <Icon icon='material-symbols:list' className='mr-2 size-9' />
          <Text variant='heading-md' desktopVariant='heading-lg'>
            รีวิวร้านค้า
          </Text>
        </div>
      </div>
      <div className='flex flex-col gap-3 sm:grid sm:grid-cols-[repeat(2,minmax(0,1fr))] sm:p-3 md:grid-cols-[repeat(3,minmax(0,1fr))] lg:grid-cols-[repeat(4,minmax(0,1fr))]'>
        {reviews.map((review) => {
          return (
            <SellerReviewCard
              key={review.id}
              review={review} // TODO add photo and sellerName in review response
              photo='data:image/png;base64,mfkirjIDSFIj32asdf...'
              sellerName='John Doe'
            />
          )
        })}
      </div>
    </div>
  )
}
