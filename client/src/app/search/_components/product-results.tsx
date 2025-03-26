import { ProductCard } from '@/components/product-card'

import type { ArtToy } from '@/generated/api'

interface ProductResultsProps {
  products: ArtToy[]
}

export function ProductResults({ products }: ProductResultsProps) {
  return (
    <div className='grid size-full grid-cols-2 place-content-center gap-3 p-3 sm:grid-cols-4'>
      {products.map((product) => (
        <ProductCard key={product.id} product={product} />
      ))}
    </div>
  )
}
