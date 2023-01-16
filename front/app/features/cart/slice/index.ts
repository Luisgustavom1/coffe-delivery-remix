import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'
import type { CartProduct, } from '@/@types/Api';

interface ICartState {
  cart: Array<CartProduct>
  cartTotal: {
    items: number
    freight: number
  }
}

const initialState: ICartState = {
  cart: [],
  cartTotal: {
    items: 0,
    freight: 0
  }
}

const cartSlice = createSlice({
  name: 'cart',
  initialState,
  reducers: {
    addCartProduct: (state, action: PayloadAction<CartProduct>) => {      
      const { product: currentProduct } = action.payload

      const productAlreadyExistsInCart = state.cart.find(
        ({product}) => product.id === currentProduct.id
      )

      if (productAlreadyExistsInCart) return

      state.cart = state.cart.concat(action.payload)
    },
    updateCartProduct: (state, action: PayloadAction<CartProduct>) => {      
      const { product } = action.payload

      state.cart = state.cart.map(
        (c) => 
          c.product.id === product.id 
            ? {
              ...action.payload
            }
            : c
          )
      return;
    },
    deleteCartProduct: (state, action: PayloadAction<number>) => {      
      const cartId = action.payload
      
      state.cart = state.cart.filter(({ id }) => id !== cartId)
      return;
    },
    clearCartProduct: (state) => {
      state.cart = []
      state.cartTotal = {
        freight: 0,
        items: 0
      }
      return
    },
    setCartProduct: (state, action: PayloadAction<Array<CartProduct>>) => {
      state.cart = action.payload
    },
    calculateCartTotal: (state) => {
      const cart = state.cart
      const totalItems = cart.reduce((acc, currentProduct) => acc + (currentProduct.product.price * currentProduct.quantity), 0)

      state.cartTotal.items = totalItems;
      state.cartTotal.freight = 350 * cart.length;
    }
  },
})

const { actions: cartActions } = cartSlice

export type { ICartState }
export { cartActions, cartSlice }
export default cartSlice.reducer