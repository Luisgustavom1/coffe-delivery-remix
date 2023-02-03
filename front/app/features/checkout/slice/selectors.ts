import type { RootState } from "@/@types/RootState";
import { createDraftSafeSelector } from "@reduxjs/toolkit";

const selectSelfSlice = (state: RootState) => state.checkout

export const checkoutDataSelector = createDraftSafeSelector(
  selectSelfSlice, 
  (state) => state
)