'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { type SubmitHandler, useForm } from 'react-hook-form'
import { toast } from 'sonner'
import { z } from 'zod'

import { Text } from '@/components/text'

import { passwordResetRequests } from '@/generated/api'

import { ForgotPasswordForm } from '../_components/fogot-password-form'

interface ForgotPasswordContainerProps {
  onSwitchMode: () => void
}

const forgotPasswordFormSchema = z.object({
  email: z.string().email('Invalid email'),
})

export type ForgotPasswordFormSchema = z.infer<typeof forgotPasswordFormSchema>

export function ForgotPasswordContainer({
  onSwitchMode,
}: ForgotPasswordContainerProps) {
  const form = useForm<ForgotPasswordFormSchema>({
    resolver: zodResolver(forgotPasswordFormSchema),
    defaultValues: {
      email: '',
    },
  })

  const onSubmit: SubmitHandler<ForgotPasswordFormSchema> = async (data) => {
    const { response, error } = await passwordResetRequests({
      body: {
        email: data.email,
      },
    })

    if (!response.ok) {
      toast.error(error?.detail)
      return
    }

    toast.success('Email sent successfully')
    form.reset()
  }

  return (
    <article className='flex w-full flex-col gap-8'>
      <Text variant='heading-md'>ลืมรหัสผ่าน</Text>

      <ForgotPasswordForm
        onSubmit={onSubmit}
        form={form}
        onSwitchMode={onSwitchMode}
      />
    </article>
  )
}
