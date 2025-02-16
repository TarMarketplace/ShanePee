import { zodResolver } from '@hookform/resolvers/zod'
import { Icon } from '@iconify/react/dist/iconify.js'
import { useRouter } from 'next/navigation'
import { useEffect } from 'react'
import { type SubmitHandler, useForm } from 'react-hook-form'
import { toast } from 'sonner'
import { z } from 'zod'

import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import { updateUser } from '@/generated/api'

import { AddressForm } from '../_components/address-form'

const addressFormSchema = z.object({
  details: z.string().min(1, 'Details address is required'),
  district: z.string().min(1, 'Distric is required'),
  province: z.string().min(1, 'Province is required'),
  postalCode: z
    .string()
    .min(5, 'Postal code is required')
    .regex(/^\d{5}$/, 'Invalid postal code'),
})

export type AddressFormSchema = z.infer<typeof addressFormSchema>

export function AddressContainer() {
  const { user, fetchUser } = useUser()
  const router = useRouter()

  const form = useForm<AddressFormSchema>({
    resolver: zodResolver(addressFormSchema),
    defaultValues: {
      details: '',
      district: '',
      province: '',
      postalCode: '',
    },
  })

  useEffect(() => {
    if (user) {
      form.reset({
        details: user.address.house_no ?? '',
        district: user.address.district ?? '',
        province: user.address.province ?? '',
        postalCode: user.address.postcode ?? '',
      })
    }
  }, [user, form])

  const onSubmit: SubmitHandler<AddressFormSchema> = async (data) => {
    const { response, error } = await updateUser({
      body: {
        address: {
          house_no: data.details,
          district: data.district,
          province: data.province,
          postcode: data.postalCode,
        },
      },
    })

    if (response.ok) {
      toast.success('Updated successfully')
      fetchUser()
    } else if (response.status == 401) {
      router.push('/login')
    } else {
      toast.error('Something went wrong')
      toast.error(error?.detail)
    }
  }

  return (
    <div className='my-6 flex flex-col gap-8 md:m-6 md:w-full'>
      <div className='hidden items-center gap-2 md:flex'>
        <Icon icon='ic:round-home' className='size-10' />
        <Text variant='heading-lg'>ที่อยู่สำหรับการจัดส่งสินค้า</Text>
      </div>
      <AddressForm onSubmit={onSubmit} form={form} />
    </div>
  )
}
