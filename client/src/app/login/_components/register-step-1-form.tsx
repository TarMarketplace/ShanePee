import { formatPhoneNumber } from '@/utils/input-formatter'
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
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/select'

import { type RegisterStep1FormSchema } from '../_containers/register-container'

interface RegisterStep1FormProps {
  onSubmit: SubmitHandler<RegisterStep1FormSchema>
  form: UseFormReturn<RegisterStep1FormSchema>
}

export function RegisterStep1Form({ onSubmit, form }: RegisterStep1FormProps) {
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className='flex w-full flex-col gap-8'
      >
        <div className='flex w-full flex-col gap-7'>
          <FormField
            control={form.control}
            name='name'
            render={({ field }) => (
              <FormItem>
                <FormLabel>ชื่อจริง</FormLabel>
                <FormControl>
                  <Input {...field} placeholder='กรุณากรอกชื่อจริง' />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name='surname'
            render={({ field }) => (
              <FormItem>
                <FormLabel>นามสกุล</FormLabel>
                <FormControl>
                  <Input {...field} placeholder='กรุณากรอกนามสกุล' />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name='email'
            render={({ field }) => (
              <FormItem>
                <FormLabel>อีเมล</FormLabel>
                <FormControl>
                  <Input
                    {...field}
                    placeholder='example@gmail.com'
                    type='email'
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name='phone'
            render={({ field }) => (
              <FormItem>
                <FormLabel>เบอร์โทรศัพท์</FormLabel>
                <FormControl>
                  <Input
                    placeholder='0xx-xxx-xxxx'
                    type='tel'
                    onChange={(e) => {
                      const phone = e.target.value
                        .replaceAll('-', '')
                        .slice(0, 10)
                      field.onChange(phone)
                    }}
                    value={formatPhoneNumber(field.value)}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name='gender'
            render={({ field }) => (
              <FormItem>
                <FormLabel>เพศ</FormLabel>
                <Select
                  onValueChange={field.onChange}
                  defaultValue={field.value}
                >
                  <FormControl>
                    <SelectTrigger>
                      <SelectValue placeholder='กรุณาเลือก' />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent>
                    <SelectItem value='MALE'>ชาย</SelectItem>
                    <SelectItem value='FEMALE'>หญิง</SelectItem>
                    <SelectItem value='OTHER'>ไม่ระบุ</SelectItem>
                  </SelectContent>
                </Select>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <Button variant='filled' type='submit' className='w-full'>
          ถัดไป
        </Button>
      </form>
    </Form>
  )
}
