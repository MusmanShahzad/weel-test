"use client";

import React, { useEffect, useState } from "react";
import { useAppDispatch, useAppSelector } from "@/lib/redux/hooks";
import { fetchOrdersRequest } from "@/lib/redux/slices/ordersSlice";
import { logout } from "@/lib/redux/slices/authSlice";
import { Button } from "@/components/atoms/Button";
import { StatCard } from "@/components/molecules/StatCard";
import { OrdersTable } from "@/components/organisms/OrdersTable";
import { OrderFilters } from "@/components/organisms/OrderFilters";
import { CreateOrderDialog } from "@/components/organisms/CreateOrderDialog";
import { ViewOrderDialog } from "@/components/organisms/ViewOrderDialog";
import {
  ShoppingCart,
  Package,
  Truck,
  Store,
  Clock,
  CheckCircle,
  XCircle,
  Loader,
} from "lucide-react";

export default function DashboardPage() {
  const dispatch = useAppDispatch();
  const { user } = useAppSelector((state) => state.auth);
  const { orders, stats, filters, sortBy, sortOrder, loading } = useAppSelector(
    (state) => state.orders
  );
  const [showCreateDialog, setShowCreateDialog] = useState(false);
  const [selectedOrder, setSelectedOrder] = useState<any>(null);

  useEffect(() => {
    dispatch(fetchOrdersRequest());
  }, [dispatch]);

  const handleLogout = () => {
    dispatch(logout());
  };

  const handleViewOrder = (order: any) => {
    setSelectedOrder(order);
  };

  const handleUpdateStatus = (orderId: number, status: string) => {
  };

  const filteredOrders = orders.filter((order) => {
    if (filters.status && order.status !== filters.status) return false;
    if (
      filters.deliveryPreference &&
      order.delivery_preference !== filters.deliveryPreference
    )
      return false;
    if (filters.search) {
      const searchLower = filters.search.toLowerCase();
      return (
        order.summary.toLowerCase().includes(searchLower) ||
        order.id.toString().includes(searchLower)
      );
    }
    return true;
  });

  const sortedOrders = [...filteredOrders].sort((a, b) => {
    const modifier = sortOrder === "asc" ? 1 : -1;
    if (sortBy === "created_at") {
      return (
        (new Date(a.created_at).getTime() - new Date(b.created_at).getTime()) * modifier
      );
    }
    if (sortBy === "status") {
      return a.status.localeCompare(b.status) * modifier;
    }
    if (sortBy === "delivery_preference") {
      return a.delivery_preference.localeCompare(b.delivery_preference) * modifier;
    }
    return 0;
  });

  return (
    <div className="min-h-screen bg-gray-50">
      <header className="bg-white border-b">
        <div className="container mx-auto px-4 py-4 flex items-center justify-between">
          <div>
            <h1 className="text-2xl font-bold">Pharmacy Orders</h1>
            <p className="text-sm text-muted-foreground">
              Welcome back, {user?.first_name || "User"}
            </p>
          </div>
          <div className="flex gap-2">
            <Button onClick={() => setShowCreateDialog(true)}>+ New Order</Button>
            <Button variant="outline" onClick={handleLogout}>
              Logout
            </Button>
          </div>
        </div>
      </header>

      <main className="container mx-auto px-4 py-8">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
          <StatCard
            title="Total Orders"
            value={stats?.total || 0}
            icon={ShoppingCart}
            description="All time orders"
          />
          <StatCard
            title="Pending"
            value={stats?.pending || 0}
            icon={Clock}
            description="Awaiting processing"
          />
          <StatCard
            title="Completed"
            value={stats?.completed || 0}
            icon={CheckCircle}
            description="Successfully fulfilled"
          />
          <StatCard
            title="Cancelled"
            value={stats?.cancelled || 0}
            icon={XCircle}
            description="Cancelled orders"
          />
        </div>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-8">
          <StatCard
            title="Delivery"
            value={stats?.delivery || 0}
            icon={Truck}
            description="Home delivery"
          />
          <StatCard
            title="In-Store"
            value={stats?.inStore || 0}
            icon={Store}
            description="Store pickup"
          />
          <StatCard
            title="Curbside"
            value={stats?.curbside || 0}
            icon={Package}
            description="Curbside pickup"
          />
        </div>

        <div className="bg-white rounded-lg shadow">
          <div className="p-6">
            <h2 className="text-xl font-semibold mb-4">Your Orders</h2>
            
            <OrderFilters />

            <div className="mt-4">
              {loading ? (
                <div className="flex items-center justify-center py-12">
                  <Loader className="h-8 w-8 animate-spin text-primary" />
                  <span className="ml-2">Loading orders...</span>
                </div>
              ) : (
                <OrdersTable
                  orders={sortedOrders}
                  onViewOrder={handleViewOrder}
                  onUpdateStatus={handleUpdateStatus}
                />
              )}
            </div>
          </div>
        </div>
      </main>

      {showCreateDialog && (
        <CreateOrderDialog onClose={() => setShowCreateDialog(false)} />
      )}

      {selectedOrder && (
        <ViewOrderDialog order={selectedOrder} onClose={() => setSelectedOrder(null)} />
      )}
    </div>
  );
}
