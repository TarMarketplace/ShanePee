'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { type SubmitHandler, useForm } from 'react-hook-form'
import { toast } from 'sonner'
import { z } from 'zod'

import { useUser } from '@/providers/user-provider'

import { env } from '@/env'

import { ProductForm } from '../_components/product-form'

const productFormSchema = z.object({
  name: z.string().min(1, 'Name is required'),
  description: z.string().min(1, 'Description is required'),
  price: z.number().min(1, 'Price is required'),
  image: z
    .instanceof(File)
    .refine((file) => file.size !== 0, 'Please upload an image'),
})

export type ProductFormSchema = z.infer<typeof productFormSchema>

interface ProductContainerProps {
  id?: number
  defaultValues?: ProductFormSchema
}

export function ProductFormContainer({
  id,
  defaultValues,
}: ProductContainerProps) {
  const { user } = useUser()

  const form = useForm<ProductFormSchema>({
    resolver: zodResolver(productFormSchema),
    defaultValues: defaultValues ?? {
      name: '',
      description: '',
      price: 0,
      image: new File([], ''),
    },
  })

  const onSubmit: SubmitHandler<ProductFormSchema> = async (data) => {
    if (!user) {
      toast.error('Please login first')
      return
    }

    const reader = new FileReader()
    reader.readAsDataURL(data.image)
    reader.onloadend = async () => {
      const imageAsBase64 = reader.result as string

      const response = await fetch(
        `${env.NEXT_PUBLIC_BASE_API_URL}/arttoy` + (id ? `/${id}` : ''),
        {
          method: id ? 'PUT' : 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            name: data.name,
            description: data.description,
            price: data.price,
            photo: imageAsBase64,
            owner_id: user.id,
          }),
        }
      )

      if (response.ok) {
        toast.success('Product created successfully')
        form.reset()
      } else {
        toast.error('Failed to create product')
      }
    }
  }

  return (
    <main className='grid size-full grid-cols-1 place-items-center bg-white p-4 md:p-12'>
      <ProductForm
        onSubmit={onSubmit}
        form={form}
        isEditing={!!defaultValues}
      />
    </main>
  )
}
