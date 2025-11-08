"use client";

import React from "react";
import { Button } from "../atoms/Button";
import { useAppDispatch, useAppSelector } from "@/lib/redux/hooks";
import { setFilters, setSorting } from "@/lib/redux/slices/ordersSlice";

export const OrderFilters: React.FC = () => {
  const dispatch = useAppDispatch();
  const { filters, sortBy, sortOrder } = useAppSelector((state) => state.orders);

  const handleFilterChange = (key: string, value: string) => {
    dispatch(setFilters({ ...filters, [key]: value }));
  };

  const handleSortChange = (field: "created_at" | "status" | "delivery_preference") => {
    if (sortBy === field) {
      dispatch(setSorting({ sortBy: field, sortOrder: sortOrder === "asc" ? "desc" : "asc" }));
    } else {
      dispatch(setSorting({ sortBy: field, sortOrder: "desc" }));
    }
  };

  return (
    <div className="flex flex-wrap gap-4 p-4 bg-muted/50 rounded-lg">
      <div className="flex-1 min-w-[200px]">
        <select
          value={filters.status || ""}
          onChange={(e) => handleFilterChange("status", e.target.value)}
          className="flex h-9 w-full rounded-md border border-input bg-white px-3 py-1 text-sm"
        >
          <option value="">All Statuses</option>
          <option value="pending">Pending</option>
          <option value="processing">Processing</option>
          <option value="completed">Completed</option>
          <option value="cancelled">Cancelled</option>
        </select>
      </div>

      <div className="flex-1 min-w-[200px]">
        <select
          value={filters.deliveryPreference || ""}
          onChange={(e) => handleFilterChange("deliveryPreference", e.target.value)}
          className="flex h-9 w-full rounded-md border border-input bg-white px-3 py-1 text-sm"
        >
          <option value="">All Delivery Types</option>
          <option value="IN_STORE">In-Store</option>
          <option value="DELIVERY">Delivery</option>
          <option value="CURBSIDE">Curbside</option>
        </select>
      </div>

      <div className="flex gap-2">
        <Button
          variant={sortBy === "created_at" ? "default" : "outline"}
          size="sm"
          onClick={() => handleSortChange("created_at")}
        >
          Date {sortBy === "created_at" && (sortOrder === "asc" ? "↑" : "↓")}
        </Button>
        <Button
          variant={sortBy === "status" ? "default" : "outline"}
          size="sm"
          onClick={() => handleSortChange("status")}
        >
          Status {sortBy === "status" && (sortOrder === "asc" ? "↑" : "↓")}
        </Button>
      </div>

      {(filters.status || filters.deliveryPreference) && (
        <Button
          variant="ghost"
          size="sm"
          onClick={() => dispatch(setFilters({}))}
        >
          Clear Filters
        </Button>
      )}
    </div>
  );
};

