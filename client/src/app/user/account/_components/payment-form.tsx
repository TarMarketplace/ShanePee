import {
  formatCVV,
  formatCardNumber,
  formatExpirationDate,
} from '@/utils/input-formatter'
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

import { type PaymentFormSchema } from '../_containers/payment-container'

interface PaymentFormProps {
  onSubmit: SubmitHandler<PaymentFormSchema>
  form: UseFormReturn<PaymentFormSchema>
}

export function PaymentForm({ onSubmit, form }: PaymentFormProps) {
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className='flex w-full flex-col gap-8'
      >
        <div className='grid w-full grid-cols-2 gap-4'>
          <div className='col-span-2'>
            <FormField
              control={form.control}
              name='cardNumber'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>หมายเลขบัตร</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder='xxxx-xxxx-xxxx-xxxx'
                      onChange={(e) => {
                        const value = e.target.value
                          .replaceAll(' ', '')
                          .slice(0, 16)
                        field.onChange(value)
                      }}
                      value={formatCardNumber(field.value)}
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
              name='expirationDate'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>วันหมดอายุ (ดด/ปป)</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder='xx/xx'
                      onChange={(e) => {
                        const value = e.target.value.slice(0, 5)
                        field.onChange(value)
                      }}
                      value={formatExpirationDate(field.value)}
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
              name='cvv'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>CVV</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder='xxx'
                      onChange={(e) => {
                        const value = e.target.value
                          .replaceAll(' ', '')
                          .slice(0, 3)
                        field.onChange(value)
                      }}
                      value={formatCVV(field.value)}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <div className='col-span-2'>
            <FormField
              control={form.control}
              name='cardHolderName'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>ชื่อเจ้าของบัตร</FormLabel>
                  <FormControl>
                    <Input {...field} placeholder='กรุณากรอกชื่อเจ้าของบัตร' />
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
