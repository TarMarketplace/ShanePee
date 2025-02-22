'use client'

import { Icon } from '@iconify/react/dist/iconify.js'
import { usePathname, useRouter, useSearchParams } from 'next/navigation'

import { Skeleton } from '@/components/skeleton'
import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import { AddressContainer } from './address-container'
import { PasswordContainer } from './password-container'
import { PaymentContainer } from './payment-container'
import { UserInfoContainer } from './user-info-container'

export function AccountContainer() {
  const searchParams = useSearchParams()
  const pathname = usePathname()
  const router = useRouter()
  const { user } = useUser()

  const mode = searchParams.get('mode') || 'info'

  const handleSwitchMode = (newMode: string) => {
    const params = new URLSearchParams(searchParams.toString())
    params.set('mode', newMode)
    router.push(`${pathname}?${params.toString()}`)
  }

  const renderMode = () => {
    if (!user) return <Skeleton className='size-full rounded-r-lg' />

    switch (mode) {
      case 'info':
        return <UserInfoContainer user={user} />
      case 'address':
        return <AddressContainer user={user} />
      case 'payment':
        return <PaymentContainer user={user} />
      case 'password':
        return <PasswordContainer />
    }
  }

  const modeList = [
    {
      label: 'บัญชีของฉัน',
      value: 'info',
      icon: 'mdi:account',
    },
    {
      label: 'ที่อยู่สำหรับการจัดส่งสินค้า',
      value: 'address',
      icon: 'ic:round-home',
    },
    { label: 'ช่องทางการชำระเงิน', value: 'payment', icon: 'famicons:card' },
    { label: 'เปลี่ยนรหัสผ่าน', value: 'password', icon: 'solar:key-bold' },
  ]

  return (
    <div className='flex w-full flex-col rounded-xl md:flex-row md:border'>
      <div className='flex w-full overflow-x-auto whitespace-nowrap pb-2 md:w-auto md:min-w-[300px] md:flex-col md:gap-2 md:border-r md:p-6'>
        {modeList.map(({ label, value, icon }) => (
          <button
            key={value}
            className={`inline-flex h-10 items-center justify-start gap-2 px-3 py-2 md:rounded-sm ${
              mode === value
                ? 'border-b-2 border-b-primary text-primary md:border-none md:bg-primary md:text-white'
                : 'text-black hover:bg-grey-50'
            }`}
            onClick={() => handleSwitchMode(value)}
          >
            {icon && <Icon icon={icon} className='size-5' />}
            <Text variant={mode === value ? 'md-semibold' : 'md-regular'}>
              {label}
            </Text>
          </button>
        ))}
      </div>
      {renderMode()}
    </div>
  )
}
