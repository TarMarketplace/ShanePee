export function formatCVV(value: string) {
  const cleaned = ('' + value).replace(/\D/g, '')

  if (cleaned.length > 3) {
    return cleaned.slice(0, 3)
  } else {
    return cleaned
  }
}
