import { Icon } from '@iconify/react'
import Image from 'next/image'
import Link from 'next/link'

import { Text } from '@/components/text'

import type { User } from '@/types/users'

import { UserMenu } from './user-menu'

interface NavbarProps {
  user: User | null
  onLogout: () => void
  searchValue: string
  onChangeSearchValue: (value: string) => void
  onSearch: () => void
}

const Navbar = ({
  user,
  onLogout,
  searchValue,
  onChangeSearchValue,
  onSearch,
}: NavbarProps) => {
  return (
    <header className='sticky top-0 z-50 flex w-full items-center justify-center bg-primary-gradient'>
      <nav className='flex w-full max-w-screen-xl items-center justify-between px-3 py-2 text-white'>
        <div className='flex w-full max-w-2xl items-center gap-[1.125rem]'>
          <Link href='/' className='hidden md:block'>
            <Image src='/logo.png' alt='' width={160} height={80} />
          </Link>
          <div className='flex w-full'>
            <input
              className='w-full rounded-l-lg px-3 py-2 text-black'
              placeholder='Art toys...'
              value={searchValue}
              onChange={(e) => onChangeSearchValue(e.target.value)}
            />
            <button
              className='rounded-r-lg bg-secondary-500 p-1.5'
              onClick={onSearch}
            >
              <Icon icon='mdi:magnify' className='size-8 text-black' />
            </button>
          </div>
        </div>
        {user ? (
          <UserMenu user={user} onLogout={onLogout} />
        ) : (
          <div className='flex items-center divide-x divide-white text-nowrap'>
            <Link href='/login' className='px-3'>
              <Text variant='md-semibold'>เข้าสู่ระบบ</Text>
            </Link>
            <Link href='/login?mode=register' className='px-3'>
              <Text variant='md-semibold'>สมัครใช้งาน</Text>
            </Link>
          </div>
        )}
      </nav>
    </header>
  )
}

export { Navbar }
