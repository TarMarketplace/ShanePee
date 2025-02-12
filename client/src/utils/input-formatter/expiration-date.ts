export function formatExpirationDate(value: string) {
  const cleaned = ('' + value).replace(/\D/g, '')

  if (cleaned.length <= 2) {
    return cleaned
  } else if (cleaned.length <= 4) {
    return `${cleaned.slice(0, 2)}/${cleaned.slice(2)}`
  } else {
    return `${cleaned.slice(0, 2)}/${cleaned.slice(2, 4)}`
  }
}
