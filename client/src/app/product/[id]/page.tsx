import { getArtToyById } from '@/generated/api'

import { ProductPageContainer } from '../_containers/product-page-container'

const getProduct = async (id: string) => {
  const { data, response } = await getArtToyById({
    path: {
      id: parseInt(id),
    },
    cache: 'no-cache',
  })

  if (response.ok) {
    return data
  } else {
    return null
  }
}

export default async function ProductPage({
  params,
}: {
  params: { id: string }
}) {
  const product = await getProduct(params.id)
  return <ProductPageContainer product={product} />
}
