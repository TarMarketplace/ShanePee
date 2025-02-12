'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'
import { type SubmitHandler, useForm } from 'react-hook-form'
import { toast } from 'sonner'
import { z } from 'zod'

import { Text } from '@/components/text'

import { env } from '@/env'

import { RegisterStep1Form } from '../_components/register-step-1-form'
import { RegisterStep2Form } from '../_components/register-step-2-form'

interface RegisterContainerProps {
  onSwitchMode: () => void
}

const registerStep1FormSchema = z.object({
  name: z.string().min(1, 'Name is required'),
  surname: z.string().min(1, 'Surname is required'),
  phone: z
    .string()
    .min(10, 'Phone number is required')
    .refine((value) => {
      return /^\d+$/.test(value)
    }),
  gender: z.enum(['MALE', 'FEMALE', 'OTHER']),
})

const registerStep2FormSchema = z
  .object({
    email: z.string().email('Invalid email'),
    password: z.string().min(8, 'Password must be at least 8 characters'),
    confirmPassword: z
      .string()
      .min(8, 'Password must be at least 8 characters'),
  })
  .refine((data) => data.password === data.confirmPassword, {
    path: ['confirmPassword'],
    message: 'Password does not match',
  })

export type RegisterStep1FormSchema = z.infer<typeof registerStep1FormSchema>
export type RegisterStep2FormSchema = z.infer<typeof registerStep2FormSchema>

export function RegisterContainer({ onSwitchMode }: RegisterContainerProps) {
  const [step, setStep] = useState(1)

  const step1Form = useForm<RegisterStep1FormSchema>({
    resolver: zodResolver(registerStep1FormSchema),
    defaultValues: {
      name: '',
      surname: '',
      phone: '',
      gender: undefined,
    },
  })
  const step2Form = useForm<RegisterStep2FormSchema>({
    resolver: zodResolver(registerStep2FormSchema),
    defaultValues: {
      email: '',
      password: '',
      confirmPassword: '',
    },
  })

  const onSubmitStep1: SubmitHandler<RegisterStep1FormSchema> = () => {
    setStep(2)
  }

  const handleUpdateUserDetails = async (data: RegisterStep1FormSchema) => {
    const res = await fetch(`${env.NEXT_PUBLIC_BASE_API_URL}/user`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify({
        first_name: data.name,
        last_name: data.surname,
        gender: data.gender,
        tel: data.phone,
      }),
    })

    if (!res.ok) {
      toast.error('Cannot update user details, please try again')
      return
    }
  }

  const onSubmitStep2: SubmitHandler<RegisterStep2FormSchema> = async (
    data
  ) => {
    const res = await fetch(`${env.NEXT_PUBLIC_BASE_API_URL}/auth/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify({
        email: data.email,
        password: data.password,
      }),
    })

    if (res.ok) {
      toast.success('Registered successfully')
      await handleUpdateUserDetails(step1Form.getValues())
    } else {
      toast.error('Something went wrong')
    }
  }

  const renderFormByStep = () => {
    switch (step) {
      case 1:
        return <RegisterStep1Form onSubmit={onSubmitStep1} form={step1Form} />
      case 2:
        return <RegisterStep2Form onSubmit={onSubmitStep2} form={step2Form} />
    }
  }

  return (
    <article className='flex w-full flex-col gap-8'>
      <Text variant='heading-md'>สมัครใช้งาน</Text>

      {renderFormByStep()}
      <button onClick={onSwitchMode}>
        <Text variant='sm-regular' className='text-muted-foreground'>
          เคยสมัครใช้งานแล้วหรือไม่?{' '}
          <span className='text-primary-500 hover:underline'>เข้าสู่ระบบ</span>
        </Text>
      </button>
    </article>
  )
}
