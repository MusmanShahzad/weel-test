"use client";

import React from "react";
import { Order } from "@/types";
import { Badge } from "../atoms/Badge";
import { Button } from "../atoms/Button";
import { format } from "date-fns";

interface OrdersTableProps {
  orders: Order[];
  onViewOrder: (order: Order) => void;
  onUpdateStatus: (orderId: number, status: string) => void;
}

export const OrdersTable: React.FC<OrdersTableProps> = ({
  orders,
  onViewOrder,
  onUpdateStatus,
}) => {
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

  return (
    <div className="rounded-md border">
      <table className="w-full">
        <thead>
          <tr className="border-b bg-muted/50">
            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground">
              Order ID
            </th>
            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground">
              Summary
            </th>
            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground">
              Delivery
            </th>
            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground">
              Status
            </th>
            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground">
              Date
            </th>
            <th className="h-12 px-4 text-left align-middle font-medium text-muted-foreground">
              Actions
            </th>
          </tr>
        </thead>
        <tbody>
          {orders.map((order) => (
            <tr key={order.id} className="border-b transition-colors hover:bg-muted/50">
              <td className="p-4 align-middle">#{order.id}</td>
              <td className="p-4 align-middle max-w-xs truncate">{order.summary}</td>
              <td className="p-4 align-middle">
                <Badge variant={getDeliveryBadgeVariant(order.delivery_preference)}>
                  {order.delivery_preference}
                </Badge>
              </td>
              <td className="p-4 align-middle">
                <Badge variant={getStatusBadgeVariant(order.status)}>
                  {order.status}
                </Badge>
              </td>
              <td className="p-4 align-middle">
                {format(new Date(order.created_at), "MMM dd, yyyy")}
              </td>
              <td className="p-4 align-middle">
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => onViewOrder(order)}
                >
                  View
                </Button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      {orders.length === 0 && (
        <div className="p-8 text-center text-muted-foreground">
          No orders found. Create your first order!
        </div>
      )}
    </div>
  );
};

