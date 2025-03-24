import Image from 'next/image'

import { Button } from '@/components/button'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/popover'
import { Text } from '@/components/text'

export default function AddReviewContainerl() {
  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button variant='filled'>
          <Image src='/review-button-star.svg' width={24} height={24} alt='' />
          <Text desktopVariant='md-semibold'>ให้คะแนน</Text>
        </Button>
      </PopoverTrigger>
      <PopoverContent className='size-fit'>
        {/* TODO: Review modal  */}
      </PopoverContent>
    </Popover>
  )
}
