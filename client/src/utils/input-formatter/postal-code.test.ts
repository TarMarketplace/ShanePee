import { describe, expect, it } from 'vitest'

import { formatPostalCode } from './postal-code'

describe('formatPostalCode', () => {
  it('should format a postal code with non-digit characters', () => {
    expect(formatPostalCode('123-45')).toBe('12345')
  })

  it('should format a postal code with more than 5 digits', () => {
    expect(formatPostalCode('123456789')).toBe('12345')
  })

  it('should return the same postal code if it has 5 digits', () => {
    expect(formatPostalCode('12345')).toBe('12345')
  })

  it('should return the same postal code if it has less than 5 digits', () => {
    expect(formatPostalCode('123')).toBe('123')
  })

  it('should return an empty string if the input is empty', () => {
    expect(formatPostalCode('')).toBe('')
  })

  it('should return an empty string if the input has no digits', () => {
    expect(formatPostalCode('abcde')).toBe('')
  })
})
