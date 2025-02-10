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

import { type RegisterStep2FormSchema } from '../_containers/register-container'

interface RegisterStep2FormProps {
  onSubmit: SubmitHandler<RegisterStep2FormSchema>
  form: UseFormReturn<RegisterStep2FormSchema>
}

export function RegisterStep2Form({ onSubmit, form }: RegisterStep2FormProps) {
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className='flex w-full flex-col gap-8'
      >
        <div className='flex w-full flex-col gap-7'>
          <FormField
            control={form.control}
            name='email'
            render={({ field }) => (
              <FormItem>
                <FormLabel>อีเมล</FormLabel>
                <FormControl>
                  <Input {...field} placeholder='กรุณากรอกอีเมล' />
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
                  <Input {...field} placeholder='รหัสผ่าน' type='password' />
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
                  <Input {...field} placeholder='รหัสผ่าน' type='password' />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <Button variant='filled' type='submit' className='w-full'>
          สมัครใช้งาน
        </Button>
      </form>
    </Form>
  )
}
