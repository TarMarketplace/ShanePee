import { Icon } from '@iconify/react'
import Image from 'next/image'
import Link from 'next/link'

import { Text } from '@/components/text'

export interface RecommendedStoreProps {
  name: string
  review_count: number
  average_rating: number
  image_src: string
  seller_id: number
}

export function RecommendedStore({
  name,
  review_count,
  average_rating,
  image_src,
  seller_id,
}: RecommendedStoreProps) {
  return (
    <Link href={`/seller/${seller_id}`}>
      <div className='flex flex-col items-center justify-center gap-2.5 rounded-lg bg-grey-50 pb-5 pt-2.5'>
        <div className='relative flex aspect-square size-full max-h-20 max-w-20 items-center justify-center overflow-hidden rounded-full'>
          <Image src={image_src} alt='' fill />
        </div>
        <div className='flex w-full flex-col items-center justify-center text-center'>
          <Text variant='sm-semibold' desktopVariant='md-semibold'>
            {name}
          </Text>
          <div className='flex items-center justify-center'>
            <Icon
              icon='material-symbols:star-rounded'
              className='size-4 text-warning'
            />
            <Text variant='sm-regular' className='text-grey-700'>
              {average_rating} | {review_count} รีวิว
            </Text>
          </div>
        </div>
      </div>
    </Link>
  )
}
