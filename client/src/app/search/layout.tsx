import { SearchPanel } from '@/components/search-panel'

export default function SearchLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <div className='flex size-full'>
      <SearchPanel />
      {children}
    </div>
  )
}
