'use client'

import Image from 'next/image'
import { ChangeEvent, FormEvent, useState } from 'react'

import { ButtonCapsule } from '@/components/button-capsule'
import { InputField } from '@/components/input-field'
import { Text } from '@/components/text'

type UserData = {
  user: string
  password: string
}

type Errors = {
  [key in keyof UserData]: string
}

export default function Login() {
  const [userData, setUserData] = useState<UserData>({
    user: '',
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
    <main className='flex w-full flex-col items-center gap-8 bg-background'>
      <section className='flex w-full items-center justify-center gap-10 bg-background-gradient p-12'>
        <div className='flex justify-center gap-8 rounded-xl bg-white px-24 py-12'>
          <Image
            src='https://placehold.co/360x120.png'
            alt='logo'
            width={360}
            height={120}
          />
          <div className='flex w-96 flex-col gap-6 rounded-xl border bg-grey-50 p-8 shadow-sm'>
            <h2 className='text-center text-2xl font-bold'>เข้าสู่ระบบ</h2>
            <Login1
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
      </section>
    </main>
  )
}

function Login1({ userData, handleInputChange, nextSection }) {
  const [errors, setErrors] = useState({
    user: '',
    password: '',
  })

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (validateForm()) {
      nextSection()
    } else {
      console.log('Form has errors:', errors)
    }
  }

  const validateForm = () => {
    let formIsValid = true
    const errs: Errors = { ...errors }

    if (!userData.user) {
      errs.user = 'username/email is required'
      formIsValid = false
    } else {
      errs.user = ''
    }

    if (!userData.password) {
      errs.password = 'password is required'
      formIsValid = false
    } else {
      errs.password = ''
    }

    setErrors(errs)
    return formIsValid
  }

  return (
    <form onSubmit={handleSubmit} className='flex flex-col gap-4'>
      <InputField
        size='lg'
        label='ชื่อผู้ใช้งาน'
        name='user'
        value={userData.user}
        onChange={handleInputChange}
        error={errors.user}
      />
      <InputField
        type='password'
        size='lg'
        label='รหัสผ่าน'
        name='password'
        value={userData.password}
        onChange={handleInputChange}
        error={errors.password}
      />

      <ButtonCapsule type='submit' className='mt-4 bg-primary-500 text-white'>
        เข้าสู่ระบบ
      </ButtonCapsule>
    </form>
  )
}
