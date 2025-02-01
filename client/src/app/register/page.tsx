'use client'

import { formatPhoneNumber } from '@/util/FormatPhoneNumber'
import { Option, Select } from '@material-tailwind/react'
import Image from 'next/image'
import { ChangeEvent, FormEvent, useState } from 'react'

import { ButtonCapsule } from '@/components/button-capsule'
import { InputField } from '@/components/input-field'
import { Text } from '@/components/text'

type UserData = {
  name: string
  lastname: string
  email: string
  phone: string
  gender: string
  username: string
  password: string
  passwordConfirm: string
}

type Errors = {
  [key in keyof UserData]: string
}

export default function Register() {
  const [section, setSection] = useState(0)
  const [userData, setUserData] = useState<UserData>({
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
            <h2 className='text-center text-2xl font-bold'>สมัครใช้งาน</h2>
            {section == 0 ? (
              <Register1
                userData={userData}
                handleInputChange={handleInputChange}
                handleSelectChange={handleSelectChange}
                nextSection={nextSection}
              />
            ) : (
              <Register2
                userData={userData}
                handleInputChange={handleInputChange}
                handleSelectChange={handleSelectChange}
                nextSection={nextSection}
              />
            )}
            <Text className='text-center text-sm'>
              เคยสมัครใช้งานแล้วหรือไม่?{' '}
              <a href='/login' className='text-primary-500 hover:underline'>
                เข้าสู่ระบบ
              </a>
            </Text>
          </div>
        </div>
      </section>
    </main>
  )
}

function Register1({
  userData,
  handleInputChange,
  handleSelectChange,
  nextSection,
}) {
  const [errors, setErrors] = useState({
    name: '',
    lastname: '',
    email: '',
    phone: '',
    gender: '',
  })

  const data = ['name', 'lastname', 'email', 'phone', 'gender']

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

    data.forEach((field) => {
      if (!userData[field]) {
        errs[field] = `${field} is required`
        formIsValid = false
      } else {
        errs[field] = ''
      }
    })

    const phoneRegex = /^(\(\d{3}\)\s?|\d{3}-)\d{3}-\d{4}$/
    if (!phoneRegex.test(userData.phone)) {
      errs.phone = 'Invalid phone number format (e.g., 123-456-7890)'
      formIsValid = false
    }

    setErrors(errs)
    return formIsValid
  }

  return (
    <form onSubmit={handleSubmit} className='flex flex-col gap-4'>
      <InputField
        size='lg'
        label='ชื่อจริง'
        name='name'
        value={userData.name}
        onChange={handleInputChange}
        error={errors.name}
      />
      <InputField
        size='lg'
        label='นามสกุล'
        name='lastname'
        value={userData.lastname}
        onChange={handleInputChange}
        error={errors.lastname}
      />
      <InputField
        size='lg'
        label='อีเมล'
        name='email'
        value={userData.email}
        onChange={handleInputChange}
        error={errors.email}
      />
      <InputField
        size='lg'
        label='เบอร์โทรศัพท์'
        name='phone'
        value={userData.phone}
        onChange={handleInputChange}
        error={errors.phone}
      />
      <div className='h-14'>
        <div className='bg-white'>
          <Select
            size='lg'
            label='เพศ'
            name='gender'
            value={userData.gender}
            onChange={handleSelectChange}
          >
            <Option value='male'>ชาย</Option>
            <Option value='female'>หญิง</Option>
            <Option value='other'>อื่นๆ</Option>
          </Select>
        </div>
        {errors.gender && (
          <p className='text-sm text-red-500'>{errors.gender}</p>
        )}
      </div>

      <ButtonCapsule type='submit' className='mt-4 bg-primary-500 text-white'>
        ถัดไป
      </ButtonCapsule>
    </form>
  )
}

function Register2({ userData, handleInputChange, nextSection }) {
  const [errors, setErrors] = useState({
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
    const errs: Errors = { ...errors }

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
        ถัดไป
      </ButtonCapsule>
    </form>
  )
}
