import { apiClient } from "./client";
import {
  LoginRequest,
  LoginResponse,
  SignupRequest,
  User,
  GetAISuggestionsRequest,
  GetAISuggestionsResponse,
  CreateOrderRequest,
  UpdateOrderRequest,
  Order,
  FeatureFlagsResponse,
} from "@/types";

// Auth API
export const authAPI = {
  login: (data: LoginRequest) =>
    apiClient.post<LoginResponse>("/auth/login", data),

  signup: (data: SignupRequest) =>
    apiClient.post<User>("/users", data),

  getCurrentUser: () =>
    apiClient.get<User>("/me"),
};

// Orders API
export const ordersAPI = {
  getAISuggestions: (data: GetAISuggestionsRequest) =>
    apiClient.post<GetAISuggestionsResponse>("/orders/suggestions", data),

  createOrder: (data: CreateOrderRequest) =>
    apiClient.post<Order>("/orders", data),

  getOrders: (params?: {
    status?: string;
    delivery_preference?: string;
    sort_by?: string;
    sort_order?: string;
    limit?: number;
    offset?: number;
  }) => {
    const queryParams = new URLSearchParams();
    if (params) {
      Object.entries(params).forEach(([key, value]) => {
        if (value !== undefined && value !== null && value !== "") {
          queryParams.append(key, value.toString());
        }
      });
    }
    const queryString = queryParams.toString();
    return apiClient.get<{ orders: Order[]; count: number }>(
      `/orders${queryString ? `?${queryString}` : ""}`
    );
  },

  getOrder: (id: number) =>
    apiClient.get<Order>(`/orders/${id}`),

  updateOrder: (id: number, data: UpdateOrderRequest) =>
    apiClient.put<Order>(`/orders/${id}`, data),
};

// Feature Flags API
export const featureFlagsAPI = {
  getFeatureFlags: () =>
    apiClient.get<FeatureFlagsResponse>("/feature-flags"),
};

