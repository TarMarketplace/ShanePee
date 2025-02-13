import { zodResolver } from '@hookform/resolvers/zod'
import { Icon } from '@iconify/react/dist/iconify.js'
import { useRouter } from 'next/navigation'
import { useEffect } from 'react'
import { type SubmitHandler, useForm } from 'react-hook-form'
import { toast } from 'sonner'
import { z } from 'zod'

import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import { env } from '@/env'

import { UserInfoForm } from '../_components/user-info-form'

const userInfoFormSchema = z.object({
  name: z.string().min(1, 'Name is required'),
  surname: z.string().min(1, 'Surname is required'),
  gender: z.enum(['MALE', 'FEMALE', 'OTHER']),
  email: z.string().email('Invalid email'),
  phone: z
    .string()
    .min(10, 'Phone number is required')
    .refine((value) => {
      return /^\d+$/.test(value)
    }),
})

export type UserInfoFormSchema = z.infer<typeof userInfoFormSchema>

export function UserInfoContainer() {
  const { user, fetchUser } = useUser()
  const router = useRouter()

  const form = useForm<UserInfoFormSchema>({
    resolver: zodResolver(userInfoFormSchema),
    defaultValues: {
      name: '',
      surname: '',
      email: '',
      phone: '',
    },
  })

  useEffect(() => {
    if (user) {
      form.reset({
        name: user.first_name ?? '',
        surname: user.last_name ?? '',
        email: user.email ?? '',
        phone: user.tel ?? '',
        gender: user.gender as 'MALE' | 'FEMALE' | 'OTHER',
      })
    }
  }, [user, form])

  const onSubmit: SubmitHandler<UserInfoFormSchema> = async (data) => {
    const res = await fetch(`${env.NEXT_PUBLIC_BASE_API_URL}/user`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify({
        first_name: data.name,
        last_name: data.surname,
        email: data.email,
        tel: data.phone,
        gender: data.gender,
      }),
    })

    if (res.ok) {
      toast.success('Updated successfully')
      fetchUser()
    } else if (res.status == 401) {
      router.push('/login')
    } else {
      toast.error('Something went wrong')
    }
  }

  const handleChangePicture = () => {
    console.log('handleChangePicture')
  }

  return (
    <div className='my-6 flex flex-col gap-8 md:m-6 md:w-full'>
      <div className='hidden items-center gap-2 md:flex'>
        <Icon icon='mdi:account' className='size-10' />
        <Text variant='heading-lg'>บัญชีของฉัน</Text>
      </div>
      {form.getValues('gender') && ( // Leon, please fix this. I spent thousands of years with Tar and Boom trying to debug this but no luck at all 😔
        <UserInfoForm
          onSubmit={onSubmit}
          handleChangePicture={handleChangePicture}
          form={form}
        />
      )}
    </div>
  )
}
