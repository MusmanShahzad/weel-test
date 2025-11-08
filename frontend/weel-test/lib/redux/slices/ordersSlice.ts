import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import {
  Order,
  AISuggestedProduct,
  GetAISuggestionsRequest,
  CreateOrderRequest,
  OrderStats,
} from "@/types";

interface OrdersState {
  orders: Order[];
  selectedOrder: Order | null;
  aiSuggestions: AISuggestedProduct[];
  stats: OrderStats | null;
  loading: boolean;
  error: string | null;
  filters: {
    status?: string;
    deliveryPreference?: string;
    search?: string;
  };
  sortBy: "created_at" | "status" | "delivery_preference";
  sortOrder: "asc" | "desc";
}

const initialState: OrdersState = {
  orders: [],
  selectedOrder: null,
  aiSuggestions: [],
  stats: null,
  loading: false,
  error: null,
  filters: {},
  sortBy: "created_at",
  sortOrder: "desc",
};

const ordersSlice = createSlice({
  name: "orders",
  initialState,
  reducers: {
    // Fetch Orders
    fetchOrdersRequest: (state) => {
      state.loading = true;
      state.error = null;
    },
    fetchOrdersSuccess: (state, action: PayloadAction<Order[]>) => {
      state.orders = action.payload;
      state.loading = false;
      state.error = null;
    },
    fetchOrdersFailure: (state, action: PayloadAction<string>) => {
      state.loading = false;
      state.error = action.payload;
    },

    // Get AI Suggestions
    getAISuggestionsRequest: (state, action: PayloadAction<GetAISuggestionsRequest>) => {
      state.loading = true;
      state.error = null;
    },
    getAISuggestionsSuccess: (state, action: PayloadAction<AISuggestedProduct[]>) => {
      state.aiSuggestions = action.payload;
      state.loading = false;
      state.error = null;
    },
    getAISuggestionsFailure: (state, action: PayloadAction<string>) => {
      state.loading = false;
      state.error = action.payload;
    },

    // Create Order
    createOrderRequest: (state, action: PayloadAction<CreateOrderRequest>) => {
      state.loading = true;
      state.error = null;
    },
    createOrderSuccess: (state, action: PayloadAction<Order>) => {
      state.orders.unshift(action.payload);
      state.aiSuggestions = [];
      state.loading = false;
      state.error = null;
    },
    createOrderFailure: (state, action: PayloadAction<string>) => {
      state.loading = false;
      state.error = action.payload;
    },

    // Update Order
    updateOrderRequest: (state, action: PayloadAction<{ id: number; status: string }>) => {
      state.loading = true;
      state.error = null;
    },
    updateOrderSuccess: (state, action: PayloadAction<Order>) => {
      const index = state.orders.findIndex((o) => o.id === action.payload.id);
      if (index !== -1) {
        state.orders[index] = action.payload;
      }
      state.loading = false;
      state.error = null;
    },
    updateOrderFailure: (state, action: PayloadAction<string>) => {
      state.loading = false;
      state.error = action.payload;
    },

    // Filters & Sorting
    setFilters: (state, action: PayloadAction<OrdersState["filters"]>) => {
      state.filters = action.payload;
    },
    setSorting: (
      state,
      action: PayloadAction<{ sortBy: OrdersState["sortBy"]; sortOrder: OrdersState["sortOrder"] }>
    ) => {
      state.sortBy = action.payload.sortBy;
      state.sortOrder = action.payload.sortOrder;
    },

    // Calculate Stats
    calculateStats: (state) => {
      const stats: OrderStats = {
        total: state.orders.length,
        pending: 0,
        processing: 0,
        completed: 0,
        cancelled: 0,
        inStore: 0,
        delivery: 0,
        curbside: 0,
      };

      state.orders.forEach((order) => {
        // Status counts
        switch (order.status) {
          case "pending":
            stats.pending++;
            break;
          case "processing":
            stats.processing++;
            break;
          case "completed":
            stats.completed++;
            break;
          case "cancelled":
            stats.cancelled++;
            break;
        }

        // Delivery preference counts
        switch (order.delivery_preference) {
          case "IN_STORE":
            stats.inStore++;
            break;
          case "DELIVERY":
            stats.delivery++;
            break;
          case "CURBSIDE":
            stats.curbside++;
            break;
        }
      });

      state.stats = stats;
    },

    clearAISuggestions: (state) => {
      state.aiSuggestions = [];
    },
  },
});

export const {
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
  setFilters,
  setSorting,
  calculateStats,
  clearAISuggestions,
} = ordersSlice.actions;

export default ordersSlice.reducer;

