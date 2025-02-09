import { zodResolver } from '@hookform/resolvers/zod'
import { Icon } from '@iconify/react/dist/iconify.js'
import { type SubmitHandler, useForm } from 'react-hook-form'
import { z } from 'zod'

import { Text } from '@/components/text'

import { UserInfoForm } from '../_components/user-info-form'

const userInfoFormSchema = z.object({
  username: z.string().min(1, 'Username is required'),
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
  const form = useForm<UserInfoFormSchema>({
    resolver: zodResolver(userInfoFormSchema),
    defaultValues: {
      username: '',
      name: '',
      surname: '',
      email: '',
      phone: '',
      gender: undefined,
    },
  })

  const onSubmit: SubmitHandler<UserInfoFormSchema> = (data) => {
    console.log(data)
  }

  const handleChangePicture = () => {
    console.log('habdleChangePicture')
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
