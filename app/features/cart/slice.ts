import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'
import type { CartProduct, } from '@/@types/Api';

interface CounterState {
  cart: Array<CartProduct>
  quantityOfProducts: 0
}

const initialState: CounterState = {
  cart: [],
  quantityOfProducts: 0
}

const cartSlice = createSlice({
  name: 'cart',
  initialState,
  reducers: {
    addCartProduct: (state, action: PayloadAction<CartProduct>) => {
      state.cart = state.cart.concat(action.payload)
      state.quantityOfProducts += action.payload.quantity
    },
  },
})

const { actions: cartActions } = cartSlice

export type { CounterState }
export { cartActions, cartSlice }
export default cartSlice.reducer