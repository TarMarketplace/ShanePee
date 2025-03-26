'use client'

import { Icon } from '@iconify/react'
import { useRouter } from 'next/navigation'

import { Button } from '@/components/button'
import { Text } from '@/components/text'

export default function PaymentSuccess() {
  const router = useRouter()

  const handleGoToOrderHistory = () => {
    router.push('/buyer/order-history')
  }

  const handleGoToHome = () => {
    router.push('/')
  }
  return (
    <main className='flex size-full min-h-[calc(100dvh-60px-236px)] items-center justify-center p-4 md:p-12'>
      <div className='w-full max-w-sm rounded-lg bg-background p-6 text-center shadow-md sm:max-w-md md:max-w-lg'>
        <div className='mb-4 flex items-center justify-center'>
          <Icon
            icon='ix:success'
            className='size-14 text-[#1EB600] md:size-20'
          />
        </div>

        <Text variant='heading-md' desktopVariant='heading-lg' className='mb-2'>
          ชำระเงินสำเร็จ
        </Text>

        <Text variant='sm-regular' desktopVariant='md-regular' className='mb-6'>
          <span className='sm:hidden'>
            ขอบคุณสำหรับการชำระเงิน
            <br />
            คุณสามารถตรวจสอบการสั่งซื้อ
            <br />
            ได้ที่หน้าประวัติการสั่งซื้อ
          </span>

          <span className='hidden sm:inline'>
            ขอบคุณสำหรับการชำระเงิน คุณสามารถ
            <br />
            ตรวจสอบการสั่งซื้อได้ที่หน้าประวัติการสั่งซื้อ
          </span>
        </Text>

        <div className='mt-6 flex items-center justify-center gap-6'>
          <Button variant='filled' onClick={handleGoToHome}>
            กลับหน้าแรก
          </Button>
          <Button variant='filled' onClick={handleGoToOrderHistory}>
            ประวัติการสั่งซื้อ
          </Button>
        </div>
      </div>
    </main>
  )
}
