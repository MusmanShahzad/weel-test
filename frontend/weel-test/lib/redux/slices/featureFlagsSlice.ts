import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface FeatureFlagsState {
  flags: Record<string, boolean>;
  loading: boolean;
  error: string | null;
}

const initialState: FeatureFlagsState = {
  flags: {},
  loading: false,
  error: null,
};

const featureFlagsSlice = createSlice({
  name: "featureFlags",
  initialState,
  reducers: {
    fetchFeatureFlagsRequest: (state) => {
      state.loading = true;
      state.error = null;
    },
    fetchFeatureFlagsSuccess: (state, action: PayloadAction<Record<string, boolean>>) => {
      state.flags = action.payload;
      state.loading = false;
      state.error = null;
    },
    fetchFeatureFlagsFailure: (state, action: PayloadAction<string>) => {
      state.loading = false;
      state.error = action.payload;
    },
  },
});

export const {
  fetchFeatureFlagsRequest,
  fetchFeatureFlagsSuccess,
  fetchFeatureFlagsFailure,
} = featureFlagsSlice.actions;

export default featureFlagsSlice.reducer;

