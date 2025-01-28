'use client'

import { Text } from '@/components/text'

const SearchPanel = () => {
  const handleSelectRating = (value: string) => {
    // TODO: Implement select rating
    console.log(value)
  }

  return (
    <aside className='h-full w-72 bg-white'>
      <div className='flex w-full flex-col'>
        <Text>Customer Review</Text>
        <button onClick={() => handleSelectRating('4')}>4</button>
        <button onClick={() => handleSelectRating('3')}>3</button>
        <button onClick={() => handleSelectRating('2')}>2</button>
        <button onClick={() => handleSelectRating('1')}>1</button>
      </div>
      <div className='w-full bg-primary-gradient'>Price Range</div>
    </aside>
  )
}

export { SearchPanel }
