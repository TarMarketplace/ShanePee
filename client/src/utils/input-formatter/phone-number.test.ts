import { describe, expect, it } from 'vitest'

import { formatPhoneNumber } from '.'

describe('phone-formatter', () => {
  it('should format phone number', () => {
    const phoneNumber = '0812345678'
    const expected = '081-234-5678'

    const actual = formatPhoneNumber(phoneNumber)

    expect(actual).toBe(expected)
  })

  it('should format phone number with less than 3 digits', () => {
    const phoneNumber = '08'
    const expected = '08'

    const actual = formatPhoneNumber(phoneNumber)

    expect(actual).toBe(expected)
  })

  it('should format phone number with less than 6 digits', () => {
    const phoneNumber = '0812'
    const expected = '081-2'

    const actual = formatPhoneNumber(phoneNumber)

    expect(actual).toBe(expected)
  })

  it('should format phone number with more than 6 digits', () => {
    const phoneNumber = '081234567'
    const expected = '081-234-567'

    const actual = formatPhoneNumber(phoneNumber)

    expect(actual).toBe(expected)
  })
})
