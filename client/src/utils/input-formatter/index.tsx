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

function formatPostalCode(value: string) {
  const cleaned = ('' + value).replace(/\D/g, '')

  if (cleaned.length > 5) {
    return cleaned.slice(0, 5)
  } else {
    return cleaned
  }
}

function formatCardNumber(value: string) {
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

function formatExpirationDate(value: string) {
  const cleaned = ('' + value).replace(/\D/g, '')

  if (cleaned.length <= 2) {
    return cleaned
  } else if (cleaned.length <= 4) {
    return `${cleaned.slice(0, 2)}/${cleaned.slice(2)}`
  } else {
    return `${cleaned.slice(0, 2)}/${cleaned.slice(2, 4)}`
  }
}

function formatCVV(value: string) {
  const cleaned = ('' + value).replace(/\D/g, '')

  if (cleaned.length > 3) {
    return cleaned.slice(0, 3)
  } else {
    return cleaned
  }
}

export {
  formatPhoneNumber,
  formatPostalCode,
  formatCardNumber,
  formatExpirationDate,
  formatCVV,
}
