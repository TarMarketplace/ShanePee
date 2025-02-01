function formatPhoneNumber(value: string) {
  const cleaned = ('' + value).replace(/\D/g, '')

  if (cleaned.length <= 3) {
    return cleaned // Up to 3 digits, no dash yet
  } else if (cleaned.length <= 6) {
    return `${cleaned.slice(0, 3)}-${cleaned.slice(3)}` // 123-456
  } else {
    return `${cleaned.slice(0, 3)}-${cleaned.slice(3, 6)}-${cleaned.slice(6, 10)}` // 123-456-7890
  }
}

export { formatPhoneNumber }
