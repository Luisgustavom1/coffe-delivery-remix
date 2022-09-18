import type { CartProduct } from '@/@types/Api';
import type { PayloadAction } from '@reduxjs/toolkit';
import type { AxiosError } from 'axios';
import { call, takeLatest } from 'redux-saga/effects'
import api from '@/services/api';
import { handleError } from '@/utils/handlers'
import { CartActions } from './actions';
import { toast } from 'react-toastify';

function* addCartProduct<T extends CartProduct>(action: PayloadAction<T>) {
   try {
      if (!action.payload.quantity) return;

      yield call((newProduct: T) => {
         return api.put('/cart', newProduct)
      }, action.payload);
   } catch (e) {
      const errorMessage: string = yield handleError(e as AxiosError<string>)      
      yield toast.error(errorMessage)
   }
}

export function* cartSaga() {
  yield takeLatest(CartActions.ADD_CART_PRODUCT, addCartProduct);
}