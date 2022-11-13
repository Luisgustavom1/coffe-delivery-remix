import type { CartProduct } from '@/@types/Api'

export enum CartActions {
   ADD_CART_PRODUCT = 'cart/addCartProduct',
   UPDATE_CART_PRODUCT = 'cart/updateCartProduct',
   DELETE_CART_PRODUCT = 'cart/deleteCartProduct'
}

type AddCartProductParams = Omit<CartProduct, 'id'>

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

export const deleteCartProduct = (cartId: number) => {
   return {
      type: CartActions.DELETE_CART_PRODUCT,
      payload: cartId
   }
}