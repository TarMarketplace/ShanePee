import { Option, Select } from '@material-tailwind/react'
import { ChangeEvent, FormEvent, useState } from 'react'

import { ErrorRegister, UserRegisterDetail } from '@/types/user-register'

import { ButtonCapsule } from '../button-capsule'
import { InputField } from '../input-field'

interface RegisterDetailProps {
  userData: UserRegisterDetail
  handleInputChange: (e: ChangeEvent<HTMLInputElement>) => void
  handleSelectChange: (value: string) => void
  nextSection: () => void
}

function RegisterDetail({
  userData,
  handleInputChange,
  handleSelectChange,
  nextSection,
}: RegisterDetailProps) {
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

  const data: (keyof UserRegisterDetail)[] = [
    'name',
    'lastname',
    'email',
    'phone',
    'gender',
  ]

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
            onChange={(value) => handleSelectChange(value as string)}
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

      <ButtonCapsule type='submit' className='bg-primary-500 text-white'>
        ต่อไป
      </ButtonCapsule>
    </form>
  )
}

export default RegisterDetail
