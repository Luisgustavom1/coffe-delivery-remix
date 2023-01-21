import type { CartProduct } from '@/@types/Api'

export enum CartActions {
   SET_CART_PRODUCT = 'cart/setCartProduct',
   ADD_CART_PRODUCT = 'cart/addCartProduct',
   UPDATE_CART_PRODUCT = 'cart/updateCartProduct',
   DELETE_CART_PRODUCT = 'cart/deleteCartProduct',
   CLEAR_CART_PRODUCT = 'cart/clearCartProduct',
   CALCULATE_CART_TOTAL = 'cart/calculateCartTotal'
}

type AddCartProductParams = Omit<CartProduct, 'id'>

export const setCartProduct = (cartProducts: Array<CartProduct>) => {
   return {
      type: CartActions.SET_CART_PRODUCT,
      payload: cartProducts
   }
}

export const addCartProduct = (product: AddCartProductParams) => {
   return {
      type: CartActions.ADD_CART_PRODUCT,
      payload: product
   }
}

export const updateCartProduct = (productUpdated: CartProduct) => {
   return {
      type: CartActions.UPDATE_CART_PRODUCT,
      payload: productUpdated
   }
}

export const clearCartProduct = () => {
   return {
      type: CartActions.CLEAR_CART_PRODUCT,
      payload: [[]]
   }
}

export const deleteCartProduct = (cartId: number) => {
   return {
      type: CartActions.DELETE_CART_PRODUCT,
      payload: cartId
   }
}

export const calculateCartTotal = () => {
   return {
      type: CartActions.DELETE_CART_PRODUCT,
      payload: {}
   }
}
