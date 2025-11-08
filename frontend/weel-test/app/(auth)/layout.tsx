"use client";

import React from "react";
import { AuthGuard } from "@/lib/auth/AuthGuard";

export default function AuthLayout({ children }: { children: React.ReactNode }) {
  return <AuthGuard requireAuth={false}>{children}</AuthGuard>;
}
