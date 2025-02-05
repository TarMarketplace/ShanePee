'use client'

import Image from 'next/image'
import { ChangeEvent, useState } from 'react'

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
    <div className='flex justify-center gap-8 rounded-xl bg-white px-24 py-12'>
      <Image
        src='https://placehold.co/360x120.png'
        alt='logo'
        width={360}
        height={120}
      />
      <div className='flex w-96 flex-col gap-6 rounded-xl border bg-grey-50 p-8 shadow-sm'>
        <h2 className='text-center text-2xl font-bold'>เข้าสู่ระบบ</h2>
        <LoginCredential
          userData={userData}
          handleInputChange={handleInputChange}
          nextSection={nextSection}
        />
        <Text className='text-center text-sm'>
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
