import { Icon } from '@iconify/react'
import Image from 'next/image'

import { Text } from '@/components/text'

const Footer = () => {
  return (
    <footer className='flex w-full items-center justify-center bg-reverse-primary-gradient'>
      <div className='flex w-full max-w-screen-xl justify-center divide-x bg-reverse-primary-gradient py-4 text-white'>
        <div className='flex flex-col gap-1 px-5 py-2'>
          <Text variant='xl-semibold'>ข้อมูลการติดต่อ</Text>
          <Text variant='sm-regular'>ที่อยู่: 13 ถ.จอมใหญ่</Text>
          <Text variant='sm-regular'>โทร: 0123456789</Text>
          <Text variant='sm-regular'>อีเมล: support@shanepee.com</Text>
          <Text variant='sm-regular'>สื่อโซเชียลมีเดีย</Text>
          <div className='flex gap-2'>
            <Icon icon='mage:facebook-square' className='size-5' />
            <Icon icon='mage:instagram-square' className='size-5' />
            <Icon icon='mage:line' className='size-5' />
            <Icon icon='mage:x-square' className='size-5' />
          </div>
        </div>
        <div className='px-5 py-2'>
          <Text variant='xl-semibold'>ช่องทางการชำระเงิน</Text>
          <Image
            src='https://placehold.co/173x160.png'
            alt=''
            width={173}
            height={160}
          />
        </div>
        <div className='flex flex-col gap-1 px-5 py-2'>
          <Text variant='xl-semibold'>เกี่ยวกับเรา</Text>
          <Text variant='sm-regular'>ShanePee คืออะไร</Text>
          <Text variant='sm-regular'>นโยบายการใช้งาน</Text>
          <Text variant='sm-regular'>ข้อตกลงและเงื่อนไขการใช้งาน</Text>
          <Text variant='sm-regular'>ร่วมงานกับเรา</Text>
          <Text variant='sm-regular'>คำถามที่พบบ่อย</Text>
        </div>
        <div className='flex items-center justify-center px-5 py-2'>
          <Image
            src='https://placehold.co/360x120.png'
            alt=''
            width={360}
            height={120}
          />
        </div>
      </div>
    </footer>
  )
}

export { Footer }
