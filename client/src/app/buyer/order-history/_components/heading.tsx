import { Icon } from '@iconify/react'

import { Button } from '@/components/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/dropdown-menu'
import { Text } from '@/components/text'

import { type Order } from '@/generated/api'

const STATUS_MAP_TO_LABEL: Record<Order['status'] | 'ALL', string> = {
  ALL: 'ทั้งหมด',
  PREPARING: 'รอจัดส่ง',
  DELIVERING: 'กำลังจัดส่ง',
  COMPLETED: 'จัดส่งสำเร็จ',
}

interface HeadingProps {
  filter: Order['status'] | 'ALL'
  setFilter: (filter: Order['status'] | 'ALL') => void
}

export default function Heading({ filter, setFilter }: HeadingProps) {
  return (
    <div className='flex justify-between py-3'>
      <div className='flex items-center'>
        <Icon icon='material-symbols:list' className='size-9' />
        <Text variant='heading-md' desktopVariant='heading-lg'>
          รายการคำสั่งซื้อ
        </Text>
      </div>

      <div className='flex items-center gap-3'>
        <Text desktopVariant='md-regular' className='hidden sm:block'>
          ตัวกรองสถานะ
        </Text>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button className='min-w-12'>
              <Text desktopVariant='md-semibold' className='hidden sm:block'>
                {STATUS_MAP_TO_LABEL[filter]}
              </Text>
              <Icon icon='fa-solid:filter' className='size-3 sm:size-4' />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className='divide-y divide-grey-200 shadow-md'>
            {Object.entries(STATUS_MAP_TO_LABEL).map(([value, label]) => (
              <DropdownMenuItem
                key={value}
                onClick={() => setFilter(value as Order['status'] | 'ALL')}
              >
                {label}
              </DropdownMenuItem>
            ))}
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </div>
  )
}
