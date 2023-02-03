import type { PayloadAction } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import type { PaymentType } from "@/@types/Api/Cart";
import type { Address } from "@/features/checkout/schemas";

enum CheckoutActionsEnum {
  CHECKOUT_DATA_PRODUCT = 'cart/setCheckoutData',
}

interface CheckoutState {
  address: Address | null
  paymentType: PaymentType | null
}

const initialState: CheckoutState = {
  address: null,
  paymentType: null
}

const checkoutSlice = createSlice({
  name: 'checkout',
  initialState,
  reducers: {
    setCheckoutData: (state, action: PayloadAction<CheckoutState>) => {      
      console.log(action.payload)
      state.address = action.payload.address
      state.paymentType = action.payload.paymentType
    },
  }
})

const { actions: CheckoutActions, reducer: checkoutReducer } = checkoutSlice

export { CheckoutActionsEnum, CheckoutActions, checkoutSlice, checkoutReducer }
