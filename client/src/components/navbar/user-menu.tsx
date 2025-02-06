import { Icon } from '@iconify/react'

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/dropdown-menu'
import { Text } from '@/components/text'

import type { User } from '@/types/users'

interface UserMenuProps {
  user: User
  onLogout: () => void
}

export const UserMenu = ({ user, onLogout }: UserMenuProps) => {
  return (
    <div className='flex items-center gap-[1.125rem]'>
      <Icon icon='tdesign:cart-filled' className='size-7' />
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <button className='flex items-center gap-1 text-nowrap'>
            <Text variant='md-semibold'>สวัสดี, {user.name}</Text>
            <Icon icon='teenyicons:down-solid' className='size-3' />
          </button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className='divide-y divide-grey-200'>
          <DropdownMenuItem>บัญชีของฉัน</DropdownMenuItem>
          <DropdownMenuItem>การซื้อของฉัน</DropdownMenuItem>
          <DropdownMenuItem>ยืนยันตัวตนผู้ขาย</DropdownMenuItem>
          <DropdownMenuItem>
            <button onClick={onLogout}>ออกจากระบบ</button>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}
