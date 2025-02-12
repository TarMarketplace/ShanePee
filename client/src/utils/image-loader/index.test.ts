import { describe, expect, it, vi } from 'vitest'

import { imageLoader } from './index'

describe('imageLoader', () => {
  it('should fetch the image and return a File object', async () => {
    const mockBlob = new Blob(['image data'], { type: 'image/png' })
    const mockResponse = {
      blob: vi.fn().mockResolvedValue(mockBlob),
    }
    global.fetch = vi.fn().mockResolvedValue(mockResponse)

    const result = await imageLoader('https://example.com/image.png')

    expect(global.fetch).toHaveBeenCalledWith('https://example.com/image.png')
    expect(mockResponse.blob).toHaveBeenCalled()
    expect(result).toBeInstanceOf(File)
    expect(result.type).toBe('image/png')
    expect(result.name).toBe('image')
  })

  it('should throw an error if fetch fails', async () => {
    global.fetch = vi.fn().mockRejectedValue(new Error('Fetch failed'))

    await expect(imageLoader('https://example.com/image.png')).rejects.toThrow(
      'Fetch failed'
    )
  })
})
