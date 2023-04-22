import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'
import type { Cart, CartProduct } from '@/@types/Api/Cart';

enum CartActionsEnum {
  SET_CART_PRODUCT = 'cart/setCartProduct',
  ADD_CART_PRODUCT = 'cart/addCartProduct',
  UPDATE_CART_PRODUCT = 'cart/updateCartProduct',
  DELETE_CART_PRODUCT = 'cart/deleteCartProduct',
  CLEAR_CART_PRODUCT = 'cart/clearCartProduct',
  CALCULATE_CART_TOTAL = 'cart/calculateCartTotal'
}

interface CartState {
  cart: Cart
  cartTotal: {
    items: number
    freight: number
  },
}

const initialState: CartState = {
  // FIX THIS, TO NOT USE Type Assertions
  cart: {
    id: -1,
    products: []
  } as Cart,
  cartTotal: {
    items: 0,
    freight: 0
  },
}

const cartSlice = createSlice({
  name: 'cart',
  initialState,
  reducers: {
    addCartProduct: (state, action: PayloadAction<CartProduct>) => {      
      const { product: currentProduct } = action.payload

      const productAlreadyExistsInCart = state.cart.products.find(
        ({ product }) => product.id === currentProduct.id
      )

      if (productAlreadyExistsInCart) return

      state.cart.products = state.cart.products.concat(action.payload)
    },
    updateCartProduct: (state, action: PayloadAction<CartProduct>) => {      
      const { product } = action.payload

      state.cart.products = state.cart.products.map(
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
      
      state.cart.products = state.cart.products.filter(({ product: { id } }) => id !== cartId)
      return;
    },
    clearCart: (state) => {
      // FIX THIS, TO NOT USE Type Assertions
      state.cart = {} as Cart
      state.cartTotal = {
        freight: 0,
        items: 0
      }
      return
    },
    setCartProduct: (state, action: PayloadAction<CartState['cart']>) => {
      state.cart = action.payload
    },
    calculateCartTotal: (state) => {
      const cart = state.cart
      const totalItems = cart.products.reduce((acc, currentProduct) => acc + (currentProduct.product.price * currentProduct.quantity), 0)

      state.cartTotal.items = totalItems;
      state.cartTotal.freight = 350 * cart.products.length;
    }
  },
})

const { actions: CartActions, reducer: cartReducer } = cartSlice

export { CartActionsEnum, CartActions, cartSlice, cartReducer }
