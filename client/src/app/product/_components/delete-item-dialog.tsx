import { Button } from '@/components/button'
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/dialog'
import { Text } from '@/components/text'

type ProductFormProps = {
  productName: string
  handleDeleteProduct: () => void
}

export function DeleteItemDialog({
  productName,
  handleDeleteProduct,
}: ProductFormProps) {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant='filled' className='w-14' type='button'>
          ลบ
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>
            <Text variant='heading-md' className='w-full'>
              ต้องการลบสินค้าหรือไม่
            </Text>
          </DialogTitle>
          <DialogDescription>
            <Text variant='lg-regular' className='text-black'>
              คุณกำลังจะลบสินค้า
              <span className='mx-1 font-bold'>
                {productName || 'ไม่มีชื่อสินค้า'}
              </span>
              <br />
              <span className='text-primary underline'>
                สินค้าของคุณจะหายไปและไม่สามารถกู้คืนได้!!{' '}
              </span>
              <br />
              กรุณาตรวจสอบก่อนยืนยัน
            </Text>
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <DialogClose asChild>
            <Button variant='filled' onClick={() => handleDeleteProduct()}>
              ลบสินค้า
            </Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
