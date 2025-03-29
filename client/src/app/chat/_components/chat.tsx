import { Icon } from '@iconify/react/dist/iconify.js'
import Link from 'next/link'

import { Button } from '@/components/button'
import { Input } from '@/components/input'
import { Text } from '@/components/text'

import type { ChatMessage } from '@/generated/api'

export interface ChatProps {
  sellerName: string
  handleBackButton: () => void
  chat: ChatMessage[]
  input: string
  setInput: (input: string) => void
  handleSendMessage: () => void
}

function Chat({
  sellerName,
  handleBackButton,
  chat,
  input,
  setInput,
  handleSendMessage,
}: ChatProps) {
  // const chat = [
  // {
  //   receive: true,
  //   message:
  //     'สวัสดีครับรบกวนสอบถามว่ามีถั่วมั้ย ไม่ได้หมายถึงถั่ว เป็นเม็ดนะครับ แต่หมายถึงน้องอาเนียจาก SpyXFamily',
  //   dateTime: new Date(),
  // },
  // {
  //   receive: false,
  //   message:
  //     'สวัสดีครับ ติดต่อจากร้าน POP MART นะครับ แอดมินเจมส์พร้อมให้บริการครับ ทางร้านเรามีแต่ตุ๊กตา CRYBABY x Powerpuff Girls Series Figures พอจะสนใจไหมครับ',
  //   dateTime: new Date(),
  // },
  // { receive: true, message: 'อ๋อ งั้นไม่เป็นไรครับ', dateTime: new Date() },
  // ]

  return (
    <div className='flex min-h-[calc(100dvh-60px-236px)] w-full flex-col divide-y-2 divide-grey-200 truncate'>
      <span className='flex h-12 items-center p-4'>
        <Icon
          icon='weui:back-filled'
          className='mr-2 size-6 cursor-pointer md:hidden'
          onClick={() => handleBackButton()}
        />
        <Text variant='lg-semibold' className='w-full'>
          {sellerName}
        </Text>
        <Link href={'/seller'}>
          <Button variant='filled' className='h-8'>
            <Icon icon='tdesign:store-filled' className='size-5' />
            <Text variant='md-regular'>หน้าร้านค้า</Text>
          </Button>
        </Link>
      </span>
      <div className='h-0 grow overflow-y-auto'>
        {chat?.map((message) => {
          if (message) {
            if (message.sender == 'SELLER') {
              return (
                <div className='flex w-full justify-start p-4' key={message.id}>
                  <div className='relative mb-3 max-w-[60%] text-wrap rounded-lg bg-secondary-100 p-2 shadow'>
                    {message.content}
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
            } else {
              return (
                <div className='flex w-full justify-end p-4' key={message.id}>
                  <div className='relative mb-3 max-w-[60%] text-wrap rounded-lg bg-primary-500 p-2 text-white shadow'>
                    {message.content}
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
            }
          }
        })}
      </div>

      <span className='flex h-12 items-center justify-between p-2'>
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
        <Icon
          icon='ic:baseline-send'
          className='size-8 cursor-pointer'
          onClick={handleSendMessage}
        />
      </span>
    </div>
  )
}

export { Chat }
