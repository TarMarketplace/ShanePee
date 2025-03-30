'use client'

import { Icon } from '@iconify/react/dist/iconify.js'
import imageCompression from 'browser-image-compression'
import { useRouter } from 'next/navigation'
import { useCallback, useEffect, useRef, useState } from 'react'
import { toast } from 'sonner'

import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import { cn } from '@/lib/utils'

import type { ChatMessage } from '@/generated/api'
import { getChatDetail, pollMessage, sendMessage } from '@/generated/api'

import { Chat } from '../_components/chat'
import { ChatBox } from '../_components/chat-box'

export function ChatListContainer() {
  const [activeChat, setActiveChat] = useState(0)
  const [chat, setChat] = useState<ChatMessage[]>([])
  const [input, setInput] = useState('')
  const [previewImages, setPreviewImages] = useState<string[]>([])
  const pollChatIdRef = useRef(0)
  const isPollingRef = useRef(false)
  const { user } = useUser()
  const router = useRouter()

  const chats = [
    {
      id: 1,
      photo: 'photo',
      sellerName: 'John Do b@b',
      date: new Date(),
      message: 'last message',
      user_id: 3523329031,
    },
    {
      id: 2,
      photo: 'photo',
      sellerName: 'John Doe a@b',
      date: new Date(),
      message: 'last message',
      user_id: 4105047913,
    },
  ]

  const handleImageUpload = async (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    if (event.target.files) {
      const files = Array.from(event.target.files)
      const imageUrls: string[] = []

      for (const file of files) {
        try {
          const compressedFile = await imageCompression(file, {
            maxWidthOrHeight: 800,
            maxSizeMB: 1,
          })

          const base64String = await convertToBase64(compressedFile)
          imageUrls.push(base64String)
        } catch (error) {
          console.error('Error during image compression:', error)
        }
      }

      setPreviewImages((prev) => [...prev, ...imageUrls])
    }
  }

  const convertToBase64 = (file: File): Promise<string> => {
    return new Promise((resolve, reject) => {
      const reader = new FileReader()
      reader.readAsDataURL(file)
      reader.onload = () => resolve(reader.result as string)
      reader.onerror = (error) => reject(error)
    })
  }

  const removeImage = (index: number) => {
    setPreviewImages((prev) => prev.filter((_, i) => i !== index))
  }

  const handleSendMessage = () => {
    if (previewImages.length > 0) {
      previewImages.forEach((image) => {
        sendMessage({
          path: {
            userID: activeChat,
          },
          body: {
            content: image,
            sender: 'BUYER',
          },
        })
          .then((response) => {
            if (response.data) {
              setInput('')
              setChat((prevChat) => [...prevChat, response.data])
              pollChatIdRef.current = response.data.id
              toast.success('Image sent')
            } else {
              toast.error('Error sending image')
            }
          })
          .catch(() => {
            toast.error('Error sending image')
          })
      })
      setPreviewImages([])
    }

    sendMessage({
      path: {
        userID: activeChat,
      },
      body: {
        content: input,
        sender: 'BUYER',
      },
    })
      .then((response) => {
        if (response.data) {
          setInput('')
          setChat((prevChat) => [...prevChat, response.data])
          pollChatIdRef.current = response.data.id
          toast.success('Message sent')
        } else {
          toast.error('Error sending message')
        }
      })
      .catch(() => {
        toast.error('Error sending message')
      })
  }

  const pollChat = useCallback(async () => {
    if (!isPollingRef.current) return

    await pollMessage({
      path: {
        userID: activeChat,
      },
      query: {
        as: 'BUYER',
        chatID: pollChatIdRef.current,
      },
    })
      .then((response) => {
        const message = response?.data?.data
        if (Array.isArray(message)) {
          setChat((prevChat) => [...prevChat, ...message])
          pollChatIdRef.current = message[message.length - 1].id
        } else if (response.response.status == 401) {
          isPollingRef.current = false
          router.push('/login')
        } else {
          toast.error('Something went wrong')
        }
      })
      .then(() => {
        pollChat()
      })
      .catch(() => {
        toast.error('Something went wrong')
      })
  }, [activeChat, router])

  const getChat = useCallback(async () => {
    isPollingRef.current = true

    await getChatDetail({
      path: {
        userID: activeChat,
      },
    })
      .then((response) => {
        const message = response?.data?.data
        if (Array.isArray(message)) {
          if (message.length < 1) return
          setChat(message)
          pollChatIdRef.current = message[message.length - 1].id
        } else if (response.response.status == 401) {
          isPollingRef.current = false
          router.push('/login')
        } else {
          console.log(message)
          toast.error('Something went wrong')
        }
      })
      .then(() => {
        pollChat()
      })
      .catch((e) => {
        console.log(e)
        toast.error('Something went wrong')
      })
  }, [activeChat, pollChat, router])

  useEffect(() => {
    if (activeChat) {
      setChat([])
      if (activeChat != user?.id) getChat()
    }
  }, [activeChat, user?.id, getChat])

  const selectedChat = chats.find((chat) => chat?.id == activeChat)

  return (
    <div className='flex min-h-[calc(100dvh-60px-236px)] divide-x-2 divide-grey-200'>
      <div
        className={cn(
          'w-full flex-col md:min-w-[350px] md:max-w-[350px]',
          activeChat == 0 ? 'flex' : 'hidden md:flex'
        )}
      >
        {chats.map((chat) => {
          if (chat) {
            return (
              <div
                key={chat.id}
                onClick={() =>
                  setActiveChat(activeChat == chat.user_id ? 0 : chat.user_id)
                }
                className='cursor-pointer'
              >
                <ChatBox
                  photo={chat.photo}
                  sellerName={chat.sellerName}
                  date={chat.date}
                  message={chat.message}
                  selected={activeChat == chat.user_id}
                />
              </div>
            )
          }
        })}
      </div>
      {activeChat == 0 ? (
        <div className='hidden size-full min-h-[calc(100dvh-60px-236px)] flex-col place-items-center justify-center gap-2 md:flex'>
          <Icon icon='ep:chat-dot-round' className='size-20' />
          <Text>เลือกแชททางด้านซ้ายมือเพื่อเริ่มบทสนทนา</Text>
        </div>
      ) : (
        <Chat
          sellerName={selectedChat?.sellerName || ''}
          chat={chat}
          input={input}
          setInput={setInput}
          previewImages={previewImages}
          handleImageUpload={handleImageUpload}
          removeImage={removeImage}
          handleBackButton={() => setActiveChat(0)}
          handleSendMessage={handleSendMessage}
        />
      )}
    </div>
  )
}
