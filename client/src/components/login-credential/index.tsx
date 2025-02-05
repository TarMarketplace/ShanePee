'use client'

import { ChangeEvent, FormEvent, useState } from 'react'

import { ErrorLogin, UserLoginDetail } from '@/types/user-login'

import { ButtonCapsule } from '../button-capsule'
import { InputField } from '../input-field'

interface LoginCredentialProps {
  userData: UserLoginDetail
  handleInputChange: (e: ChangeEvent<HTMLInputElement>) => void
  nextSection: () => void
}

function LoginCredential({
  userData,
  handleInputChange,
  nextSection,
}: LoginCredentialProps) {
  const [errors, setErrors] = useState({
    username: '',
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
    const errs: ErrorLogin = { ...errors }

    if (!userData.username) {
      errs.username = 'username/email is required'
      formIsValid = false
    } else {
      errs.username = ''
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
        name='username'
        value={userData.username}
        onChange={handleInputChange}
        error={errors.username}
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

export default LoginCredential
