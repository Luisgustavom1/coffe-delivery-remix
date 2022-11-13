import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'
import type { CartProduct, } from '@/@types/Api';

interface ICartState {
  cart: Array<CartProduct>
  quantityOfProducts: number
  cartTotal: {
    items: number
    freight: number
  }
}

const initialState: ICartState = {
  cart: [],
  quantityOfProducts: 0,
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
      state.quantityOfProducts += 1;
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
      state.quantityOfProducts -= 1;
      return;
    },
    setCart: (state, action: PayloadAction<Array<CartProduct>>) => {
      state.cart = action.payload
    },
    setCartQuantity: (state, action: PayloadAction<number>) => {
      state.quantityOfProducts = action.payload
    }
  },
})

const { actions: cartActions } = cartSlice

export type { ICartState }
export { cartActions, cartSlice }
export default cartSlice.reducer