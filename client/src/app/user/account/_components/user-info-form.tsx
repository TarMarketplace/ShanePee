import Image from 'next/image'
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
import { Text } from '@/components/text'

import { formatPhoneNumber } from '@/utils/input-formatter'

import { type UserInfoFormSchema } from '../_containers/user-info-container'

interface UserInfoFormProps {
  onSubmit: SubmitHandler<UserInfoFormSchema>
  handleChangePicture: () => void
  form: UseFormReturn<UserInfoFormSchema>
}

export function UserInfoForm({
  onSubmit,
  handleChangePicture,
  form,
}: UserInfoFormProps) {
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className='flex w-full flex-col gap-8'
      >
        <div className='flex flex-col items-center gap-4 md:flex-row'>
          <div className='w-32'>
            <Image
              src='https://placehold.co/128x128.png'
              alt=''
              width={128}
              height={128}
            />
            <button
              className={`mt-2 inline-flex h-10 w-full items-center justify-center gap-2 rounded-sm bg-primary px-3 py-2 text-white`}
              onClick={() => handleChangePicture()}
            >
              <Text variant='sm-semibold'>เปลี่ยนรูป</Text>
            </button>
          </div>
          <div className='grid w-full grid-cols-2 gap-4 md:grid-cols-6'>
            <div className='col-span-2 md:col-span-6'>
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
            </div>
            <div className='col-span-1 md:col-span-2'>
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
            </div>
            <div className='col-span-1 md:col-span-2'>
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
            </div>
            <div className='col-span-2'>
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
            <div className='col-span-2 md:col-span-3'>
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
            </div>
            <div className='col-span-2 md:col-span-3'>
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
            </div>
          </div>
        </div>
        <div className='flex justify-end'>
          <Button variant='filled' type='submit'>
            บันทึก
          </Button>
        </div>
      </form>
    </Form>
  )
}
