"use client";

import React, { useState } from "react";
import { Order, AISuggestedProduct } from "@/types";
import { Button } from "../atoms/Button";
import { Badge } from "../atoms/Badge";
import { useAppDispatch } from "@/lib/redux/hooks";
import { updateOrderRequest } from "@/lib/redux/slices/ordersSlice";
import { X } from "lucide-react";

interface ViewOrderDialogProps {
  order: Order;
  onClose: () => void;
}

export const ViewOrderDialog: React.FC<ViewOrderDialogProps> = ({ order, onClose }) => {
  const dispatch = useAppDispatch();
  const [selectedStatus, setSelectedStatus] = useState(order.status);
  const [isEditing, setIsEditing] = useState(false);

  const handleUpdateStatus = () => {
    if (selectedStatus !== order.status) {
      dispatch(updateOrderRequest({ id: order.id, status: selectedStatus }));
    }
    setIsEditing(false);
    onClose();
  };

  const getStatusBadgeVariant = (status: string) => {
    switch (status) {
      case "pending":
        return "warning";
      case "processing":
        return "default";
      case "completed":
        return "success";
      case "cancelled":
        return "destructive";
      default:
        return "outline";
    }
  };

  const getDeliveryBadgeVariant = (preference: string) => {
    switch (preference) {
      case "DELIVERY":
        return "default";
      case "IN_STORE":
        return "secondary";
      case "CURBSIDE":
        return "outline";
      default:
        return "outline";
    }
  };

  let aiProducts: AISuggestedProduct[] = [];
  if (order.ai_suggested_products) {
    try {
      aiProducts = JSON.parse(order.ai_suggested_products);
    } catch (e) {
      console.error("Failed to parse AI products:", e);
    }
  }

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/50">
      <div className="bg-white rounded-lg shadow-lg w-full max-w-3xl max-h-[90vh] overflow-y-auto m-4">
        <div className="p-6">
          {/* Header */}
          <div className="flex items-center justify-between mb-6">
            <div>
              <h2 className="text-2xl font-bold text-gray-900">Order #{order.id}</h2>
              <p className="text-sm text-gray-500">
                Created: {new Date(order.created_at).toLocaleString()}
              </p>
            </div>
            <button
              onClick={onClose}
              className="text-gray-400 hover:text-gray-600 p-2"
            >
              <X className="h-5 w-5" />
            </button>
          </div>

          {/* Order Summary */}
          <div className="space-y-6">
            <div>
              <h3 className="text-sm font-semibold text-gray-700 mb-2">Order Summary</h3>
              <p className="text-gray-900 bg-gray-50 p-4 rounded-md border border-gray-200">
                {order.summary}
              </p>
            </div>

            {/* Status */}
            <div>
              <h3 className="text-sm font-semibold text-gray-700 mb-2">Status</h3>
              {isEditing ? (
                <div className="flex gap-2 items-center">
                  <select
                    value={selectedStatus}
                    onChange={(e) => setSelectedStatus(e.target.value as any)}
                    className="flex h-9 rounded-md border border-gray-300 bg-white px-3 py-1 text-sm text-gray-900"
                  >
                    <option value="pending">Pending</option>
                    <option value="processing">Processing</option>
                    <option value="completed">Completed</option>
                    <option value="cancelled">Cancelled</option>
                  </select>
                  <Button size="sm" onClick={handleUpdateStatus}>
                    Save
                  </Button>
                  <Button size="sm" variant="outline" onClick={() => setIsEditing(false)}>
                    Cancel
                  </Button>
                </div>
              ) : (
                <div className="flex gap-2 items-center">
                  <Badge variant={getStatusBadgeVariant(order.status)}>
                    {order.status.toUpperCase()}
                  </Badge>
                  <Button size="sm" variant="outline" onClick={() => setIsEditing(true)}>
                    Edit Status
                  </Button>
                </div>
              )}
            </div>

            {/* Delivery Details */}
            <div className="grid grid-cols-2 gap-4">
              <div>
                <h3 className="text-sm font-semibold text-gray-700 mb-2">Delivery Type</h3>
                <Badge variant={getDeliveryBadgeVariant(order.delivery_preference)}>
                  {order.delivery_preference}
                </Badge>
              </div>
              {order.delivery_address && (
                <div>
                  <h3 className="text-sm font-semibold text-gray-700 mb-2">Delivery Address</h3>
                  <p className="text-sm text-gray-900">{order.delivery_address}</p>
                  {order.postal_code && (
                    <p className="text-xs text-gray-500 mt-1">Postal Code: {order.postal_code}</p>
                  )}
                </div>
              )}
            </div>

            {/* AI Suggested Products */}
            {aiProducts.length > 0 && (
              <div>
                <h3 className="text-sm font-semibold text-gray-700 mb-3">Selected Products</h3>
                <div className="space-y-2">
                  {aiProducts.map((product, index) => (
                    <div
                      key={index}
                      className="p-4 border border-gray-200 rounded-lg bg-white"
                    >
                      <div className="flex justify-between items-start">
                        <div className="flex-1">
                          <h4 className="font-medium text-gray-900">{product.name}</h4>
                          {product.reason && (
                            <p className="text-sm text-gray-600 mt-1">{product.reason}</p>
                          )}
                        </div>
                        <div className="text-right ml-4">
                          <p className="font-semibold text-gray-900">${product.price.toFixed(2)}</p>
                          <p className="text-sm text-gray-500">Qty: {product.quantity}</p>
                        </div>
                      </div>
                    </div>
                  ))}
                  <div className="pt-2 border-t border-gray-200 flex justify-between">
                    <span className="font-semibold text-gray-900">Total:</span>
                    <span className="font-bold text-gray-900">
                      ${aiProducts.reduce((sum, p) => sum + p.price * p.quantity, 0).toFixed(2)}
                    </span>
                  </div>
                </div>
              </div>
            )}
          </div>

          {/* Footer */}
          <div className="flex gap-2 pt-6 mt-6 border-t border-gray-200">
            <Button onClick={onClose} variant="outline" className="flex-1">
              Close
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
};

