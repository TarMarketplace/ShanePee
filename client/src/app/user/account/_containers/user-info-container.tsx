import { zodResolver } from '@hookform/resolvers/zod'
import { Icon } from '@iconify/react/dist/iconify.js'
import { useRouter } from 'next/navigation'
import { type SubmitHandler, useForm } from 'react-hook-form'
import { toast } from 'sonner'
import { z } from 'zod'

import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import type { User } from '@/generated/api'
import { updateUser } from '@/generated/api'

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

interface UserInfoContainerProps {
  user: User
}

export function UserInfoContainer({ user }: UserInfoContainerProps) {
  const { fetchUser } = useUser()
  const router = useRouter()

  const form = useForm<UserInfoFormSchema>({
    resolver: zodResolver(userInfoFormSchema),
    defaultValues: {
      name: user.first_name ?? '',
      surname: user.last_name ?? '',
      email: user.email ?? '',
      phone: user.tel ?? '',
      gender: user.gender as UserInfoFormSchema['gender'],
    },
  })

  const onSubmit: SubmitHandler<UserInfoFormSchema> = async (data) => {
    const { response, error } = await updateUser({
      body: {
        first_name: data.name,
        last_name: data.surname,
        tel: data.phone,
        gender: data.gender,
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

  const handleChangePicture = () => {
    console.log('handleChangePicture')
  }

  return (
    <div className='my-6 flex flex-col gap-8 md:m-6 md:w-full'>
      <div className='hidden items-center gap-2 md:flex'>
        <Icon icon='mdi:account' className='size-10' />
        <Text variant='heading-lg'>บัญชีของฉัน</Text>
      </div>
      <UserInfoForm
        onSubmit={onSubmit}
        handleChangePicture={handleChangePicture}
        form={form}
      />
    </div>
  )
}
