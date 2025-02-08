'use client'

import { Icon } from '@iconify/react/dist/iconify.js'
import { usePathname, useRouter, useSearchParams } from 'next/navigation'
import { useMemo } from 'react'

import { Text } from '@/components/text'

import { AddressContainer } from './address-container'
import { PasswordContainer } from './password-container'
import { PaymentContainer } from './payment-container'
import { UserInfoContainer } from './user-info-container'

export function AccountContainer() {
  const searchParams = useSearchParams()
  const pathname = usePathname()
  const router = useRouter()

  const mode = searchParams.get('mode') || 'info'

  const handleSwitchMode = (newMode: string) => {
    const params = new URLSearchParams(searchParams.toString())
    params.set('mode', newMode)
    router.push(`${pathname}?${params.toString()}`)
  }

  const renderMode = useMemo(() => {
    switch (mode) {
      case 'info':
        return <UserInfoContainer />
      case 'address':
        return <AddressContainer />
      case 'payment':
        return <PaymentContainer />
      case 'password':
        return <PasswordContainer />
    }
  }, [mode])

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
    <div className='flex w-full rounded-xl border'>
      <div className='flex w-auto min-w-[300px] flex-col gap-2 whitespace-nowrap border-r p-6'>
        {modeList.map(({ label, value, icon }) => (
          <button
            key={value}
            className={`inline-flex h-10 items-center justify-start gap-2 rounded-sm px-3 py-2 ${
              mode === value
                ? 'bg-primary text-white'
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
      {renderMode}
    </div>
  )
}
