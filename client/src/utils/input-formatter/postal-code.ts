export function formatPostalCode(value: string) {
  const cleaned = ('' + value).replace(/\D/g, '')

  if (cleaned.length > 5) {
    return cleaned.slice(0, 5)
  } else {
    return cleaned
  }
}
