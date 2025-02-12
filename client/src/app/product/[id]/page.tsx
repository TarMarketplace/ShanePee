import { env } from 'process'

const getProduct = async (id: string) => {
  const response = await fetch(`${env.NEXT_PUBLIC_BASE_API_URL}/arttoy/${id}`)

  if (response.ok) {
    const data = await response.json()
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

  if (!product) return <p>Product not found.</p>

  return <div>{product.name}</div>
}
