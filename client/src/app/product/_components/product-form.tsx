// 'use client'

// import { useState } from 'react'
// import { useForm } from 'react-hook-form'
// import { Icon } from '@iconify/react'
// import { Button } from '@/components/button'

// type ProductFormData = {
//   name: string
//   description: string
//   price: number
//   image?: string
// }

// type ProductFormProps = {
//   defaultValues?: ProductFormData
//   onSubmit: (data: ProductFormData) => void
// }

// export default function ProductForm({ defaultValues, onSubmit }: ProductFormProps) {
//   const { register, handleSubmit } = useForm({ defaultValues })
//   const [imagePreview, setImagePreview] = useState(defaultValues?.image || '')

//   const handleImageUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
//     const file = event.target.files?.[0]
//     if (file) {
//       const imageUrl = URL.createObjectURL(file)
//       setImagePreview(imageUrl)
//     }
//   }

//   return (
//     <form
//       onSubmit={handleSubmit(onSubmit)}
//       className="flex flex-col gap-[10px] p-[6px] w-full max-w-[900px] mx-auto"
//     >
//       <div className="flex items-start gap-[2px]">
//       <Icon icon={defaultValues ? "bxs:edit" : "mdi:plus"} width={24} height={24} />
//         <h2 className="text-2xl font-semibold">{defaultValues ? 'แก้ไขสินค้า' : 'วางจำหน่าย Art Toy ใหม่'}</h2>
//       </div>

//       <div className="flex flex-col sm:flex-row gap-4 items-center justify-center w-full px-[5px] sm:px-0">
//         <label className="relative w-full sm:w-[640px] aspect-[1.6/1] border-2 border-gray-300 rounded-lg flex items-center justify-center overflow-hidden cursor-pointer">
//           {imagePreview ? (
//             <img src={imagePreview} alt="Product" className="w-full h-full object-cover rounded-lg" />
//           ) : (
//             <div className="absolute inset-0 flex justify-center items-center transition">
//               <input type="file" className="hidden" accept="image/*" onChange={handleImageUpload} />
//               <Button variant="filled" type="button">
//                 เลือกรูป
//               </Button>
//             </div>
//           )}
//         </label>

//         <div className="flex flex-col gap-2 w-full sm:w-[614px] sm:h-full">
//           <input {...register('name')} type="text" placeholder="ชื่อสินค้า" className="border p-2 rounded-md w-full h-[40px]" required />
//           <textarea {...register('description')} placeholder="รายละเอียด" className="border p-2 rounded-md w-full h-[160px] py-2" required />
//           <input {...register('price')} type="number" placeholder="ราคา" className="border p-2 rounded-md w-full h-[40px]" required />
//         </div>
//       </div>

//       <div className="flex items-center justify-end gap-4 mt-auto">
//         {defaultValues && (
//           <button type="button" className="font-bold underline text-grey-500">
//             ยกเลิก
//           </button>
//         )}
//         <Button variant="filled" type="submit">
//           {defaultValues ? 'แก้ไข' : 'วางจำหน่าย'}
//         </Button>
//       </div>
//     </form>
//   )
// }
