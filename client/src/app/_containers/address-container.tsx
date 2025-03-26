'use client'

import { Icon } from '@iconify/react'
import Link from 'next/link'

import { Skeleton } from '@/components/skeleton'
import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import { Address } from '../../components/address'

interface AddressContainerProps {
  enablePencil?: boolean
}

export function AddressContainer({ enablePencil }: AddressContainerProps) {
  const { user } = useUser()
  const pencilOn = enablePencil ?? false

  return (
    <div className='flex flex-col gap-2'>
      <div className='flex items-center gap-2.5'>
        <Text desktopVariant='heading-sm' variant='md-semibold'>
          ที่อยู่สำหรับการจัดส่ง
        </Text>
        {pencilOn ? (
          <Link href='user/account?mode=address'>
            <Icon icon='mdi:pencil' className='size-5 md:size-6' />
          </Link>
        ) : null}
      </div>
      {!user ? (
        <div className='flex flex-col gap-2'>
          <Skeleton className='h-5 w-2/5' />
          <Skeleton className='h-5 w-3/5' />
          <Skeleton className='h-5 w-1/5' />
        </div>
      ) : (
        <Address user={user} />
      )}
    </div>
  )
}
