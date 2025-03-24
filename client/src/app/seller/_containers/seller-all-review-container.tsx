'use client'

import { Icon } from '@iconify/react/dist/iconify.js'
import { useEffect, useState } from 'react'

import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import { type Review } from '@/generated/api'

import { SellerReviewCard } from '../_components/seller-review-card'

// type SellerAllReviewContainerProps = {
//   sellerId: string
// }

export function SellerAllReviewContainer() {
  // {
  //   sellerId,
  // }: SellerAllReviewContainerProps
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
