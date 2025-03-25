import { AvatarImage } from '@radix-ui/react-avatar'

import { Avatar, AvatarFallback } from '@/components/avatar'
import { Text } from '@/components/text'

export interface ChatBoxProps {
  photo: string
  sellerName: string
  date: Date
  message: string
  selected: boolean
}

function ChatBox({ photo, sellerName, date, message, selected }: ChatBoxProps) {
  const dateTime = new Date(date)

  return (
    <div
      className={
        selected
          ? 'flex h-20 w-full items-center gap-2 truncate bg-primary-50 p-2'
          : 'flex h-20 w-full items-center gap-2 truncate p-2'
      }
    >
      <Avatar className='size-14'>
        <AvatarImage src={photo} alt={sellerName} />
        <AvatarFallback>JD</AvatarFallback>
      </Avatar>
      <div className='flex w-full flex-col justify-center gap-1 truncate'>
        <span className='flex justify-between'>
          <Text variant='md-semibold'>{sellerName}</Text>
          <Text variant='sm-regular' className='text-grey-500'>
            {String(dateTime.getHours()).padStart(2, '0')}:
            {String(dateTime.getMinutes()).padStart(2, '0')}
          </Text>
        </span>
        <Text
          variant='sm-regular'
          className='overflow-auto truncate text-grey-500'
        >
          {message}
        </Text>
      </div>
    </div>
  )
}

export { ChatBox }
