'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useSearchParams } from 'next/navigation'
import { type SubmitHandler, useForm } from 'react-hook-form'
import { toast } from 'sonner'
import { z } from 'zod'

import { Text } from '@/components/text'

import { resetPassword } from '@/generated/api'

import { ForgotPasswordForm } from '../_components/fogot-password-form'

const forgotPasswordFormSchema = z
  .object({
    password: z.string().min(8, 'Password must be at least 8 characters'),
    confirmPassword: z
      .string()
      .min(8, 'Password must be at least 8 characters'),
  })
  .refine((data) => data.password === data.confirmPassword, {
    path: ['confirmPassword'],
    message: 'Password does not match',
  })

export type ForgotPasswordFormSchema = z.infer<typeof forgotPasswordFormSchema>

export function ForgotPasswordContainer() {
  const searchParams = useSearchParams()
  const request_id = searchParams.get('request_id')
  const token = searchParams.get('token')

  const form = useForm<ForgotPasswordFormSchema>({
    resolver: zodResolver(forgotPasswordFormSchema),
    defaultValues: {
      password: '',
      confirmPassword: '',
    },
  })

  const onSubmit: SubmitHandler<ForgotPasswordFormSchema> = async (data) => {
    const { response, error } = await resetPassword({
      body: {
        new_password: data.password,
        request_id: Number(request_id),
        token: token!,
      },
    })

    if (!response.ok) {
      toast.error(error?.detail)
      return
    }

    toast.success('Email sent successfully')
    form.reset()
  }

  if (!request_id || !token) {
    return (
      <div className='flex w-full items-center justify-center rounded-xl border p-6 shadow-sm'>
        <Text>Invalid request</Text>
      </div>
    )
  }

  return (
    <div className='flex w-full rounded-xl border p-6 shadow-sm'>
      <article className='flex w-full flex-col gap-8'>
        <Text variant='heading-md'>ตั้งค่ารหัสผ่านใหม่</Text>
        <ForgotPasswordForm onSubmit={onSubmit} form={form} />
      </article>
    </div>
  )
}
