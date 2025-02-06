import Image from 'next/image'

function AuthBanner() {
  return (
    <div className='hidden w-full flex-col items-center justify-center p-4 md:flex'>
      <Image src='/logo.png' alt='logo' width={360} height={134} />
      <h1 className='text-center text-xl font-bold lg:text-3xl'>
        ตลาด Art Toys มือสองอันดับหนึ่งในประเทศไทย
      </h1>
    </div>
  )
}

export { AuthBanner }
