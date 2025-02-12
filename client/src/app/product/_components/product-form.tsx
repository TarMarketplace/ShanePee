import { Icon } from '@iconify/react'
import Image from 'next/image'
import { useMemo, useState } from 'react'
import type { SubmitHandler, UseFormReturn } from 'react-hook-form'

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
import { Text } from '@/components/text'
import { Textarea } from '@/components/textarea'

import { cn } from '@/lib/utils'

import { type ProductFormSchema } from '../_containers/product-form-container'

type ProductFormProps = {
  form: UseFormReturn<ProductFormSchema>
  onSubmit: SubmitHandler<ProductFormSchema>
  isEditing: boolean
}

export function ProductForm({ form, onSubmit, isEditing }: ProductFormProps) {
  const [isShowingButton, setIsShowingButton] = useState(false)

  const image = form.watch('image')

  const previewUrl = useMemo(() => {
    if (image.size > 0) {
      return URL.createObjectURL(image)
    }
  }, [image])

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className='mx-auto flex w-full max-w-screen-lg flex-col gap-2.5 p-1.5'
      >
        <div className='flex items-center gap-0.5'>
          <Icon
            icon={isEditing ? 'bxs:edit' : 'mdi:plus'}
            width={32}
            height={32}
          />
          <Text variant='heading-lg'>
            {isEditing ? 'แก้ไขสินค้า' : 'วางจำหน่าย Art Toy ใหม่'}
          </Text>
        </div>

        <div className='flex w-full flex-col justify-center gap-4 px-1 sm:flex-row sm:px-0'>
          <FormField
            control={form.control}
            name='image'
            render={({ field }) => (
              <FormItem>
                <FormLabel
                  isInput={false}
                  className='relative flex aspect-[1.6/1] w-full cursor-pointer items-center justify-center overflow-hidden rounded-lg border-2 border-grey-300 sm:w-[640px]'
                >
                  {previewUrl && (
                    <Image
                      src={previewUrl}
                      alt=''
                      fill
                      className='size-full object-cover'
                      onMouseEnter={() => setIsShowingButton(true)}
                      onMouseLeave={() => setIsShowingButton(false)}
                    />
                  )}
                  <Button
                    variant='filled'
                    type='button'
                    className={cn(
                      'pointer-events-none z-50 opacity-0',
                      (!previewUrl || isShowingButton) && 'opacity-100'
                    )}
                  >
                    {previewUrl ? 'เปลี่ยนรูป' : 'เลือกรูป'}
                  </Button>
                </FormLabel>
                <FormControl>
                  <input
                    type='file'
                    className='hidden'
                    accept='image/*'
                    onChange={(e) => {
                      const file = e.target.files?.[0]
                      if (file) {
                        field.onChange(file)
                      }
                    }}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <div className='flex w-full flex-col gap-3 sm:h-full sm:w-[614px]'>
            <FormField
              name='name'
              control={form.control}
              render={({ field }) => (
                <FormItem>
                  <FormLabel>ชื่อสินค้า</FormLabel>
                  <FormControl>
                    <Input {...field} type='text' placeholder='ชื่อสินค้า' />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              name='description'
              control={form.control}
              render={({ field }) => (
                <FormItem>
                  <FormLabel>รายละเอียด</FormLabel>
                  <FormControl>
                    <Textarea {...field} placeholder='รายละเอียด' />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              name='price'
              control={form.control}
              render={({ field }) => (
                <FormItem>
                  <FormLabel>ราคา</FormLabel>
                  <Text variant='md-regular' className='absolute left-3 top-2'>
                    ฿
                  </Text>
                  <FormControl>
                    <Input
                      {...field}
                      onChange={(e) => {
                        if (e.target.value) {
                          field.onChange(parseInt(e.target.value))
                        } else {
                          field.onChange('')
                        }
                      }}
                      type='number'
                      placeholder='ราคา'
                      className='pl-7'
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
        </div>

        <div className='mt-auto flex items-center justify-end gap-4'>
          {isEditing && (
            <button type='button' className='font-bold text-grey-500 underline'>
              ยกเลิก
            </button>
          )}
          <Button variant='filled' type='submit'>
            {isEditing ? 'แก้ไข' : 'วางจำหน่าย'}
          </Button>
        </div>
      </form>
    </Form>
  )
}
