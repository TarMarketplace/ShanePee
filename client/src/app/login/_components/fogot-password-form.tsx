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
import { Text } from '@/components/text'

import { type ForgotPasswordFormSchema } from '../_containers/forgot-password-container'

interface ForgotPasswordFormProps {
  onSubmit: SubmitHandler<ForgotPasswordFormSchema>
  form: UseFormReturn<ForgotPasswordFormSchema>
  onSwitchMode: () => void
}

export function ForgotPasswordForm({
  onSubmit,
  form,
  onSwitchMode,
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
            name='email'
            render={({ field }) => (
              <FormItem>
                <FormLabel>อีเมล</FormLabel>
                <FormControl>
                  <Input {...field} placeholder='กรุณากรอกชื่อผู้ใช้งาน' />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <div className='flex flex-col gap-2.5'>
          <Button variant='filled' type='submit' className='w-full'>
            ส่งอีเมลเปลี่ยนรหัสผ่าน
          </Button>
          <button onClick={onSwitchMode}>
            <Text variant='sm-regular' className='text-muted-foreground'>
              ย้อนกลับ
            </Text>
          </button>
        </div>
      </form>
    </Form>
  )
}
