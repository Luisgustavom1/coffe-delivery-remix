import type { CartProduct } from '@/@types/Api/Cart';
import type { PayloadAction } from '@reduxjs/toolkit';
import type { AxiosError } from 'axios';
import { call, put, takeLatest, select } from 'redux-saga/effects'
import api from '@/services/api';
import { handleError } from '@/utils/handlers'
import { toast } from 'react-toastify';
import { CartActions, CartActionsEnum } from '.';
import { cartSelector } from './selectors';

type CartSelector = ReturnType<typeof cartSelector>

function* setCartProduct() {
   try {
      yield put(CartActions.calculateCartTotal())
   } catch (e) {
      const errorMessage: string = yield handleError(e as AxiosError<string>)      
      yield toast.error(errorMessage)
   }
}

function* createCart<T extends CartProduct>(action: PayloadAction<T>) {
   try {
      if (!action.payload.quantity) return;

      yield call((newProduct: T) => {
         return api.post('/cart', {
            quantity: newProduct.quantity,
            productId: newProduct.product.id
         })
      }, action.payload);
      yield put(CartActions.calculateCartTotal())
   } catch (e) {
      const errorMessage: string = yield handleError(e as AxiosError<string>)      
      yield toast.error(errorMessage)
   }
}

function* updateCartProduct<T extends CartProduct>(action: PayloadAction<T>) {
   const cart: CartSelector = yield select(cartSelector)
   try {
      yield call(({ quantity, product: productUpdated }: T) => {
         return api.put(`/cart/${cart.id}`, {
            productId: productUpdated.id,
            quantity: quantity,
         })
      }, action.payload);
      yield put(CartActions.calculateCartTotal())
   } catch (e) {
      const errorMessage: string = yield handleError(e as AxiosError<string>)      
      yield toast.error(errorMessage)
   }
}

function* deleteCartProduct<T extends number>(action: PayloadAction<T>) {
   const cart: CartSelector = yield select(cartSelector)

   try {
      yield call((productId: T) => {
         return api.delete(`/cart/${cart.id}/product/${productId}`)
      }, action.payload);
      yield put(CartActions.calculateCartTotal())
   } catch (e) {
      const errorMessage: string = yield handleError(e as AxiosError<string>)      
      yield toast.error(errorMessage)
   }
}

export function* cartSaga() {
  yield takeLatest(CartActionsEnum.SET_CART_PRODUCT, setCartProduct);
  yield takeLatest(CartActionsEnum.ADD_CART_PRODUCT, createCart);
  yield takeLatest(CartActionsEnum.UPDATE_CART_PRODUCT, updateCartProduct);
  yield takeLatest(CartActionsEnum.DELETE_CART_PRODUCT, deleteCartProduct);
}