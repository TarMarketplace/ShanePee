'use client'

import { Icon } from '@iconify/react'
import Image from 'next/image'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { useState } from 'react'

import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import { UserMenu } from './user-menu'

const Navbar = () => {
  const router = useRouter()
  const { user, setUser } = useUser()
  const [search, setSearch] = useState('')

  const handleLogin = () => {
    setUser({ id: '1', name: 'lnwJoZaSodaSing+' })
  }

  const handleLogout = () => {
    setUser(null)
  }

  const handleSearch = () => {
    // TODO: Implement search
    console.log(search)
    router.push(`/search?query=${search}`)
  }

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
            value={search}
            onChange={(e) => setSearch(e.target.value)}
          />
          <button
            className='rounded-r-lg bg-secondary-500 p-1.5'
            onClick={handleSearch}
          >
            <Icon icon='mdi:magnify' className='size-8 text-black' />
          </button>
        </div>
      </div>
      {user ? (
        <div className='flex items-center gap-[1.125rem]'>
          <Icon icon='tdesign:cart-filled' className='size-7' />
          <UserMenu user={user} onLogout={handleLogout} />
        </div>
      ) : (
        <div className='flex items-center divide-x divide-white text-nowrap'>
          <button className='px-3' onClick={handleLogin}>
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
