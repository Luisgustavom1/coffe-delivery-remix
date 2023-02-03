import type { FormCheckout } from "../schemas";

export enum CheckoutActions {
  SET_CHECKOUT_DATA = 'cart/setCheckoutData',
}

export const setCheckoutData = (checkoutData: FormCheckout) => {
  return {
    payload: checkoutData,
    type: CheckoutActions.SET_CHECKOUT_DATA
  }
}