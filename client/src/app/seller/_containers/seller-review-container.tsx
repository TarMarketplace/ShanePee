'use client'

import Link from 'next/link'
import { useEffect, useState } from 'react'
import { toast } from 'sonner'

import { Text } from '@/components/text'

import type { ArrayResponseReview } from '@/generated/api'
import { getReviewsOfSeller } from '@/generated/api'

import { SellerReviewCard } from '../_components/seller-review-card'

type SellerReviewContainerProps = {
  sellerId: string
}

export function SellerReviewContainer({
  sellerId,
}: SellerReviewContainerProps) {
  const [reviews, setReviews] = useState<ArrayResponseReview>({
    data: [],
  })

  useEffect(() => {
    getReviewsOfSeller({
      path: {
        sellerID: parseInt(sellerId),
      },
    })
      .then((response) => {
        if (response?.data) {
          setReviews(response.data)
        } else {
          toast.error('Something went wrong')
        }
      })
      .catch(() => {
        toast.error('Something went wrong')
      })
  }, [sellerId])

  return (
    <div className='my-4 w-full max-w-5xl'>
      <div className='flex justify-between py-3'>
        <div className='flex items-center'>
          <Text variant='heading-md'>รีวิวร้านค้า</Text>
        </div>
        <Link href={'/seller/review/' + sellerId}>
          <Text variant='sm-regular' className='text-primary underline'>
            ดูรีวิวทั้งหมด
          </Text>
        </Link>
      </div>
      <div className='flex gap-3 overflow-x-auto p-4'>
        {reviews.data?.map((review) => {
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
