import { Icon } from '@iconify/react'
import Image from 'next/image'
import Link from 'next/link'

import { Text } from '@/components/text'

import type { User } from '@/types/user'

import { UserMenu } from './user-menu'

interface NavbarProps {
  user: User | null
  onLogin: () => void
  onLogout: () => void
  searchValue: string
  onChangeSearchValue: (value: string) => void
  onSearch: () => void
}

const Navbar = ({
  user,
  onLogin,
  onLogout,
  searchValue,
  onChangeSearchValue,
  onSearch,
}: NavbarProps) => {
  return (
    <nav className='sticky z-50 flex w-full items-center justify-between bg-primary-gradient px-3 py-2 text-white'>
      <div className='flex w-full max-w-2xl items-center gap-[1.125rem]'>
        <Link href='/'>
          <Image
            src='https://placehold.co/150x40.png'
            alt=''
            width={150}
            height={40}
          />
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
          <button className='px-3' onClick={onLogin}>
            <Text variant='md-semibold'>เข้าสู่ระบบ</Text>
          </button>
          <button className='px-3'>
            <Text variant='md-semibold'>สมัครใช้งาน</Text>
          </button>
        </div>
      )}
    </nav>
  )
}

export { Navbar }
