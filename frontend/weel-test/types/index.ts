// User Types
export interface User {
  id: number;
  email: string;
  first_name: string;
  last_name: string;
  created_at: string;
  updated_at: string;
}

// Order Types
export type OrderStatus = "pending" | "processing" | "completed" | "cancelled";
export type DeliveryPreference = "IN_STORE" | "DELIVERY" | "CURBSIDE";

export interface AISuggestedProduct {
  name: string;
  quantity: number;
  price: number;
  reason?: string;
}

export interface Order {
  id: number;
  user_id: number;
  summary: string;
  delivery_preference: DeliveryPreference;
  delivery_address?: string;
  postal_code?: string;
  ai_suggested_products?: string; // JSON string
  status: OrderStatus;
  created_at: string;
  updated_at: string;
}

// Feature Flags
export interface FeatureFlag {
  id: number;
  name: string;
  description: string;
  enabled: boolean;
  created_at: string;
  updated_at: string;
}

export interface FeatureFlagsResponse {
  flags: Record<string, boolean>;
  details: FeatureFlag[];
}

// API Request/Response Types
export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}

export interface SignupRequest {
  email: string;
  password: string;
  first_name: string;
  last_name: string;
}

export interface GetAISuggestionsRequest {
  summary: string;
  delivery_address?: string;
}

export interface GetAISuggestionsResponse {
  suggestions: AISuggestedProduct[];
  count: number;
}

export interface CreateOrderRequest {
  summary: string;
  delivery_preference: DeliveryPreference;
  delivery_address?: string;
  postal_code?: string;
  selected_products?: AISuggestedProduct[];
}

export interface UpdateOrderRequest {
  status?: OrderStatus;
}

// Order Statistics
export interface OrderStats {
  total: number;
  pending: number;
  processing: number;
  completed: number;
  cancelled: number;
  inStore: number;
  delivery: number;
  curbside: number;
}
