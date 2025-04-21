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

import { type ForgotPasswordFormSchema } from '../_containers/forgot-password-container'

interface ForgotPasswordFormProps {
  onSubmit: SubmitHandler<ForgotPasswordFormSchema>
  form: UseFormReturn<ForgotPasswordFormSchema>
}

export function ForgotPasswordForm({
  onSubmit,
  form,
}: ForgotPasswordFormProps) {
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className='flex w-full flex-col gap-8'
      >
        <div className='flex w-full flex-col gap-7'>
          <FormField
            control={form.control}
            name='password'
            render={({ field }) => (
              <FormItem>
                <FormLabel>รหัสผ่าน</FormLabel>
                <FormControl>
                  <Input
                    {...field}
                    placeholder='กรุณากรอกรหัสผ่าน'
                    type='password'
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name='confirmPassword'
            render={({ field }) => (
              <FormItem>
                <FormLabel>ยืนยันรหัสผ่าน</FormLabel>
                <FormControl>
                  <Input
                    {...field}
                    placeholder='กรุณากรอกรหัสผ่านใหม่อีกครั้ง'
                    type='password'
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <div className='flex flex-col gap-2.5'>
          <Button variant='filled' type='submit' className='w-full'>
            ตั้งค่ารหัสผ่านใหม่
          </Button>
        </div>
      </form>
    </Form>
  )
}
