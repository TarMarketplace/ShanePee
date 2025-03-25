'use client'

import { Icon } from '@iconify/react/dist/iconify.js'
import { useState } from 'react'

import { Text } from '@/components/text'

import { Chat } from '../_components/chat'
import { ChatBox } from '../_components/chat-box'

export function ChatListContainer() {
  const [activeChat, setActiveChat] = useState(0)

  const chats = [
    {
      id: 1,
      photo: 'photo',
      sellerName: 'John Doe',
      date: new Date(),
      message: 'last message',
    },
    {
      id: 2,
      photo: 'photo',
      sellerName: 'John Doe 2',
      date: new Date(),
      message: 'last message',
    },
    {
      id: 3,
      photo: 'photo',
      sellerName: 'John Doe 3',
      date: new Date(),
      message: 'last message',
    },
  ]

  const selectedChat = chats.find((chat) => chat?.id == activeChat)

  return (
    <div className='flex min-h-[calc(100dvh-60px-236px)] divide-x-2 divide-grey-200'>
      <div
        className={
          activeChat == 0
            ? 'flex w-full flex-col divide-y-2 divide-grey-200 md:min-w-[350px] md:max-w-[350px]'
            : 'hidden w-full flex-col divide-y-2 divide-grey-200 md:flex md:min-w-[350px] md:max-w-[350px]'
        }
      >
        {chats?.map((chat) => {
          if (chat) {
            return (
              <div
                key={chat.id}
                onClick={() =>
                  setActiveChat(activeChat == chat.id ? 0 : chat.id)
                }
                className='cursor-pointer'
              >
                <ChatBox
                  photo={chat.photo}
                  sellerName={chat.sellerName}
                  date={chat.date}
                  message={chat.message}
                  selected={activeChat == chat.id}
                />
              </div>
            )
          }
        })}
        <div></div>
      </div>
      {activeChat == 0 ? (
        <div className='hidden size-full min-h-[calc(100dvh-60px-236px)] flex-col place-items-center justify-center gap-2 md:flex'>
          <Icon icon='ep:chat-dot-round' className='size-20' />
          <Text>เลือกแชททางด้านซ้ายมือเพื่อเริ่มบทสนทนา</Text>
        </div>
      ) : (
        <Chat
          sellerName={selectedChat?.sellerName || ''}
          handleBackButton={() => setActiveChat(0)}
        />
      )}
    </div>
  )
}
