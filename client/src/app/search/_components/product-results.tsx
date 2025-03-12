import { ProductCard } from '@/components/product-card'

import type { ArtToy } from '@/generated/api'

interface ProductResultsProps {
  products: ArtToy[]
}

export function ProductResults({ products }: ProductResultsProps) {
  return (
    <div className='grid grid-cols-[repeat(auto-fit,minmax(170px,max-content))] place-content-center gap-3 p-3 sm:grid-cols-[repeat(auto-fit,minmax(250px,max-content))]'>
      {products.map((product) => (
        <ProductCard key={product.id} product={product} />
      ))}
    </div>
  )
}
