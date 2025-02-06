'use client'

import Link from 'next/link'

import { AuthBanner } from '@/components/auth-page-banner'
import { Text } from '@/components/text'

const LoginContainer = () => {
  return (
    <div className='flex w-full justify-center gap-12 rounded-xl bg-white shadow-lg md:max-w-5xl md:p-12 lg:gap-24 lg:px-24'>
      <AuthBanner />
      <div className='flex w-full flex-col gap-6 rounded-xl border p-6 shadow-sm'>
        <h2 className='text-2xl font-bold'>เข้าสู่ระบบ</h2>
        <div className='flex flex-col gap-2'>
          <Link
            className='text-end text-xs text-primary-500 hover:underline'
            href='/reset-password'
          >
            ลืมรหัสผ่าน?
          </Link>
        </div>
        <Text className='text-center text-sm text-grey-500'>
          เพิ่งเคยใช้งานครั้งแรกหรือไม่?{' '}
          <a href='/register' className='text-primary-500 hover:underline'>
            สมัครใช้งาน
          </a>
        </Text>
      </div>
    </div>
  )
}

export { LoginContainer }
