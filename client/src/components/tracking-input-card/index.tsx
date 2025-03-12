import { type SubmitHandler, type UseFormReturn } from 'react-hook-form'

import { Text } from '@/components/text'

import { Button } from '../button'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '../form'
import { Input } from '../input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '../select'
import type { TrackingInputCardSchema } from './index.stories'

export interface TrackingInputCardProps {
  id: string
  name: string
  onSubmit: SubmitHandler<TrackingInputCardSchema>
  form: UseFormReturn<TrackingInputCardSchema>
}

const TrackingInputCard = ({
  id,
  name,
  form,
  onSubmit,
}: TrackingInputCardProps) => {
  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className='flex w-[500px] flex-col gap-5 rounded-lg bg-white p-4'
      >
        <div className='items-center'>
          <Text variant='heading-md'>เพิ่มรหัสติดตามสินค้า</Text>
        </div>
        <div className='flex flex-col gap-5'>
          <div>
            <Text variant='md-semibold' className='text-grey-500'>
              รายการสินค้า: {id}
            </Text>
            <Text variant='md-semibold'>{name}</Text>
          </div>
        </div>
        <FormField
          control={form.control}
          name='trackingNumberValue'
          render={({ field }) => (
            <FormItem>
              <FormLabel>รหัสติดตามสินค้า</FormLabel>
              <FormControl>
                <Input {...field} placeholder='กรุณากรอกรหัสติดตามสินค้า' />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name='deliveryCompanyValue'
          render={({ field }) => (
            <FormItem>
              <FormLabel>บริษัท</FormLabel>
              <Select onValueChange={field.onChange} defaultValue={field.value}>
                <FormControl>
                  <SelectTrigger>
                    <SelectValue placeholder='กรุณาเลือก' />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  <SelectItem value='Shopee express'>Shopee express</SelectItem>
                  <SelectItem value='Kerry'>Kerry</SelectItem>
                  <SelectItem value='Flash'>Flash</SelectItem>
                </SelectContent>
              </Select>
              <FormMessage />
            </FormItem>
          )}
        />
        <div className='flex justify-end pt-2'>
          <Button variant='filled' type='submit'>
            เพิ่มรหัสติดตามสินค้า
          </Button>
        </div>
      </form>
    </Form>
  )
}

export { TrackingInputCard }
