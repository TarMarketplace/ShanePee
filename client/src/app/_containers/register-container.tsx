'use client'

import { AuthBanner } from '@/components/auth-page-banner'
import { Text } from '@/components/text'

const RegisterContainer = () => {
  return (
    <div className='flex w-full justify-center gap-12 rounded-xl bg-white shadow-lg md:max-w-[1080px] md:p-12 lg:gap-24 lg:px-24'>
      <AuthBanner />
      <div className='flex w-full flex-col gap-6 rounded-xl border p-6 shadow-sm'>
        <h2 className='text-2xl font-bold'>สมัครใช้งาน</h2>
        <Text className='text-center text-xs text-grey-500'>
          เคยสมัครใช้งานแล้วหรือไม่?{' '}
          <a href='/login' className='text-primary-500 hover:underline'>
            เข้าสู่ระบบ
          </a>
        </Text>
      </div>
    </div>
  )
}

export { RegisterContainer }
