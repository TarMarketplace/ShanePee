import { zodResolver } from '@hookform/resolvers/zod'
import { useRouter } from 'next/navigation'
import { type SubmitHandler, useForm } from 'react-hook-form'
import { toast } from 'sonner'
import { z } from 'zod'

import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import { LoginForm } from '../_components/login-form'

const loginFormSchema = z.object({
  email: z.string().email('Invalid email'),
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
  const router = useRouter()
  const { setUser } = useUser()

  const form = useForm<LoginFormSchema>({
    resolver: zodResolver(loginFormSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  })

  const onSubmit: SubmitHandler<LoginFormSchema> = async (data) => {
    const res = await fetch('/api/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })

    if (res.ok) {
      toast.success('Successfully logged in')
      const resData = await res.json()

      setUser({ ...resData })
      router.push('/')
    } else {
      toast.error('Something went wrong')
    }
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
