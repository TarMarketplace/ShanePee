import { Icon } from '@iconify/react/dist/iconify.js'
import Image from 'next/image'
import Link from 'next/link'
import { useEffect, useRef, useState } from 'react'

import { Button } from '@/components/button'
import { Input } from '@/components/input'
import { Text } from '@/components/text'

import type { ChatMessage } from '@/generated/api'

export interface ChatProps {
  sender_id: number | null
  sellerName: string
  chat: ChatMessage[]
  input: string
  setInput: (input: string) => void
  previewImages: string[]
  handleImageUpload: (event: React.ChangeEvent<HTMLInputElement>) => void
  removeImage: (index: number) => void
  handleBackButton: () => void
  handleSendMessage: () => void
}

function Chat({
  sender_id,
  sellerName,
  chat,
  input,
  setInput,
  previewImages,
  handleImageUpload,
  removeImage,
  handleBackButton,
  handleSendMessage,
}: ChatProps) {
  const [imageDimensions, setImageDimensions] = useState<
    { width: number; height: number }[]
  >([])
  const messagesEndRef = useRef<HTMLDivElement | null>(null)

  const handleDownload = (url: string) => {
    const link = document.createElement('a')
    link.href = url
    link.download = 'downloaded-image'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
  }

  useEffect(() => {
    const fetchImageDimensions = async () => {
      const dimensions = await Promise.all(
        previewImages.map(async (image) => {
          const img = new window.Image()
          img.src = image
          await img.decode()
          return { width: img.width, height: img.height }
        })
      )
      setImageDimensions(dimensions)
    }

    fetchImageDimensions()
  }, [previewImages])

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'auto' })
  }, [chat, previewImages])

  return (
    <div className='flex min-h-[calc(100dvh-62px-256px)] w-full flex-col divide-y-2 divide-grey-200 truncate'>
      <span className='flex h-12 items-center p-4'>
        <Icon
          icon='weui:back-filled'
          className='mr-2 size-6 cursor-pointer md:hidden'
          onClick={() => handleBackButton()}
        />
        <Text variant='lg-semibold' className='w-full'>
          {sellerName}
        </Text>
        <Link href={'/seller/' + sender_id}>
          <Button variant='filled' className='h-8'>
            <Icon icon='tdesign:store-filled' className='size-5' />
            <Text variant='md-regular'>หน้าร้านค้า</Text>
          </Button>
        </Link>
      </span>
      <div className='h-0 grow overflow-y-auto'>
        {chat?.map((message) => {
          if (message) {
            if (message.sender_id == sender_id) {
              return (
                <div className='flex w-full justify-end p-4' key={message.id}>
                  <div className='relative mb-3 max-w-[60%] text-wrap rounded-lg bg-primary-500 text-white shadow'>
                    {message.message_type == 'IMAGE' ? (
                      <Image
                        src={message.content}
                        alt='chat image'
                        width={400}
                        height={400}
                        className='size-auto cursor-pointer rounded-lg'
                        onClick={() => handleDownload(message.content)}
                      />
                    ) : (
                      <div className='p-2'>message.content</div>
                    )}
                    <Text className='absolute bottom-[-30px] right-2 text-grey-500'>
                      {String(new Date(message.created_at).getHours()).padStart(
                        2,
                        '0'
                      )}
                      :
                      {String(
                        new Date(message.created_at).getMinutes()
                      ).padStart(2, '0')}
                    </Text>
                  </div>
                </div>
              )
            } else {
              return (
                <div className='flex w-full justify-start p-4' key={message.id}>
                  <div className='relative mb-3 max-w-[60%] text-wrap rounded-lg bg-secondary-100 shadow'>
                    {message.message_type == 'IMAGE' ? (
                      <Image
                        src={message.content}
                        alt='chat image'
                        width={400}
                        height={400}
                        className='size-auto cursor-pointer rounded-lg'
                        onClick={() => handleDownload(message.content)}
                      />
                    ) : (
                      <div className='p-2'>message.content</div>
                    )}
                    <Text className='absolute bottom-[-30px] left-2 text-grey-500'>
                      {String(new Date(message.created_at).getHours()).padStart(
                        2,
                        '0'
                      )}
                      :
                      {String(
                        new Date(message.created_at).getMinutes()
                      ).padStart(2, '0')}
                    </Text>
                  </div>
                </div>
              )
            }
          }
        })}
        <div ref={messagesEndRef} />
      </div>

      {previewImages.length > 0 && (
        <div className='flex gap-2 overflow-x-auto p-2'>
          {previewImages.map((image, index) => {
            const { width, height } = imageDimensions[index] || {}
            if (!width || !height) return null

            const aspectRatio = width / height
            const newWidth = 192 * aspectRatio

            return (
              <div key={index} className='relative h-48 shrink-0'>
                <Image
                  src={image}
                  alt='Preview'
                  height={192}
                  width={newWidth}
                  className='h-48 rounded-lg'
                />
                <div
                  className='absolute right-2 top-2 flex size-6 cursor-pointer items-center justify-center rounded-full bg-error-500'
                  onClick={() => removeImage(index)}
                >
                  <Icon
                    icon='solar:trash-bin-trash-bold'
                    className='size-4 text-white'
                  />
                </div>
              </div>
            )
          })}
        </div>
      )}

      <span className='flex h-12 items-center justify-between gap-2 p-2'>
        <Input
          placeholder='ส่งข้อความ'
          className='truncate border-transparent focus:border-none focus:border-transparent focus:ring-0'
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === 'Enter' && !e.shiftKey) {
              e.preventDefault()
              handleSendMessage()
            }
          }}
        ></Input>
        <label htmlFor='image-upload' className='cursor-pointer'>
          <Icon icon='ic:baseline-image' className='size-8' />
          <input
            type='file'
            id='image-upload'
            className='hidden'
            accept='image/*'
            multiple
            onChange={handleImageUpload}
          />
        </label>
        <Icon
          icon='ic:baseline-send'
          className='size-8 cursor-pointer'
          onClick={() => handleSendMessage()}
        />
      </span>
    </div>
  )
}

export { Chat }
