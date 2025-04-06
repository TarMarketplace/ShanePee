'use client'

import { Icon } from '@iconify/react/dist/iconify.js'
import { useRouter } from 'next/navigation'
import { useCallback, useEffect, useRef, useState } from 'react'
import { toast } from 'sonner'

import { Text } from '@/components/text'

import { useUser } from '@/providers/user-provider'

import { cn } from '@/lib/utils'

import type { ChatMessage ,
  ChatList} from '@/generated/api'
import {
  getChatList,
  getChatMessage,
  sendMessage,
} from '@/generated/api'

import { Chat } from '../_components/chat'
import { ChatBox } from '../_components/chat-box'

export function ChatListContainer() {
  const [chatList, setChatList] = useState<ChatList[]>([])
  const [activeChat, setActiveChat] = useState(0)
  const [chat, setChat] = useState<ChatMessage[]>([])
  const [input, setInput] = useState('')
  const pollChatIdRef = useRef(0)
  const isPollingRef = useRef(false)
  const pollChatListIdRef = useRef(0)
  const isPollingChatListRef = useRef(false)
  const { user } = useUser()
  const router = useRouter()

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
          toast.error('Something went wrong')
        }
      })
      .then(() => {
        pollChatList()
      })
      .catch(() => {
        toast.error('Something went wrong')
      })
  }, [router])

  const getUserChatList = useCallback(async () => {
    isPollingChatListRef.current = true

    await getChatList({
      query: {
        poll: false,
      },
    })
      .then((response) => {
        const list = response?.data?.data
        if (Array.isArray(list)) {
          if (list.length < 1) return
          setChatList(list)
          pollChatListIdRef.current = list[list.length - 1].id
        } else if (response.response.status == 401) {
          isPollingChatListRef.current = false
          router.push('/login')
        } else {
          toast.error('Something went wrong')
        }
      })
      .then(() => {
        pollChatList()
      })
      .catch((e) => {
        console.log(e)
        toast.error('Something went wrong')
      })
  }, [pollChatList, router])

  useEffect(() => {
    getUserChatList()
  }, [getUserChatList])

  const handleSendMessage = (message_type: 'MESSAGE' | 'IMAGE') => {
    if (!input) {
      toast.warning('Please enter a message')
      return
    }

    sendMessage({
      path: {
        userID: activeChat,
      },
      body: {
        content: input,
        message_type: message_type,
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
          setChat(message)
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

  const selectedChat = chatList.find((chat) => chat?.target_id == activeChat)

  return (
    <div className='flex min-h-[calc(100dvh-60px-236px)] divide-x-2 divide-grey-200'>
      <div
        className={cn(
          'w-full flex-col md:min-w-[350px] md:max-w-[350px]',
          activeChat == 0 ? 'flex' : 'hidden md:flex'
        )}
      >
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
          sellerName={
            (selectedChat?.target_first_name ?? '') +
            ' ' +
            (selectedChat?.target_last_name ?? '')
          }
          handleBackButton={() => setActiveChat(0)}
          chat={chat}
          input={input}
          setInput={setInput}
          handleSendMessage={handleSendMessage}
        />
      )}
    </div>
  )
}
