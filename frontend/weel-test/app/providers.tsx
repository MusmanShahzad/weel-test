"use client";

import React, { useEffect } from "react";
import { Provider } from "react-redux";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { store } from "@/lib/redux/store";
import { initializeAuth } from "@/lib/redux/slices/authSlice";
import { fetchFeatureFlagsRequest } from "@/lib/redux/slices/featureFlagsSlice";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      retry: 1,
    },
  },
});

export function Providers({ children }: { children: React.ReactNode }) {
  useEffect(() => {
    // Initialize auth from localStorage
    if (typeof window !== "undefined") {
      const token = localStorage.getItem("token");
      const userStr = localStorage.getItem("user");
      if (token && userStr) {
        try {
          const user = JSON.parse(userStr);
          store.dispatch(initializeAuth({ user, token }));
        } catch (e) {
          console.error("Failed to parse stored user:", e);
        }
      }
    }

    // Fetch feature flags on app load
    store.dispatch(fetchFeatureFlagsRequest());
  }, []);

  return (
    <Provider store={store}>
      <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
    </Provider>
  );
}

