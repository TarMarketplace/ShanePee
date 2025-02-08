import { zodResolver } from '@hookform/resolvers/zod'
import { Icon } from '@iconify/react/dist/iconify.js'
import { type SubmitHandler, useForm } from 'react-hook-form'
import { z } from 'zod'

import { Text } from '@/components/text'

import { AddressForm } from '../_components/address-form'

const addressFormSchema = z.object({
  details: z.string().min(1, 'Details address is required'),
  distric: z.string().min(1, 'Distric is required'),
  province: z.string().min(1, 'Province is required'),
  postalCode: z
    .string()
    .min(5, 'Postal code is required')
    .regex(/^\d{5}$/, 'Invalid postal code'),
})

export type AddressFormSchema = z.infer<typeof addressFormSchema>

export function AddressContainer() {
  const form = useForm<AddressFormSchema>({
    resolver: zodResolver(addressFormSchema),
    defaultValues: {
      details: '',
      distric: '',
      province: '',
      postalCode: '',
    },
  })

  const onSubmit: SubmitHandler<AddressFormSchema> = (data) => {
    console.log(data)
  }

  return (
    <div className='m-6 flex w-full flex-col gap-8'>
      <div className='flex items-center gap-2'>
        <Icon icon='ic:round-home' className='size-10' />
        <Text variant='heading-lg'>ที่อยู่สำหรับการจัดส่งสินค้า</Text>
      </div>
      <AddressForm onSubmit={onSubmit} form={form} />
    </div>
  )
}
