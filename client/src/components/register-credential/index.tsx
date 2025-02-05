'use client'

import { ChangeEvent, FormEvent, useState } from 'react'

import { ErrorRegister, UserRegisterDetail } from '@/types/user-register'

import { ButtonCapsule } from '../button-capsule'
import { InputField } from '../input-field'

interface RegisterCredentialProps {
  userData: UserRegisterDetail
  handleInputChange: (e: ChangeEvent<HTMLInputElement>) => void
  handleSelectChange: (value: string) => void
  nextSection: () => void
}

function RegisterCredential({
  userData,
  handleInputChange,
  nextSection,
}: RegisterCredentialProps) {
  const [errors, setErrors] = useState({
    name: '',
    lastname: '',
    email: '',
    phone: '',
    gender: '',
    username: '',
    password: '',
    passwordConfirm: '',
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
    const errs: ErrorRegister = { ...errors }

    if (!userData.username) {
      errs.username = 'username is required'
      formIsValid = false
    } else {
      // TODO implement check if username avaliable
      errs.username = ''
    }

    if (!userData.password) {
      errs.password = 'password is required'
      formIsValid = false
    } else {
      if (userData.password.length < 8) {
        errs.password = 'password need to be at least 8 characters'
        formIsValid = false
      } else {
        if (userData.password != userData.passwordConfirm) {
          errs.password = 'password does not match'
          errs.passwordConfirm = 'password does not match'
          formIsValid = false
        } else {
          errs.password = ''
          errs.passwordConfirm = ''
        }
      }
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
      <InputField
        type='password'
        size='lg'
        label='ยืนยันรหัสผ่าน'
        name='passwordConfirm'
        value={userData.passwordConfirm}
        onChange={handleInputChange}
        error={errors.passwordConfirm}
      />

      <ButtonCapsule type='submit' className='mt-4 bg-primary-500 text-white'>
        สมัครใช้งาน
      </ButtonCapsule>
    </form>
  )
}

export default RegisterCredential
