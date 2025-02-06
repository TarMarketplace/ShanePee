import Image from 'next/image'

function AuthPageBanner() {
  return (
    <div className='flex flex-col justify-center p-4'>
      <Image src='/logo.png' alt='logo' width={360} height={134} />
      <h1 className='text-center text-3xl font-bold'>
        ตลาด Art Toys มือสอง
        <br />
        อันดับหนึ่งในประเทศไทย
      </h1>
    </div>
  )
}

export default AuthPageBanner
