import { Icon } from '@iconify/react/dist/iconify.js'
import { AvatarImage } from '@radix-ui/react-avatar'

import { Avatar, AvatarFallback } from '@/components/avatar'
import { Button } from '@/components/button'
import { Text } from '@/components/text'

export interface SellerStatsCardProps {
  photo: string
  sellerName: string
  store_rating: number
  total_art_toys: number
  sold_art_toys: number
  total_review: number
  joined_for: number
  isSeller: boolean
}

function SellerStatsCard({
  photo,
  sellerName,
  store_rating,
  total_art_toys,
  sold_art_toys,
  total_review,
  joined_for,
  isSeller,
}: SellerStatsCardProps) {
  return (
    <div className='grid grid-cols-1 items-center justify-center gap-1 md:grid-cols-3 md:gap-6'>
      <div className='flex size-full items-center gap-2 rounded-md border border-grey-50 bg-grey-50 p-4 shadow'>
        <Avatar className='size-20'>
          <AvatarImage src={photo} alt={sellerName} />
          <AvatarFallback>JD</AvatarFallback>
        </Avatar>
        <Text variant='lg-semibold'>{sellerName}</Text>
      </div>
      <div className='flex h-full flex-col justify-between gap-1 pt-4 md:p-1 md:pt-0'>
        <span className='flex items-center gap-2'>
          <Icon icon='iconoir:star-solid' className='size-5' />
          <Text variant='md-regular'>คะแนนร้านค้า: </Text>
          <Text variant='md-regular' className='text-primary'>
            {store_rating}
          </Text>
        </span>
        <span className='flex items-center gap-2'>
          <Icon icon='material-symbols-light:box-rounded' className='size-5' />
          <Text variant='md-regular'>จำนวนสินค้าทั้งหมด: </Text>
          <Text variant='md-regular' className='text-primary'>
            {total_art_toys} ชิ้น
          </Text>
        </span>
        <span className='flex items-center gap-2'>
          <Icon icon='mdi:cart' className='size-5' />
          <Text variant='md-regular'>จำนวนสินค้าที่ขายได้: </Text>
          <Text variant='md-regular' className='text-primary'>
            {sold_art_toys} ชิ้น
          </Text>
        </span>
      </div>
      <div className='flex h-full flex-col justify-between gap-1 md:p-1'>
        <span className='flex items-center gap-2'>
          <Icon icon='material-symbols:chat' className='size-5' />
          <Text variant='md-regular'>จำนวนรีวิว: </Text>
          <Text variant='md-regular' className='text-primary'>
            {total_review}
          </Text>
        </span>
        <span className='flex items-center gap-2'>
          <Icon icon='fa-solid:sign' className='size-5' />
          <Text variant='md-regular'>เข้าร่วมเมื่อ: </Text>
          <Text variant='md-regular' className='text-primary'>
            {joined_for} วันที่ผ่านมา
          </Text>
        </span>
        <span className='mt-4 flex items-center gap-2 md:mt-0'>
          <Button
            variant='filled'
            className={isSeller ? 'invisible h-8 md:h-6' : 'h-8 md:h-6'}
          >
            <Icon icon='material-symbols:chat' className='size-5' />
            <Text variant='md-regular'>แชทเลย</Text>
          </Button>
        </span>
      </div>
    </div>
  )
}

export { SellerStatsCard }
