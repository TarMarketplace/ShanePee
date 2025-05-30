import { Icon } from '@iconify/react/dist/iconify.js'
import { AvatarImage } from '@radix-ui/react-avatar'

import { Avatar, AvatarFallback } from '@/components/avatar'
import { Text } from '@/components/text'

import type { ReviewWithTruncatedBuyer } from '@/generated/api'

export interface SellerReviewCardProps {
  review: ReviewWithTruncatedBuyer
  photo: string
  sellerName: string
}

export function SellerReviewCard({
  review,
  photo,
  sellerName,
}: SellerReviewCardProps) {
  return (
    <div className='flex aspect-[250/140] size-full min-w-[180px] max-w-[250px] flex-col rounded-xl p-2 shadow md:aspect-[250/320]'>
      <div className='flex h-10 w-full items-center gap-2 p-1'>
        <Avatar className='size-8'>
          <AvatarImage src={photo} alt={sellerName} />
          <AvatarFallback>{sellerName}</AvatarFallback>
        </Avatar>
        <div className='w-full'>{sellerName}</div>
      </div>
      <div className='flex h-full flex-col justify-between truncate p-2'>
        <Text
          className='truncate whitespace-normal break-words'
          variant='sm-regular'
        >
          {review.comment}
        </Text>
      </div>
      <div className='flex flex-col items-center justify-center truncate p-4'>
        <div className='flex gap-1'>
          {[...Array(5)].map((_, index) => {
            if (index < review.rating) {
              return (
                <Icon
                  key={index}
                  icon='iconoir:star-solid'
                  className='size-6 text-warning-500'
                />
              )
            } else {
              return (
                <Icon
                  key={index}
                  icon='iconoir:star'
                  className='size-6 text-warning-500'
                />
              )
            }
          })}
        </div>
      </div>
    </div>
  )
}
