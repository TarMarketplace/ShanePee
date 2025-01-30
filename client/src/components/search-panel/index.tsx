'use client'

import { Text } from '@/components/text'

const SearchPanel = () => {
  return (
    <aside className='h-full w-72 bg-white'>
      <div className='flex w-full flex-col'>
        <Text>Customer Review</Text>
        <button onClick={() => console.log('4')}>4</button>
        <button onClick={() => console.log('3')}>3</button>
        <button onClick={() => console.log('2')}>2</button>
        <button onClick={() => console.log('1')}>1</button>
      </div>
      <div className='w-full bg-primary-gradient'>Price Range</div>
    </aside>
  )
}

export { SearchPanel }
