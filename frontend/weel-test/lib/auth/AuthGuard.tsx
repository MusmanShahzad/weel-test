"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { useAppSelector } from "@/lib/redux/hooks";

interface AuthGuardProps {
  children: React.ReactNode;
  requireAuth?: boolean;
}

export const AuthGuard: React.FC<AuthGuardProps> = ({ children, requireAuth = true }) => {
  const router = useRouter();
  const { token } = useAppSelector((state) => state.auth);
  const [isChecking, setIsChecking] = useState(true);

  useEffect(() => {
    if (typeof window === "undefined") return;

    const storageToken = localStorage.getItem("token");
    const hasToken = token || storageToken;

    if (requireAuth && !hasToken) {
      router.replace("/login");
    } else if (!requireAuth && hasToken) {
      router.replace("/dashboard");
    } else {
      setIsChecking(false);
    }
  }, [token, requireAuth, router]);

  if (isChecking) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-gray-900"></div>
      </div>
    );
  }

  return <>{children}</>;
};

