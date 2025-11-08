import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { User, LoginRequest, SignupRequest } from "@/types";

interface AuthState {
  user: User | null;
  token: string | null;
  loading: boolean;
  error: string | null;
}

const initialState: AuthState = {
  user: null,
  token: null,
  loading: false,
  error: null,
};

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    // Login
    loginRequest: (state, action: PayloadAction<LoginRequest>) => {
      state.loading = true;
      state.error = null; // Clear previous errors
    },
    loginSuccess: (state, action: PayloadAction<{ user: User; token: string }>) => {
      state.user = action.payload.user;
      state.token = action.payload.token;
      state.loading = false;
      state.error = null;
    },
    loginFailure: (state, action: PayloadAction<string>) => {
      state.loading = false;
      state.error = action.payload;
      state.user = null;
      state.token = null;
    },

    // Signup
    signupRequest: (state, action: PayloadAction<SignupRequest>) => {
      state.loading = true;
      state.error = null;
    },
    signupSuccess: (state) => {
      state.loading = false;
      state.error = null;
    },
    signupFailure: (state, action: PayloadAction<string>) => {
      state.loading = false;
      state.error = action.payload;
    },

    // Logout
    logout: (state) => {
      state.user = null;
      state.token = null;
      state.loading = false;
      state.error = null;
    },

    // Initialize auth from storage
    initializeAuth: (state, action: PayloadAction<{ user: User; token: string } | null>) => {
      if (action.payload) {
        state.user = action.payload.user;
        state.token = action.payload.token;
      }
    },
  },
});

export const {
  loginRequest,
  loginSuccess,
  loginFailure,
  signupRequest,
  signupSuccess,
  signupFailure,
  logout,
  initializeAuth,
} = authSlice.actions;

export default authSlice.reducer;

