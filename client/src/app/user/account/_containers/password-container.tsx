import { zodResolver } from '@hookform/resolvers/zod'
import { Icon } from '@iconify/react/dist/iconify.js'
import { type SubmitHandler, useForm } from 'react-hook-form'
import { toast } from 'sonner'
import { z } from 'zod'

import { Text } from '@/components/text'

import { changePassword } from '@/generated/api'

import { PasswordForm } from '../_components/password-form'

const passwordFormSchema = z
  .object({
    oldPassword: z.string().min(1, 'Password is required'),
    password: z.string().min(8, 'New password must be at least 8 characters'),
    confirmPassword: z
      .string()
      .min(8, 'New password must be at least 8 characters'),
  })
  .refine((data) => data.password === data.confirmPassword, {
    path: ['confirmPassword'],
    message: 'Password does not match',
  })

export type PasswordFormSchema = z.infer<typeof passwordFormSchema>

export function PasswordContainer() {
  const form = useForm<PasswordFormSchema>({
    resolver: zodResolver(passwordFormSchema),
    defaultValues: {
      oldPassword: '',
      password: '',
      confirmPassword: '',
    },
  })

  const onSubmit: SubmitHandler<PasswordFormSchema> = async (data) => {
    const { response, error } = await changePassword({
      body: {
        new_password: data.password,
        old_password: data.oldPassword,
      },
    })

    if (response.ok) {
      toast.success('Updated successfully')
      form.reset()
    } else {
      toast.error('Something went wrong')
      toast.error(error?.detail)
    }
  }

  return (
    <div className='my-6 flex flex-col gap-8 md:m-6 md:w-full'>
      <div className='hidden items-center gap-2 md:flex'>
        <Icon icon='solar:key-bold' className='size-10' />
        <Text variant='heading-lg'>เปลี่ยนรหัสผ่าน</Text>
      </div>
      <PasswordForm onSubmit={onSubmit} form={form} />
    </div>
  )
}
