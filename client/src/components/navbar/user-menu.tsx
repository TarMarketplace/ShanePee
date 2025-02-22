import { Icon } from '@iconify/react'
import Link from 'next/link'
import { useState } from 'react'

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/dropdown-menu'
import { Text } from '@/components/text'

import type { User } from '@/generated/api'

import { Sidebar } from '../side-menu'

interface UserMenuProps {
  user: User | null
  onLogout: () => void
}

export const UserMenu = ({ user, onLogout }: UserMenuProps) => {
  const [isSideMenuOpen, setIsSideMenuOpen] = useState(false)

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
          <button className='hidden md:block'>
            {user && (
              <div className='hidden items-center gap-1 text-nowrap md:flex'>
                <Text variant='md-semibold'>สวัสดี, {user.first_name} </Text>
                <Icon icon='teenyicons:down-solid' className='size-3' />
              </div>
            )}
          </button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className='divide-y divide-grey-200'>
          <DropdownMenuItem>
            <Link href='/user/account'>บัญชีของฉัน</Link>
          </DropdownMenuItem>
          <DropdownMenuItem>การซื้อของฉัน</DropdownMenuItem>
          <DropdownMenuItem>ยืนยันตัวตนผู้ขาย</DropdownMenuItem>
          <DropdownMenuItem>
            <button onClick={onLogout} className='text-error'>
              ออกจากระบบ
            </button>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
      <div className='md:hidden'>
        <Sidebar
          isOpen={isSideMenuOpen}
          setIsOpen={setIsSideMenuOpen}
          onLogout={onLogout}
        />
      </div>
    </div>
  )
}
