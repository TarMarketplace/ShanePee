import { zodResolver } from '@hookform/resolvers/zod'
import { type SubmitHandler, useForm } from 'react-hook-form'
import { z } from 'zod'

import { Text } from '@/components/text'

import { LoginForm } from '../_components/login-form'

const loginFormSchema = z.object({
  username: z.string().min(1, 'Username is required'),
  password: z.string().min(1, 'Password is required'),
})

export type LoginFormSchema = z.infer<typeof loginFormSchema>

interface LoginContainerProps {
  onForgotPassword: () => void
  onSwitchMode: () => void
}

export function LoginContainer({
  onForgotPassword,
  onSwitchMode,
}: LoginContainerProps) {
  const form = useForm<LoginFormSchema>({
    resolver: zodResolver(loginFormSchema),
    defaultValues: {
      username: '',
      password: '',
    },
  })

  const onSubmit: SubmitHandler<LoginFormSchema> = (data) => {
    console.log(data)
  }

  return (
    <article className='flex w-full flex-col gap-8'>
      <Text variant='heading-md'>เข้าสู่ระบบ</Text>
      <div className='flex w-full flex-col gap-2.5'>
        <LoginForm onSubmit={onSubmit} form={form} />
        <button
          onClick={onForgotPassword}
          className='ml-auto w-fit px-3 text-right text-primary-500 hover:underline'
        >
          <Text variant='sm-regular'>ลืมรหัสผ่าน?</Text>
        </button>
      </div>
      <button onClick={onSwitchMode}>
        <Text variant='sm-regular' className='text-muted-foreground'>
          เพิ่งเคยใช้งานครั้งแรกหรือไม่?{' '}
          <span className='text-primary-500 hover:underline'>สมัครใช้งาน</span>
        </Text>
      </button>
    </article>
  )
}
