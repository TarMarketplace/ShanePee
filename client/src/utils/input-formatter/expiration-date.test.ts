import { describe, expect, it } from 'vitest'

import { formatExpirationDate } from './expiration-date'

describe('formatExpirationDate', () => {
  it('should format a single digit correctly', () => {
    expect(formatExpirationDate('1')).toBe('1')
  })

  it('should format two digits correctly', () => {
    expect(formatExpirationDate('12')).toBe('12')
  })

  it('should format three digits correctly', () => {
    expect(formatExpirationDate('123')).toBe('12/3')
  })

  it('should format four digits correctly', () => {
    expect(formatExpirationDate('1234')).toBe('12/34')
  })

  it('should format more than four digits correctly', () => {
    expect(formatExpirationDate('12345')).toBe('12/34')
  })

  it('should remove non-digit characters', () => {
    expect(formatExpirationDate('12a34')).toBe('12/34')
    expect(formatExpirationDate('1b2c')).toBe('12')
  })
})
