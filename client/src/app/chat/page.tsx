import { Suspense } from 'react'

import { ChatListContainer } from './_containers/chat-list-container'

export default function Chat() {
  return (
    <main className='size-full'>
      <Suspense fallback={<div>Loading...</div>}>
        <ChatListContainer />
      </Suspense>
    </main>
  )
}
