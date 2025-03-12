import { searchArtToys } from '@/generated/api'

import { ProductResults } from './_components/product-results'
import { SearchHeading } from './_components/search-heading'

const SORT_OPTIONS_MAP: Record<
  string,
  { sort_key: 'release_date' | 'price'; reverse: boolean }
> = {
  latest: { sort_key: 'release_date', reverse: true },
  oldest: { sort_key: 'release_date', reverse: false },
  highest_price: { sort_key: 'price', reverse: true },
  lowest_price: { sort_key: 'price', reverse: false },
}

export default async function Search({
  searchParams,
}: {
  searchParams: {
    query?: string
    sort?: string
  }
}) {
  const sortOptions = SORT_OPTIONS_MAP[searchParams.sort ?? 'latest']

  const { data } = await searchArtToys({
    query: {
      keyword: searchParams.query,
      sort_key: sortOptions?.sort_key,
      reverse: sortOptions?.reverse,
    },
    cache: 'no-cache',
  })

  const products = data?.data ?? []

  return (
    <main className='flex size-full flex-col items-center justify-center p-4'>
      <div className='flex w-full flex-col sm:w-fit sm:min-w-[60%] sm:divide-y-4 sm:divide-primary'>
        <SearchHeading query={searchParams.query} />
        <ProductResults products={products} />
      </div>
    </main>
  )
}
