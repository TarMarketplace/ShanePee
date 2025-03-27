'use client'

import Link from 'next/link'
import { useEffect, useState } from 'react'
import { toast } from 'sonner'

import { Text } from '@/components/text'

import type { ReviewWithTruncatedBuyer } from '@/generated/api'
import { getReviewsOfSeller } from '@/generated/api'

import { SellerReviewCard } from '../_components/seller-review-card'

type SellerReviewContainerProps = {
  sellerId: string
}

export function SellerReviewContainer({
  sellerId,
}: SellerReviewContainerProps) {
  const [reviews, setReviews] = useState<ReviewWithTruncatedBuyer[]>([])

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
      <div className='flex flex-col gap-3 sm:grid sm:grid-cols-[repeat(2,minmax(0,1fr))] sm:p-3 md:grid-cols-[repeat(3,minmax(0,1fr))] lg:grid-cols-[repeat(4,minmax(0,1fr))]'>
        {reviews.map((review) => {
          return (
            <SellerReviewCard
              key={review.buyer_truncated_first_name}
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
