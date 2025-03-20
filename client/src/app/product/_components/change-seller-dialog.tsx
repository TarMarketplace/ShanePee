import { Button } from '@/components/button'
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/dialog'
import { Text } from '@/components/text'

type ChangeSellerDialogProps = {
  sellerNameFrom: string
  sellerNameTo: string
  showDialog: boolean
  setShowDialog: (show: boolean) => void
  handleConfirmDifferentSeller: () => void
}

export function ChangeSellerDialog({
  sellerNameFrom,
  sellerNameTo,
  showDialog,
  setShowDialog,
  handleConfirmDifferentSeller,
}: ChangeSellerDialogProps) {
  return (
    <Dialog open={showDialog} onOpenChange={setShowDialog}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>
            <Text variant='heading-md' className='w-full'>
              ต้องการจะเปลี่ยนร้านหรือไม่
            </Text>
          </DialogTitle>
          <DialogDescription>
            <Text variant='lg-regular' className=''>
              คุณกำลังจะเปลี่ยนร้านจาก
              <span className='mx-1 font-bold'>{sellerNameFrom}</span>
              เป็น
              <span className='mx-1 font-bold'>{sellerNameTo}</span>
              <br />
              หากคุณเปลี่ยนร้าน <br className='md:hidden' />
              <span className='text-primary underline'>
                สินค้าในตะกร้าทั้งหมดของคุณจะหายไป!!
              </span>
              กรุณาตรวจสอบสินค้าในตะกร้าก่อนยืนยัน
            </Text>
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <DialogClose asChild>
            <Button
              variant='filled'
              onClick={() => handleConfirmDifferentSeller()}
            >
              เปลี่ยนร้าน
            </Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
