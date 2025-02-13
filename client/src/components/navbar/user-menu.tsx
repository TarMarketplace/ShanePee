import { Icon } from '@iconify/react'
import Link from 'next/link'

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/dropdown-menu'
import { Text } from '@/components/text'

import type { User } from '@/types/users'

interface UserMenuProps {
  user: User | null
  onLogout: () => void
}

export const UserMenu = ({ user, onLogout }: UserMenuProps) => {
  return (
    <div className='ml-4 flex items-center gap-[1.125rem]'>
      <Link
        href={user ? '/cart' : '/login'}
        className={user ? '' : 'md:hidden'}
      >
        <Icon icon='tdesign:cart-filled' className='size-7' />
      </Link>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <button>
            {user && (
              <div className='hidden items-center gap-1 text-nowrap md:flex'>
                <Text variant='md-semibold'>สวัสดี, {user.first_name} </Text>
                <Icon icon='teenyicons:down-solid' className='size-3' />
              </div>
            )}
            <div className='block md:hidden'>
              <Icon
                icon='akar-icons:three-line-horizontal'
                className='size-8'
              />
            </div>
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
