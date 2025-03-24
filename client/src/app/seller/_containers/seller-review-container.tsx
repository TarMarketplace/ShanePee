'use client'

import Link from 'next/link'
import { useEffect, useState } from 'react'

import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import { type Review } from '@/generated/api'

import { SellerReviewCard } from '../_components/seller-review-card'

type SellerReviewContainerProps = {
  sellerId: string
}

export function SellerReviewContainer({
  sellerId,
}: SellerReviewContainerProps) {
  const [reviews, setReviews] = useState<Review[] | null>(null)
  const { user } = useUser()

  useEffect(() => {
    // if (user?.id) {
    //   getSellerReviews(user.id)
    //     .then((response) => {
    //       if (response?.data) {
    //         setReviews(response.data)
    //       } else {
    //         toast.error('Something went wrong')
    //       }
    //     })
    //     .catch(() => {
    //       toast.error('Something went wrong')
    //     })
    // }
    // TODO get seller review api
    setReviews([
      {
        art_toy_id: 4279569719,
        comment:
          'การบริหารดีมากครับ 10 เต็ม 10 ไปเลย แต่ว่า ผมพบปัญหานิดหน่อยครับ แต่ไม่มั่นใจว่าสามารถใช้ Binary Search ในการหาคำตอบได้มั้ย ...',
        id: 9007199254740991,
        rating: 5,
      },
      {
        art_toy_id: 4279569719,
        comment:
          'การบริหารดีมากครับ 10 เต็ม 10 ไปเลย แต่ว่า ผมพบปัญหานิดหน่อยครับ แต่ไม่มั่นใจว่าสามารถใช้ Binary Search ในการหาคำตอบได้มั้ย ...',
        id: 9007199254740991,
        rating: 2,
      },
      {
        art_toy_id: 4279569719,
        comment:
          'การบริหารดีมากครับ 10 เต็ม 10 ไปเลย แต่ว่า ผมพบปัญหานิดหน่อยครับ แต่ไม่มั่นใจว่าสามารถใช้ Binary Search ในการหาคำตอบได้มั้ย ...',
        id: 9007199254740991,
        rating: 3,
      },
      {
        art_toy_id: 4279569719,
        comment:
          'การบริหารดีมากครับ 10 เต็ม 10 ไปเลย แต่ว่า ผมพบปัญหานิดหน่อยครับ แต่ไม่มั่นใจว่าสามารถใช้ Binary Search ในการหาคำตอบได้มั้ย ...',
        id: 9007199254740991,
        rating: 3,
      },
      {
        art_toy_id: 4279569719,
        comment:
          'การบริหารดีมากครับ 10 เต็ม 10 ไปเลย แต่ว่า ผมพบปัญหานิดหน่อยครับ แต่ไม่มั่นใจว่าสามารถใช้ Binary Search ในการหาคำตอบได้มั้ย ...',
        id: 9007199254740991,
        rating: 3,
      },
    ])
  }, [user])

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
        {reviews?.map((review) => {
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
