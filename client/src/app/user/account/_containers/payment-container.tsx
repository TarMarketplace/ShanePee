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

import { PaymentForm } from '../_components/payment-form'

const paymentFormSchema = z.object({
  cardNumber: z
    .string()
    .length(16, 'Card number must be exactly 16 digits')
    .regex(/^\d{16}$/, 'Card number must contain only digits'),
  cardHolderName: z.string().min(1, 'Cardholder name is required'),
  expirationDate: z
    .string()
    .length(5, 'Expiration date must be in MM/YY format')
    .regex(/^(0[1-9]|1[0-2])\/\d{2}$/, 'Invalid expiration date format'),
  cvv: z
    .string()
    .length(3, 'CVV must be exactly 3 digits')
    .regex(/^\d{3}$/, 'CVV must contain only digits'),
})

export type PaymentFormSchema = z.infer<typeof paymentFormSchema>

export function PaymentContainer() {
  const { user, fetchUser } = useUser()
  const router = useRouter()

  const form = useForm<PaymentFormSchema>({
    resolver: zodResolver(paymentFormSchema),
    defaultValues: {
      cardNumber: '',
      cardHolderName: '',
      expirationDate: '',
      cvv: '',
    },
  })

  useEffect(() => {
    if (user) {
      form.reset({
        cardNumber: user.payment_method.card_number ?? '',
        cardHolderName: user.payment_method.card_owner ?? '',
        expirationDate: user.payment_method.expire_date ?? '',
        cvv: user.payment_method.cvv ?? '',
      })
    }
  }, [user, form])

  const onSubmit: SubmitHandler<PaymentFormSchema> = async (data) => {
    const res = await fetch(`${env.NEXT_PUBLIC_BASE_API_URL}/user`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify({
        payment_method: {
          card_number: data.cardNumber,
          card_owner: data.cardHolderName,
          expire_date: data.expirationDate,
          cvv: data.cvv,
        },
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

  return (
    <div className='my-6 flex flex-col gap-8 md:m-6 md:w-full'>
      <div className='hidden items-center gap-2 md:flex'>
        <Icon icon='famicons:card' className='size-10' />
        <Text variant='heading-lg'>ช่องทางการชำระเงิน</Text>
      </div>
      <PaymentForm onSubmit={onSubmit} form={form} />
    </div>
  )
}
