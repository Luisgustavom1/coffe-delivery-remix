import type { RootState } from "@/@types/RootState";
import { createDraftSafeSelector } from "@reduxjs/toolkit";

const selectSelfSlice = (state: RootState) => state.cart

export const cartSelector = createDraftSafeSelector(
  selectSelfSlice,
  (state) => state.cart
)

export const cartTotalSelector = createDraftSafeSelector(
  selectSelfSlice,
  (state) => state.cartTotal
)
