'use client'

import { formatPhoneNumber } from '@/utils/FormatPhoneNumber'
import { ChangeEvent, useState } from 'react'

import AuthPageBanner from '@/components/auth-page-banner'
import RegisterCredential from '@/components/register-credential'
import RegisterDetail from '@/components/register-detail'
import { Text } from '@/components/text'

import { UserRegisterDetail } from '@/types/user-register'

const RegisterPageContainer = () => {
  const [section, setSection] = useState(0)
  const [userData, setUserData] = useState<UserRegisterDetail>({
    name: '',
    lastname: '',
    email: '',
    phone: '',
    gender: '',
    username: '',
    password: '',
    passwordConfirm: '',
  })

  const nextSection = () => {
    if (section == 1) {
      // TODO Handle user register
      console.log('Handle user register')
      console.log(userData)
    } else {
      setSection(1)
    }
  }

  const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    const formattedValue = name === 'phone' ? formatPhoneNumber(value) : value

    setUserData({
      ...userData,
      [name]: formattedValue,
    })
  }

  const handleSelectChange = (value: string | undefined) => {
    setUserData((prevState) => ({
      ...prevState,
      gender: value || '',
    }))
  }

  return (
    <div className='flex justify-center gap-24 rounded-xl bg-white px-24 py-12'>
      <AuthPageBanner />
      <div className='flex w-96 flex-col gap-6 rounded-xl border p-6 shadow-sm'>
        <h2 className='text-2xl font-bold'>สมัครใช้งาน</h2>
        {section == 0 ? (
          <RegisterDetail
            userData={userData}
            handleInputChange={handleInputChange}
            handleSelectChange={handleSelectChange}
            nextSection={nextSection}
          />
        ) : (
          <RegisterCredential
            userData={userData}
            handleInputChange={handleInputChange}
            handleSelectChange={handleSelectChange}
            nextSection={nextSection}
          />
        )}
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

export { RegisterPageContainer }
