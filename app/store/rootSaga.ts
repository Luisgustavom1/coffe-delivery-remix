import { cartSaga } from "@/features/cart/saga";
import { all } from "redux-saga/effects";

export default function* rootSaga() {
  yield all([
    cartSaga()
  ])
}