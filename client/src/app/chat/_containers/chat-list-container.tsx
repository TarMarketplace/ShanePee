'use client'

import { Icon } from '@iconify/react/dist/iconify.js'
import imageCompression from 'browser-image-compression'
import { useRouter, useSearchParams } from 'next/navigation'
import { useCallback, useEffect, useRef, useState } from 'react'
import { toast } from 'sonner'

import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import { cn } from '@/lib/utils'

import type { ChatList, ChatMessage } from '@/generated/api'
import {
  getChatList,
  getChatMessage,
  getSellerById,
  sendMessage,
} from '@/generated/api'

import { Chat } from '../_components/chat'
import { ChatBox } from '../_components/chat-box'

export function ChatListContainer() {
  const [newChat, setNewChat] = useState<ChatList>()
  const [chatList, setChatList] = useState<ChatList[]>([])
  const [activeChat, setActiveChat] = useState(0)
  const [chat, setChat] = useState<ChatMessage[]>([])
  const [input, setInput] = useState('')
  const [previewImages, setPreviewImages] = useState<string[]>([])
  const pollChatIdRef = useRef(0)
  const isPollingRef = useRef(false)
  const pollChatListIdRef = useRef(0)
  const isPollingChatListRef = useRef(false)
  const { user } = useUser()
  const router = useRouter()
  const searchParams = useSearchParams()

  const pollChatList = useCallback(async () => {
    if (!isPollingChatListRef.current) return

    await getChatList({
      query: {
        poll: true,
        chatID: pollChatListIdRef.current,
      },
    })
      .then((response) => {
        const list = response?.data?.data
        if (Array.isArray(list)) {
          setChatList((prevChatList) => [...prevChatList, ...list])
          pollChatListIdRef.current = list[list.length - 1].id
        } else if (response.response.status == 401) {
          isPollingChatListRef.current = false
          router.push('/login')
        } else {
          toast.error('Something went wrong (PollChatList)')
        }
      })
      .then(() => {
        pollChatList()
      })
      .catch(() => {
        toast.error('Something went wrong (PollChatList)')
      })
  }, [router])

  function timeoutPromise<T>(
    promise: Promise<T>,
    timeoutMs: number
  ): Promise<T> {
    return Promise.race([
      promise,
      new Promise<never>((_, reject) =>
        setTimeout(() => reject(new Error('getChatList timeout')), timeoutMs)
      ),
    ])
  }

  const getUserChatList = useCallback(async () => {
    timeoutPromise(getChatList({ query: { poll: false } }), 500)
      .then((response) => {
        const list = response?.data?.data
        if (Array.isArray(list)) {
          if (list.length < 1) return
          isPollingChatListRef.current = true
          setChatList(list)
          pollChatListIdRef.current = list[list.length - 1].id

          const personId = searchParams.get('id')
          if (personId) {
            if (list.find((chat) => chat?.target_id == Number(personId))) {
              setActiveChat(Number(personId))
            } else {
              fetchNewChat(Number(personId))
            }
          }
        } else if (response.response.status == 401) {
          isPollingChatListRef.current = false
          router.push('/login')
        } else {
          toast.error('Something went wrong (ChatList)')
        }
      })
      .then(() => {
        pollChatList()
      })
      .catch((e) => {
        console.log(e)
        const personId = searchParams.get('id')
        if (personId) {
          fetchNewChat(Number(personId))
        } else {
          toast.error('Something went wrong (ChatList)')
        }
      })
  }, [searchParams, pollChatList, router])

  useEffect(() => {
    getUserChatList()
  }, [getUserChatList])

  const fetchNewChat = async (newPersonId: number) => {
    const response = await getSellerById({ path: { id: newPersonId } })
    const newPerson = response?.data
    const newChatListItem: ChatList = {
      id: 0,
      last_chat_content: '',
      last_chat_message_type: 'MESSAGE',
      last_chat_time: Date.now().toString(),
      target_first_name: newPerson?.first_name || '',
      target_id: newPerson?.id ?? 0,
      target_last_name: newPerson?.last_name || '',
      target_photo: newPerson?.photo || '',
    }
    setActiveChat(newPersonId)
    setNewChat(newChatListItem)
  }

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

  const handleSendMessage = async () => {
    if (previewImages.length == 0) {
      if (!input) {
        toast.warning('Please enter a message')
        return
      }

      await sendMessage({
        path: {
          userID: activeChat,
        },
        body: {
          content: input,
          message_type: 'MESSAGE',
        },
      })
        .then((response) => {
          if (response.data) {
            setInput('')
            pollChatIdRef.current = response.data.id
            toast.success('Message sent')
          } else {
            toast.error('Error sending message')
          }
        })
        .catch(() => {
          toast.error('Error sending message')
        })
    } else {
      const promises = previewImages.map(
        async (img) =>
          await sendMessage({
            path: {
              userID: activeChat,
            },
            body: {
              content: img,
              message_type: 'IMAGE',
            },
          })
      )

      Promise.all(promises)
        .then((responses) => {
          const hasError = responses.some((res) => !res?.data)
          if (hasError) {
            toast.error('Some images failed to send')
          } else {
            toast.success('Image(s) sent')
            setPreviewImages([])
          }
        })
        .catch(() => {
          toast.error('Error sending images')
        })
    }
    if (activeChat == newChat?.target_id) {
      window.location.reload()
    }
  }

  const pollChat = useCallback(async () => {
    if (!isPollingRef.current) return

    await getChatMessage({
      path: {
        userID: activeChat,
      },
      query: {
        poll: true,
        chatID: pollChatIdRef.current,
      },
    })
      .then((response) => {
        const message = response?.data?.data
        if (Array.isArray(message)) {
          if (message.length < 1) return
          setChat((prevChat) => [...prevChat, ...message])
          pollChatIdRef.current = message[message.length - 1].id
        } else if (response.response.status == 401) {
          isPollingRef.current = false
          router.push('/login')
        } else {
          toast.error('Something went wrong (PollChat)')
        }
      })
      .then(() => {
        pollChat()
      })
      .catch(() => {
        toast.error('Something went wrong (PollChat)')
      })
  }, [activeChat, router])

  const getChat = useCallback(async () => {
    if (!chatList.find((chat) => chat?.target_id == activeChat)) return

    await getChatMessage({
      path: {
        userID: activeChat,
      },
      query: {
        poll: false,
      },
    })
      .then((response) => {
        const message = response?.data?.data
        if (Array.isArray(message)) {
          if (message.length < 1) return
          isPollingRef.current = true
          setChat(message)
          pollChatIdRef.current = message[message.length - 1].id
        } else if (response.response.status == 401) {
          isPollingRef.current = false
          router.push('/login')
        } else {
          toast.error('Something went wrong (GetChat)')
        }
      })
      .then(() => {
        pollChat()
      })
      .catch((e) => {
        console.log(e)
        toast.error('Something went wrong (GetChat)')
      })
  }, [activeChat, pollChat, router, chatList])

  useEffect(() => {
    if (activeChat) {
      setChat([])
      if (activeChat != user?.id) getChat()
    }
  }, [activeChat, user?.id, getChat])

  const selectedChat = [...chatList, newChat].find(
    (chat) => chat?.target_id == activeChat
  )

  return (
    <div className='flex max-h-[calc(100dvh-62px-256px)] divide-x-2 divide-grey-200'>
      <div
        className={cn(
          'w-full flex-col md:min-w-[350px] md:max-w-[350px]',
          activeChat == 0 ? 'flex' : 'hidden md:flex'
        )}
      >
        {newChat && (
          <div
            key={newChat.id}
            onClick={() =>
              setActiveChat(
                activeChat == newChat.target_id ? 0 : newChat.target_id
              )
            }
            className='cursor-pointer'
          >
            <ChatBox
              photo={newChat.target_photo ?? undefined}
              sellerName={
                newChat.target_first_name + ' ' + newChat.target_last_name
              }
              sellerNameFallback={
                (newChat.target_first_name?.charAt(0).toUpperCase() ?? '') +
                (newChat.target_last_name?.charAt(0).toUpperCase() ?? '')
              }
              date={new Date()}
              message={
                newChat.last_chat_message_type == 'MESSAGE'
                  ? newChat.last_chat_content
                  : 'Image'
              }
              selected={activeChat == newChat.target_id}
            />
          </div>
        )}
        {chatList.map((chat) => {
          if (chat) {
            return (
              <div
                key={chat.id}
                onClick={() =>
                  setActiveChat(
                    activeChat == chat.target_id ? 0 : chat.target_id
                  )
                }
                className='cursor-pointer'
              >
                <ChatBox
                  photo={chat.target_photo ?? undefined}
                  sellerName={
                    chat.target_first_name + ' ' + chat.target_last_name
                  }
                  sellerNameFallback={
                    (chat.target_first_name?.charAt(0).toUpperCase() ?? '') +
                    (chat.target_last_name?.charAt(0).toUpperCase() ?? '')
                  }
                  date={new Date(chat.last_chat_time)}
                  message={
                    chat.last_chat_message_type == 'MESSAGE'
                      ? chat.last_chat_content
                      : 'Image'
                  }
                  selected={activeChat == chat.target_id}
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
          sender_id={user?.id ?? null}
          seller_id={selectedChat?.target_id ?? null}
          sellerName={
            (selectedChat?.target_first_name ?? '') +
            ' ' +
            (selectedChat?.target_last_name ?? '')
          }
          handleBackButton={() => setActiveChat(0)}
          chat={chat}
          input={input}
          setInput={setInput}
          previewImages={previewImages}
          handleImageUpload={handleImageUpload}
          removeImage={removeImage}
          handleSendMessage={handleSendMessage}
        />
      )}
    </div>
  )
}
