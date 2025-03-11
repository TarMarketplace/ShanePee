import { Text } from '@/components/text'

import { SearchMenu } from './search-menu'

interface SearchHeadingProps {
  query?: string
}

export function SearchHeading({ query }: SearchHeadingProps) {
  return (
    <>
      <div className='flex items-center justify-end py-2 sm:justify-between'>
        <Text
          className='hidden sm:block'
          variant='heading-md'
          desktopVariant='heading-lg'
        >
          ผลการค้นหา: {query}
        </Text>
        <div className='flex items-center gap-3'>
          <Text variant='sm-regular' desktopVariant='md-regular'>
            เรียงลำดับตาม
          </Text>
          <SearchMenu />
        </div>
      </div>
    </>
  )
}
