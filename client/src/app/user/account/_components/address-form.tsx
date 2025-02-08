import { formatPostalCode } from '@/utils/input-formatter'
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

import { type AddressFormSchema } from '../_containers/address-container'

interface AddressFormProps {
  onSubmit: SubmitHandler<AddressFormSchema>
  form: UseFormReturn<AddressFormSchema>
}

export function AddressForm({ onSubmit, form }: AddressFormProps) {
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className='flex w-full flex-col gap-8'
      >
        <div className='grid w-full grid-cols-1 gap-4 md:grid-cols-3'>
          <div className='col-span-1 md:col-span-3'>
            <FormField
              control={form.control}
              name='details'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>บ้านเลขที่, ซอย, หมู่, ถนน, แขวง/ตำบล</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder='กรุณากรอกบ้านเลขที่, ซอย, หมู่, ถนน, แขวง/ตำบล'
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <div className='col-span-1'>
            <FormField
              control={form.control}
              name='distric'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>เขต/อำเภอ</FormLabel>
                  <FormControl>
                    <Input {...field} placeholder='กรุณากรอกเขต/อำเภอ' />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <div className='col-span-1'>
            <FormField
              control={form.control}
              name='province'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>จังหวัด</FormLabel>
                  <FormControl>
                    <Input {...field} placeholder='กรุณากรอกจังหวัด' />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <div className='col-span-1'>
            <FormField
              control={form.control}
              name='postalCode'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>รหัสไปรษณีย์</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder='กรุณากรอกรหัสไปรษณีย์'
                      onChange={(e) => {
                        const value = e.target.value.slice(0, 5)
                        field.onChange(value)
                      }}
                      value={formatPostalCode(field.value)}
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
