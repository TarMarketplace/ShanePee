export function formatCardNumber(value: string) {
  const cleaned = ('' + value).replace(/\D/g, '')

  if (cleaned.length <= 4) {
    return cleaned
  } else if (cleaned.length <= 8) {
    return `${cleaned.slice(0, 4)} ${cleaned.slice(4)}`
  } else if (cleaned.length <= 12) {
    return `${cleaned.slice(0, 4)} ${cleaned.slice(4, 8)} ${cleaned.slice(8)}`
  } else if (cleaned.length <= 16) {
    return `${cleaned.slice(0, 4)} ${cleaned.slice(4, 8)} ${cleaned.slice(8, 12)} ${cleaned.slice(12)}`
  } else {
    return `${cleaned.slice(0, 4)} ${cleaned.slice(4, 8)} ${cleaned.slice(8, 12)} ${cleaned.slice(12, 16)}`
  }
}
