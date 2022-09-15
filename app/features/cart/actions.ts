import type { CartProduct } from '@/@types/Api'

export enum CartActions {
   ADD_CART_PRODUCT = 'cart/addCartProduct'
}

export const addCartProduct = (product: CartProduct) => {
   return {
      type: CartActions.ADD_CART_PRODUCT,
      payload: product
   }
}