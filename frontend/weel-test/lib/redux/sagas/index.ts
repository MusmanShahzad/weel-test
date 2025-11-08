import { all, fork } from "redux-saga/effects";
import authSaga from "./authSaga";
import ordersSaga from "./ordersSaga";
import featureFlagsSaga from "./featureFlagsSaga";

export default function* rootSaga() {
  yield all([
    fork(authSaga),
    fork(ordersSaga),
    fork(featureFlagsSaga),
  ]);
}

