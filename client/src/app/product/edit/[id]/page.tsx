import { getArtToyById } from '@/generated/api'

import { ProductFormContainer } from '../../_containers/product-form-container'

const getProduct = async (id: string) => {
  const { response, data } = await getArtToyById({
    path: {
      id: parseInt(id),
    },
  })

  if (response.ok) {
    return data
  } else {
    return null
  }
}

export default async function EditProductPage({
  params,
}: {
  params: { id: string }
}) {
  const product = await getProduct(params.id)

  if (!product) return <p>Product not found.</p>

  return (
    <ProductFormContainer
      defaultValues={{
        name: product.name,
        description: product.description,
        price: product.price,
      }}
      defaultImage={product.photo}
      id={parseInt(params.id)}
    />
  )
}
