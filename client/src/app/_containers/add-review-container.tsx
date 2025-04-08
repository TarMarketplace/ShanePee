'use client'

import Image from 'next/image'
import { useState } from 'react'
import { toast } from 'sonner'

import { Button } from '@/components/button'
import { ReviewCard } from '@/components/review-card'
import { Text } from '@/components/text'

import { createReview } from '@/generated/api'

interface AddReviewContainerProps {
  name: string
  photo: string
  orderID: number
}

export default function AddReviewContainer({
  name,
  photo,
  orderID,
}: AddReviewContainerProps) {
  const [open, setOpen] = useState(false)
  const [rating, setRating] = useState(0)
  const [comment, setComment] = useState('')

  const handleSubmit = async () => {
    if (rating === 0) {
      toast.error('Please select a rating')
      return
    }

    try {
      const requestBody = { rating, comment }
      console.log('Request Body:', requestBody)

      if (isNaN(orderID)) {
        toast.error('Invalid order ID.')
        return
      }

      await createReview({
        path: { orderID: orderID },
        body: requestBody,
      })

      toast.success('Review successfully submitted.')
      setOpen(false)
    } catch (error) {
      toast.error('Something went wrong.')
      console.error(error)
    }
  }

  return (
    <>
      <Button variant='filled' onClick={() => setOpen(true)}>
        <Image src='/review-button-star.svg' width={24} height={24} alt='' />
        <Text desktopVariant='md-semibold'>ให้คะแนน</Text>
      </Button>

      {open && (
        <div className='fixed inset-0 z-[9999] flex items-center justify-center bg-black/40'>
          <div className='fixed inset-0' onClick={() => setOpen(false)} />
          <div className='relative z-10 flex min-h-screen w-full items-center justify-center'>
            <ReviewCard
              name={name}
              photo={photo}
              rating={rating}
              comment={comment}
              setRating={setRating}
              setComment={setComment}
              onSubmit={handleSubmit}
              onClose={() => setOpen(false)}
            />
          </div>
        </div>
      )}
    </>
  )
}
