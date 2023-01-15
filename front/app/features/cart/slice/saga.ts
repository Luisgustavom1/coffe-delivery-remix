import type { CartProduct } from '@/@types/Api';
import type { PayloadAction } from '@reduxjs/toolkit';
import type { AxiosError } from 'axios';
import { call, put, takeLatest } from 'redux-saga/effects'
import api from '@/services/api';
import { handleError } from '@/utils/handlers'
import { CartActions } from './actions';
import { toast } from 'react-toastify';
import { cartActions } from '.';

function* setCartProduct() {
   try {
      yield put(cartActions.calculateCartTotal())
   } catch (e) {
      const errorMessage: string = yield handleError(e as AxiosError<string>)      
      yield toast.error(errorMessage)
   }
}

function* addCartProduct<T extends CartProduct>(action: PayloadAction<T>) {
   try {
      if (!action.payload.quantity) return;

      yield call((newProduct: T) => {
         return api.post('/cart', {
            quantity: newProduct.quantity,
            productId: newProduct.product.id
         })
      }, action.payload);
      yield put(cartActions.calculateCartTotal())
   } catch (e) {
      const errorMessage: string = yield handleError(e as AxiosError<string>)      
      yield toast.error(errorMessage)
   }
}

function* updateCartProduct<T extends CartProduct>(action: PayloadAction<T>) {
   try {
      yield call((productUpdated: T) => {
         return api.put(`/cart/${productUpdated.id}`, {
            id: productUpdated.id,
            productId: productUpdated.product.id,
            quantity: productUpdated.quantity,
         })
      }, action.payload);
      yield put(cartActions.calculateCartTotal())
   } catch (e) {
      const errorMessage: string = yield handleError(e as AxiosError<string>)      
      yield toast.error(errorMessage)
   }
}

function* deleteCartProduct<T extends number>(action: PayloadAction<T>) {
   try {
      yield call((productId: T) => {
         return api.delete(`/cart/${productId}`)
      }, action.payload);
      yield put(cartActions.calculateCartTotal())
   } catch (e) {
      const errorMessage: string = yield handleError(e as AxiosError<string>)      
      yield toast.error(errorMessage)
   }
}

export function* cartSaga() {
  yield takeLatest(CartActions.SET_CART_PRODUCT, setCartProduct);
  yield takeLatest(CartActions.ADD_CART_PRODUCT, addCartProduct);
  yield takeLatest(CartActions.UPDATE_CART_PRODUCT, updateCartProduct);
  yield takeLatest(CartActions.DELETE_CART_PRODUCT, deleteCartProduct);
}