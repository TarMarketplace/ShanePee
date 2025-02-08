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

import { type PasswordFormSchema } from '../_containers/password-container'

interface PasswordFormProps {
  onSubmit: SubmitHandler<PasswordFormSchema>
  form: UseFormReturn<PasswordFormSchema>
}

export function PasswordForm({ onSubmit, form }: PasswordFormProps) {
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className='flex w-full flex-col gap-8'
      >
        <div className='grid w-full grid-cols-1 gap-4'>
          <div className='col-span-1'>
            <FormField
              control={form.control}
              name='oldPassword'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>รหัสผ่านเดิม</FormLabel>
                  <FormControl>
                    <Input {...field} placeholder='กรุณากรอกรหัสผ่านเดิม' />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <div className='col-span-1'>
            <FormField
              control={form.control}
              name='password'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>รหัสผ่านใหม่</FormLabel>
                  <FormControl>
                    <Input {...field} placeholder='กรุณากรอกรหัสผ่านใหม่' />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <div className='col-span-1'>
            <FormField
              control={form.control}
              name='confirmPassword'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>ยืนยันรหัสผ่านใหม่</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder='กรุณากรอกรหัสผ่านใหม่อีกครั้ง'
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
        </div>
        <div className='flex justify-end'>
          <Button variant='filled' type='submit' className=''>
            บันทึก
          </Button>
        </div>
      </form>
    </Form>
  )
}
