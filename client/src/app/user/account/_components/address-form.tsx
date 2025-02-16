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

import { formatPostalCode } from '@/utils/input-formatter'

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
                    <div>
                      <Input
                        {...field}
                        className='hidden md:block'
                        placeholder='กรุณากรอกบ้านเลขที่, ซอย, หมู่, ถนน, แขวง/ตำบล'
                      />
                      <textarea
                        {...field}
                        className='peer flex w-full rounded-md border border-input bg-transparent px-3 py-2 text-base file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:border-primary-500 focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50 md:hidden'
                        rows={3}
                        placeholder='กรุณากรอกบ้านเลขที่, ซอย, หมู่, ถนน, แขวง/ตำบล'
                      />
                    </div>
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <div className='col-span-1'>
            <FormField
              control={form.control}
              name='district'
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
          <Button variant='filled' type='submit'>
            บันทึก
          </Button>
        </div>
      </form>
    </Form>
  )
}
