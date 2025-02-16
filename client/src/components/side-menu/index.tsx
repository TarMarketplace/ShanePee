'use client'

import { Icon } from '@iconify/react'
import Link from 'next/link'

interface SidebarProps {
  isOpen: boolean
  setIsOpen: (open: boolean) => void
  onLogout: () => void
}

export function Sidebar({ isOpen, setIsOpen, onLogout }: SidebarProps) {
  const handleBackdropClick = () => {
    setIsOpen(false)
  }

  return (
    <div className='relative'>
      <div className='block md:hidden' onClick={() => setIsOpen(true)}>
        <Icon icon='akar-icons:three-line-horizontal' className='size-8' />
      </div>

      {isOpen && ( // Backdrop
        <div
          className='fixed inset-0 z-40 bg-black/10 md:hidden'
          onClick={handleBackdropClick}
        ></div>
      )}

      <div
        className={`fixed right-0 top-0 z-50 h-full w-64 bg-white shadow-lg transition-transform duration-300 ${isOpen ? 'translate-x-0' : 'translate-x-full'} md:static md:translate-x-0`}
      >
        <Icon
          icon='material-symbols:close-rounded'
          className='absolute right-4 top-4 size-8 text-black md:hidden'
          onClick={() => setIsOpen(false)}
        />

        <div
          className='mt-16 flex flex-col divide-y border-y text-right text-black'
          onClick={() => setIsOpen(false)}
        >
          <div className='p-3'>
            <Link href='/user/account'>บัญชีของฉัน</Link>
          </div>
          <div className='p-3'>การซื้อของฉัน</div>
          <div className='p-3'>ยืนยันตัวตนผู้ขาย</div>
          <div className='p-3'>
            <button onClick={onLogout} className='text-error'>
              ออกจากระบบ
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}
