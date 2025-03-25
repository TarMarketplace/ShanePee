import { Icon } from '@iconify/react'
import Image from 'next/image'
import { useState } from 'react'

import { Button } from '@/components/button'
import { Text } from '@/components/text'

export interface ReviewCardProps {
  Name: string
  Photo: string
  orderID: number
  onSubmit: (rating: number, comment: string, orderID: number) => Promise<void>
  onClose: () => void
}

const ReviewCard = ({
  Name,
  Photo,
  orderID,
  onSubmit,
  onClose,
}: ReviewCardProps) => {
  const [rating, setRating] = useState(0)
  const [comment, setComment] = useState('')

  const handleRatingClick = (value: number) => setRating(value)

  const handleSubmit = async () => {
    if (rating === 0) {
      alert('Please select a rating')
      return
    }
    await onSubmit(rating, comment, orderID)
  }

  return (
    <div className='absolute w-[353px] space-y-5 rounded-lg bg-[#FCFBF7] p-4 shadow-lg md:w-[500px]'>
      {/* Cancel Button */}
      <div className='flex justify-between'>
        <Text variant='heading-md'>ให้คะแนนร้านค้า</Text>

        <button onClick={onClose}>
          <Icon
            icon='maki:cross'
            className='size-[16px] text-[#8E8E8E]'
          />
        </button>
      </div>

      <div className='flex h-[76px] w-[321px] items-center gap-4 rounded-lg bg-[#F4F4F4] p-2 shadow-[0_0_4px_0px_#00000040] md:h-[96px] md:w-[468px] md:p-3'>
        <Image
          src={Photo}
          alt='Shop photo'
          className='size-[60px] rounded-full object-cover md:size-[80px]'
        />
        <Text variant='md-semibold'>{Name}</Text>
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
              onClick={() => handleRatingClick(star)}
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
        <Button variant='filled' className='text-white' onClick={handleSubmit}>
          ให้คะแนน
        </Button>
      </div>
    </div>
  )
}

export { ReviewCard }
