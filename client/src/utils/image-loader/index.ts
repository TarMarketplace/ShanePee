export async function imageLoader(data: string): Promise<File> {
  const response = await fetch(data)
  const blob = await response.blob()
  return new File([blob], 'image', { type: 'image/png' })
}
