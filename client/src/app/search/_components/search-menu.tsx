'use client'

import { Icon } from '@iconify/react'
import Link from 'next/link'
import { usePathname, useSearchParams } from 'next/navigation'

import { Button } from '@/components/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/dropdown-menu'
import { Text } from '@/components/text'

const SORT_OPTIONS_MAP: Record<string, string> = {
  latest: 'ใหม่ล่าสุด',
  oldest: 'เก่าที่สุด',
  highest_price: 'ราคาสูงสุด',
  lowest_price: 'ราคาต่ำสุด',
}

export function SearchMenu() {
  const searchParams = useSearchParams()
  const pathname = usePathname()

  const currentSort = searchParams.get('sort') || 'latest'

  const getSortUrl = (value: string) => {
    const params = new URLSearchParams(searchParams.toString())
    params.set('sort', value)
    return `${pathname}?${params.toString()}`
  }

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button>
          <Text variant='sm-semibold' desktopVariant='md-semibold'>
            {SORT_OPTIONS_MAP[currentSort] || 'ใหม่ล่าสุด'}
          </Text>
          <Icon icon='icon-park-solid:down-one' className='size-3 sm:size-4' />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className='divide-y divide-grey-200 shadow-md'>
        {Object.entries(SORT_OPTIONS_MAP).map(([value, label]) => (
          <DropdownMenuItem key={value}>
            {value === currentSort ? (
              label
            ) : (
              <Link href={getSortUrl(value)}>{label}</Link>
            )}
          </DropdownMenuItem>
        ))}
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
