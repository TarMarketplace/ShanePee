'use client'

import { createContext, useContext, useEffect, useState } from 'react'

import type { CartItem} from '@/generated/api';
import { getCart } from '@/generated/api'

export interface CartData {
  cartItems: CartItem[]
  setCartItems: (cartItems: CartItem[]) => void
  fetchCart: () => void
}

const CartContext = createContext<CartData | undefined>(undefined)

function CartProvider({ children }: { children?: React.ReactNode }) {
  const [cartItems, setCartItems] = useState<CartItem[]>([])

  const fetchCart = async () => {
    try {
      const { response, data } = await getCart()
      if (!data || !Array.isArray(data.data)) {
        return
      }
      if (response.ok) {
        setCartItems(data.data)
      }
    } catch {
      setCartItems([])
    }
  }

  useEffect(() => {
    fetchCart()
  }, [])

  return (
    <CartContext.Provider
      value={{
        cartItems,
        setCartItems,
        fetchCart,
      }}
    >
      {children}
    </CartContext.Provider>
  )
}

function useCart(): CartData {
  const context = useContext(CartContext)
  if (context === undefined) {
    throw new Error('useCart must be used within CartProvider')
  }
  return context
}

export { CartContext, CartProvider, useCart }
