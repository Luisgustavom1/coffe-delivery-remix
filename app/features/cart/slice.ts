import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'
import type { CartProduct, } from '@/@types/Api';

interface ICartState {
  cart: Array<CartProduct>
  quantityOfProducts: 0
}

const initialState: ICartState = {
  cart: [],
  quantityOfProducts: 0
}

const cartSlice = createSlice({
  name: 'cart',
  initialState,
  reducers: {
    addCartProduct: (state, action: PayloadAction<CartProduct>) => {      
      const { product: currentProduct, quantity } = action.payload

      const productAlreadyExistsInCart = state.cart.find(
        ({product}) => product.id === currentProduct.id
      )
      
      if (productAlreadyExistsInCart && !quantity) {
        state.cart = state.cart.filter(({ product }) => product.id !== currentProduct.id)
        state.quantityOfProducts -= 1;
        return;
      }

      if (productAlreadyExistsInCart) {
        state.cart = state.cart.map(
          (c) => 
            c.product.id === currentProduct.id 
              ? {
                ...action.payload
              }
              : c
            )
        return;
      }

      state.cart = state.cart.concat(action.payload)
      state.quantityOfProducts += 1;
    },
    setCart: (state, action: PayloadAction<Array<CartProduct>>) => {
      state.cart = action.payload
    }
  },
})

const { actions: cartActions } = cartSlice

export type { ICartState }
export { cartActions, cartSlice }
export default cartSlice.reducer