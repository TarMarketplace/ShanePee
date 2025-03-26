import Link from 'next/link'

import { Text } from '@/components/text'

import { SellerAllReviewContainer } from '../../_containers/seller-all-review-container'

export default function SellerReviewPage({
  params,
}: {
  params: { id: string }
}) {
  return (
    <main className='grid size-full grid-cols-1 place-items-center p-4 md:px-12'>
      <div className='w-full place-items-start pt-6 underline'>
        <Link href={`/seller/${params.id}`}>
          <Text variant='sm-regular'>{'< ย้อนกลับ'}</Text>
        </Link>
      </div>
      <SellerAllReviewContainer sellerId={params.id} />
    </main>
  )
}
