import { Text } from '../text'

export interface OrderStatusProps {
  status: 'PREPARING' | 'DELIVERING' | 'COMPLETED'
}

const OrderStatus = ({ status }: OrderStatusProps) => {
  switch (status) {
    case 'PREPARING':
      return (
        <div className='flex h-6 w-20 items-center justify-center rounded-xl border border-warning bg-warning-50'>
          <Text variant='sm-semibold' className='text-center text-warning'>
            รอจัดส่ง
          </Text>
        </div>
      )
    case 'DELIVERING':
      return (
        <div className='flex h-6 w-20 items-center justify-center rounded-xl border border-info bg-info-50'>
          <Text variant='sm-semibold' className='text-center text-info'>
            กำลังจัดส่ง
          </Text>
        </div>
      )
    case 'COMPLETED':
      return (
        <div className='flex h-6 w-20 items-center justify-center rounded-xl border border-success bg-success-50'>
          <Text variant='sm-semibold' className='text-center text-success'>
            จัดส่งสำเร็จ
          </Text>
        </div>
      )
  }
}

export { OrderStatus }
