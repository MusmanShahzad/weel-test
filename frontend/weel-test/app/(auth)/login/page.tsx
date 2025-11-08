"use client";

import React from "react";
import { Formik, Form } from "formik";
import * as Yup from "yup";
import Link from "next/link";
import { Button } from "@/components/atoms/Button";
import { FormField } from "@/components/molecules/FormField";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/atoms/Card";
import { useAppDispatch, useAppSelector } from "@/lib/redux/hooks";
import { loginRequest } from "@/lib/redux/slices/authSlice";

const validationSchema = Yup.object({
  email: Yup.string().email("Invalid email address").required("Email is required"),
  password: Yup.string().min(6, "Password must be at least 6 characters").required("Password is required"),
});

export default function LoginPage() {
  const dispatch = useAppDispatch();
  const { loading, error } = useAppSelector((state) => state.auth);

  const handleSubmit = async (
    values: { email: string; password: string },
    { setSubmitting, setErrors }: any
  ) => {
    try {
      dispatch(loginRequest(values));
    } catch (err: any) {
      console.error("Submit error:", err);
      setErrors({ submit: "An unexpected error occurred" });
      setSubmitting(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100 p-4">
      <Card className="w-full max-w-md">
        <CardHeader>
          <CardTitle className="text-2xl">Welcome Back</CardTitle>
          <CardDescription>Sign in to your pharmacy account</CardDescription>
        </CardHeader>
        <CardContent>
          <Formik
            initialValues={{ email: "", password: "" }}
            validationSchema={validationSchema}
            onSubmit={handleSubmit}
            validateOnChange={true}
            validateOnBlur={true}
          >
            {({ isValid, isSubmitting, handleSubmit: formikHandleSubmit }) => (
              <Form className="space-y-4" onSubmit={(e) => {
                e.preventDefault();
                e.stopPropagation();
                formikHandleSubmit(e);
              }}>
                {error && (
                  <div className="p-3 bg-red-50 border border-red-200 rounded-md text-red-800 text-sm">
                    <strong>Error:</strong> {error}
                  </div>
                )}

                <FormField
                  name="email"
                  label="Email"
                  type="email"
                  placeholder="admin@example.com"
                />

                <FormField
                  name="password"
                  label="Password"
                  type="password"
                  placeholder="••••••••"
                />

                <Button
                  type="submit"
                  disabled={!isValid || loading || isSubmitting}
                  className="w-full"
                >
                  {loading || isSubmitting ? "Signing in..." : "Sign In"}
                </Button>

                <div className="text-center text-sm text-muted-foreground">
                  Don't have an account?{" "}
                  <Link href="/signup" className="text-primary hover:underline font-medium">
                    Sign up
                  </Link>
                </div>

                <div className="mt-4 p-3 bg-blue-50 border border-blue-200 rounded-md text-blue-800 text-xs">
                  <strong>Test Account:</strong><br />
                  Email: admin@example.com<br />
                  Password: password123
                </div>
              </Form>
            )}
          </Formik>
        </CardContent>
      </Card>
    </div>
  );
}
