import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'
import type { Product } from '@/@types/Api/Cart';

enum CartActionsEnum {
  SET_CART_PRODUCT = 'cart/setCartProduct',
  ADD_CART_PRODUCT = 'cart/addCartProduct',
  UPDATE_CART_PRODUCT = 'cart/updateCartProduct',
  DELETE_CART_PRODUCT = 'cart/deleteCartProduct',
  CLEAR_CART_PRODUCT = 'cart/clearCartProduct',
  CALCULATE_CART_TOTAL = 'cart/calculateCartTotal'
}

interface CartState {
  cart: Array<Product>
  cartTotal: {
    items: number
    freight: number
  },
}

const initialState: CartState = {
  cart: [],
  cartTotal: {
    items: 0,
    freight: 0
  },
}

const cartSlice = createSlice({
  name: 'cart',
  initialState,
  reducers: {
    addCartProduct: (state, action: PayloadAction<Product>) => {      
      const { product: currentProduct } = action.payload

      const productAlreadyExistsInCart = state.cart.find(
        ({product}) => product.id === currentProduct.id
      )

      if (productAlreadyExistsInCart) return

      state.cart = state.cart.concat(action.payload)
    },
    updateCartProduct: (state, action: PayloadAction<Product>) => {      
      const { product } = action.payload

      state.cart = state.cart.map(
        (c) => 
          c.product.id === product.id 
            ? {
              ...action.payload
            }
            : c
          );
      state.cart = state.cart.map(
        (c) => 
          c.product.id === product.id 
            ? {
              ...action.payload
            }
            : c
          );
      return;
    },
    deleteCartProduct: (state, action: PayloadAction<number>) => {      
      const cartId = action.payload
      
      state.cart = state.cart.filter(({ id }) => id !== cartId)
      return;
    },
    clearCart: (state) => {
      state.cart = []
      state.cartTotal = {
        freight: 0,
        items: 0
      }
      return
    },
    setCartProduct: (state, action: PayloadAction<Array<Product>>) => {
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

const { actions: CartActions, reducer: cartReducer } = cartSlice

export { CartActionsEnum, CartActions, cartSlice, cartReducer }
