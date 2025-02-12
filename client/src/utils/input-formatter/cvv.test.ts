import { describe, expect, it } from 'vitest'

import { formatCVV } from './cvv'

describe('formatCVV', () => {
  it('should return the input if it is already a valid CVV', () => {
    expect(formatCVV('123')).toBe('123')
  })

  it('should remove non-digit characters', () => {
    expect(formatCVV('12a3')).toBe('123')
    expect(formatCVV('1!2@3')).toBe('123')
  })

  it('should truncate the input if it is longer than 3 digits', () => {
    expect(formatCVV('1234')).toBe('123')
    expect(formatCVV('12345')).toBe('123')
  })

  it('should return the input if it is less than 3 digits', () => {
    expect(formatCVV('1')).toBe('1')
    expect(formatCVV('12')).toBe('12')
  })

  it('should return an empty string if the input is empty', () => {
    expect(formatCVV('')).toBe('')
  })
})
