import { call, put, takeLatest } from "redux-saga/effects";
import { PayloadAction } from "@reduxjs/toolkit";
import { ordersAPI } from "@/lib/api/services";
import {
  fetchOrdersRequest,
  fetchOrdersSuccess,
  fetchOrdersFailure,
  getAISuggestionsRequest,
  getAISuggestionsSuccess,
  getAISuggestionsFailure,
  createOrderRequest,
  createOrderSuccess,
  createOrderFailure,
  updateOrderRequest,
  updateOrderSuccess,
  updateOrderFailure,
  calculateStats,
} from "../slices/ordersSlice";
import {
  Order,
  GetAISuggestionsRequest,
  GetAISuggestionsResponse,
  CreateOrderRequest,
} from "@/types";

function* handleFetchOrders() {
  try {
    const response: { orders: Order[]; count: number } = yield call(ordersAPI.getOrders);
    yield put(fetchOrdersSuccess(response.orders));
    yield put(calculateStats());
  } catch (error: any) {
    yield put(fetchOrdersFailure(error.response?.data?.error || "Failed to fetch orders"));
  }
}

function* handleGetAISuggestions(action: PayloadAction<GetAISuggestionsRequest>) {
  try {
    const response: GetAISuggestionsResponse = yield call(
      ordersAPI.getAISuggestions,
      action.payload
    );
    yield put(getAISuggestionsSuccess(response.suggestions));
  } catch (error: any) {
    yield put(getAISuggestionsFailure(error.response?.data?.error || "Failed to get AI suggestions"));
  }
}

function* handleCreateOrder(action: PayloadAction<CreateOrderRequest>) {
  try {
    const order: Order = yield call(ordersAPI.createOrder, action.payload);
    yield put(createOrderSuccess(order));
    yield put(calculateStats());
  } catch (error: any) {
    yield put(createOrderFailure(error.response?.data?.error || "Failed to create order"));
  }
}

function* handleUpdateOrder(action: PayloadAction<{ id: number; status: string }>) {
  try {
    const order: Order = yield call(ordersAPI.updateOrder, action.payload.id, {
      status: action.payload.status as any,
    });
    yield put(updateOrderSuccess(order));
    yield put(calculateStats());
  } catch (error: any) {
    yield put(updateOrderFailure(error.response?.data?.error || "Failed to update order"));
  }
}

export default function* ordersSaga() {
  yield takeLatest(fetchOrdersRequest.type, handleFetchOrders);
  yield takeLatest(getAISuggestionsRequest.type, handleGetAISuggestions);
  yield takeLatest(createOrderRequest.type, handleCreateOrder);
  yield takeLatest(updateOrderRequest.type, handleUpdateOrder);
}

