import { cartSaga } from "@/features/cart/slice/saga";
import { all } from "redux-saga/effects";

export default function* rootSaga() {
  yield all([
    cartSaga()
  ])
}