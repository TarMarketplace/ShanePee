import { describe, expect, it } from 'vitest'

import { formatCardNumber } from './card-number'

describe('formatCardNumber', () => {
  it('should format card number with less than 4 digits', () => {
    expect(formatCardNumber('123')).toBe('123')
  })

  it('should format card number with 4 to 8 digits', () => {
    expect(formatCardNumber('123456')).toBe('1234 56')
  })

  it('should format card number with 8 to 12 digits', () => {
    expect(formatCardNumber('1234567890')).toBe('1234 5678 90')
  })

  it('should format card number with 12 to 16 digits', () => {
    expect(formatCardNumber('12345678901234')).toBe('1234 5678 9012 34')
  })

  it('should format card number with more than 16 digits', () => {
    expect(formatCardNumber('123456789012345678')).toBe('1234 5678 9012 3456')
  })

  it('should remove non-digit characters', () => {
    expect(formatCardNumber('1234-5678-9012-3456')).toBe('1234 5678 9012 3456')
  })
})
