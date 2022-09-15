import { configureStore } from '@reduxjs/toolkit'
import createSagaMiddleware from 'redux-saga'
import cartReducer from '@/features/cart/slice'
import { cartSaga } from '@/features/cart/saga'

const sagaMiddleware = createSagaMiddleware()

export const store = configureStore({
  reducer: {
    cart: cartReducer,
  },
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(sagaMiddleware)
})

sagaMiddleware.run(cartSaga)

export type AppDispatch = typeof store.dispatch