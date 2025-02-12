import { env } from 'process'

import { ProductFormContainer } from '../../_containers/product-form-container'

const getProduct = async (id: string) => {
  const response = await fetch(`${env.NEXT_PUBLIC_BASE_API_URL}/arttoy/${id}`)

  if (response.ok) {
    const data = await response.json()
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
    <ProductFormContainer defaultValues={product} id={parseInt(params.id)} />
  )
}
