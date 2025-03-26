import { Text } from '@/components/text'

import type { User } from '@/generated/api'

interface AddressProps {
  user: User
}

export function Address({ user }: AddressProps) {
  return (
    <div className='flex flex-col'>
      <Text desktopVariant='md-regular' variant='sm-regular'>
        {user.first_name} {user.last_name}
      </Text>
      <Text desktopVariant='md-regular' variant='sm-regular'>
        {user.address.house_no} {user.address.district} {user.address.province}
        {user.address.postcode}
      </Text>
      <Text desktopVariant='md-regular' variant='sm-regular'>
        {user.tel}
      </Text>
    </div>
  )
}
