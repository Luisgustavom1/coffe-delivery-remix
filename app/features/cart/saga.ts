import type { CartProduct } from '@/@types/Api';
import api from '@/services/api';
import type { PayloadAction } from '@reduxjs/toolkit';
import type { AxiosError, AxiosResponse } from 'axios';
import { call, put, takeLatest } from 'redux-saga/effects'
import { CartActions } from './actions';
import { cartActions } from './slice'

function* addCartProduct<T extends CartProduct>(action: PayloadAction<T>) {
   try {
      const { data }: AxiosResponse<T>= yield call((newProduct: T) => {
         return api.put('/cart', {
            product: newProduct
         })
      }, action.payload);

      yield put(cartActions.addCartProduct(data));
   } catch (e) {
      yield put({type: "ADD_CART_PRODUCT_FAILED", message: (e as AxiosError).response?.data});
   }
}

export function* cartSaga() {
  yield takeLatest(CartActions.ADD_CART_PRODUCT, addCartProduct);
}