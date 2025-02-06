import { type SubmitHandler, type UseFormReturn } from 'react-hook-form'

import { Button } from '@/components/button'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/form'
import { Input } from '@/components/input'

import { type LoginFormSchema } from '../_containers/login-container'

interface LoginFormProps {
  onSubmit: SubmitHandler<LoginFormSchema>
  form: UseFormReturn<LoginFormSchema>
}

export function LoginForm({ onSubmit, form }: LoginFormProps) {
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className='flex w-full flex-col gap-8'
      >
        <div className='flex w-full flex-col gap-7'>
          <FormField
            control={form.control}
            name='username'
            render={({ field }) => (
              <FormItem>
                <FormLabel>ชื่อผู้ใช้งาน</FormLabel>
                <FormControl>
                  <Input {...field} placeholder='example@gmail.com' />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name='password'
            render={({ field }) => (
              <FormItem>
                <FormLabel>รหัสผ่าน</FormLabel>
                <FormControl>
                  <Input {...field} placeholder='password' type='password' />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <Button variant='filled' type='submit' className='w-full'>
          เข้าสู่ระบบ
        </Button>
      </form>
    </Form>
  )
}
