'use client'

import { ChangeEvent, useState } from 'react'

import AuthPageBanner from '@/components/auth-page-banner'
import LoginCredential from '@/components/login-credential'
import { Text } from '@/components/text'

import { UserLoginDetail } from '@/types/user-login'

const LoginPageContainer = () => {
  const [userData, setUserData] = useState<UserLoginDetail>({
    username: '',
    password: '',
  })

  const nextSection = () => {
    console.log('Handle user login')
    console.log(userData)
  }

  const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target

    setUserData({
      ...userData,
      [name]: value,
    })
  }

  return (
    <div className='flex justify-center gap-24 rounded-xl bg-white px-24 py-12'>
      <AuthPageBanner />
      <div className='flex w-96 flex-col gap-6 rounded-xl border p-6 shadow-sm'>
        <h2 className='text-2xl font-bold'>เข้าสู่ระบบ</h2>
        <LoginCredential
          userData={userData}
          handleInputChange={handleInputChange}
          nextSection={nextSection}
        />
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

export { LoginPageContainer }
