import { call, put, takeLatest } from "redux-saga/effects";
import { PayloadAction } from "@reduxjs/toolkit";
import { authAPI } from "@/lib/api/services";
import { apiClient } from "@/lib/api/client";
import {
  loginRequest,
  loginSuccess,
  loginFailure,
  signupRequest,
  signupSuccess,
  signupFailure,
  logout,
} from "../slices/authSlice";
import { LoginRequest, SignupRequest, LoginResponse, User } from "@/types";

function* handleLogin(action: PayloadAction<LoginRequest>) {
  try {
    const response: LoginResponse = yield call(authAPI.login, action.payload);
    
    // Save to localStorage
    if (typeof window !== "undefined") {
      localStorage.setItem("token", response.token);
      localStorage.setItem("user", JSON.stringify(response.user));
    }
    
    // Set token in API client
    apiClient.setToken(response.token);
    
    yield put(loginSuccess({ user: response.user, token: response.token }));
    
    // Redirect to dashboard using Next.js router (avoid hard refresh)
    if (typeof window !== "undefined") {
      // Use pushState to avoid page reload
      window.history.pushState({}, '', '/dashboard');
      window.location.href = "/dashboard";
    }
  } catch (error: any) {
    console.error("Login error:", error);
    const errorMessage = error.response?.data?.error || error.message || "Login failed. Please try again.";
    yield put(loginFailure(errorMessage));
  }
}

function* handleSignup(action: PayloadAction<SignupRequest>) {
  try {
    const user: User = yield call(authAPI.signup, action.payload);
    yield put(signupSuccess());
    
    // After successful signup, redirect to login
    if (typeof window !== "undefined") {
      setTimeout(() => {
        window.location.href = "/login";
      }, 1000);
    }
  } catch (error: any) {
    console.error("Signup error:", error);
    const errorMessage = error.response?.data?.error || error.message || "Signup failed. Please try again.";
    yield put(signupFailure(errorMessage));
  }
}

function* handleLogout() {
  // Clear localStorage
  localStorage.removeItem("token");
  localStorage.removeItem("user");
  
  // Redirect to login
  if (typeof window !== "undefined") {
    window.location.href = "/login";
  }
}

export default function* authSaga() {
  yield takeLatest(loginRequest.type, handleLogin);
  yield takeLatest(signupRequest.type, handleSignup);
  yield takeLatest(logout.type, handleLogout);
}

