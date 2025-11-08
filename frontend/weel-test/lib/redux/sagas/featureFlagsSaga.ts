import { call, put, takeLatest } from "redux-saga/effects";
import { featureFlagsAPI } from "@/lib/api/services";
import {
  fetchFeatureFlagsRequest,
  fetchFeatureFlagsSuccess,
  fetchFeatureFlagsFailure,
} from "../slices/featureFlagsSlice";
import { FeatureFlagsResponse } from "@/types";

function* handleFetchFeatureFlags() {
  try {
    const response: FeatureFlagsResponse = yield call(featureFlagsAPI.getFeatureFlags);
    yield put(fetchFeatureFlagsSuccess(response.flags));
  } catch (error: any) {
    yield put(fetchFeatureFlagsFailure(error.response?.data?.error || "Failed to fetch feature flags"));
  }
}

export default function* featureFlagsSaga() {
  yield takeLatest(fetchFeatureFlagsRequest.type, handleFetchFeatureFlags);
}

