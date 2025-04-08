'use client'

import { Icon } from '@iconify/react'
import Image from 'next/image'

import { Button } from '@/components/button'
import { Text } from '@/components/text'

export interface ReviewCardProps {
  name: string
  photo: string
  rating: number
  comment: string
  setRating: (val: number) => void
  setComment: (val: string) => void
  onSubmit: () => Promise<void>
  onClose: () => void
}

const ReviewCard = ({
  name,
  photo,
  rating,
  comment,
  setRating,
  setComment,
  onSubmit,
  onClose,
}: ReviewCardProps) => {
  return (
    <div className='w-[353px] space-y-5 rounded-lg bg-[#FCFBF7] p-4 shadow-lg md:w-[500px]'>
      <div className='flex items-center justify-between'>
        <Text variant='heading-md'>ให้คะแนนร้านค้า</Text>
        <button onClick={onClose}>
          <Icon icon='maki:cross' className='size-[16px] text-[#8E8E8E]' />
        </button>
      </div>

      <div className='flex h-[76px] w-full items-center gap-4 rounded-lg bg-[#F4F4F4] p-2 shadow-[0_0_4px_0px_#00000040] md:h-[96px] md:p-3'>
        <div className='relative size-[60px] md:size-[80px]'>
          <Image
            src={photo}
            alt='Shop photo'
            fill
            className='rounded-full object-cover'
          />
        </div>
        <Text variant='md-semibold'>{name}</Text>
      </div>

      <div className='flex-col space-y-2'>
        <Text variant='md-regular'>ให้คะแนนร้านค้า</Text>
        <div className='flex gap-1'>
          {[1, 2, 3, 4, 5].map((star) => (
            <Icon
              key={star}
              icon={star <= rating ? 'iconoir:star-solid' : 'iconoir:star'}
              className='size-8 cursor-pointer text-[#FFC107]'
              strokeWidth={star <= rating ? 0 : 2}
              onClick={() => setRating(star)}
            />
          ))}
        </div>

        <Text variant='md-regular'>รีวิวร้านค้า</Text>
        <textarea
          value={comment}
          onChange={(e) => setComment(e.target.value)}
          className='w-full rounded-lg border bg-[#FCFBF7] p-2'
          placeholder='พิมพ์ข้อความรีวิว...'
          rows={5}
        />
      </div>

      <div className='flex justify-end pt-2'>
        <Button variant='filled' className='text-white' onClick={onSubmit}>
          ให้คะแนน
        </Button>
      </div>
    </div>
  )
}

export { ReviewCard }
