import { ProductContainer } from '../_containers/product-container'
import { SellerReviewContainer } from '../_containers/seller-review-container'
import { SellerStatsContainer } from '../_containers/seller-stats-container'

export default function SellerPage({ params }: { params: { id: string } }) {
  return (
    <main className='grid size-full grid-cols-1 place-items-center p-4 md:px-12'>
      <SellerStatsContainer sellerId={params.id} />
      <SellerReviewContainer sellerId={params.id} />
      <ProductContainer sellerId={params.id} />
    </main>
  )
}
