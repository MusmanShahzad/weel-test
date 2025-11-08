"use client";

import React from "react";
import { Formik, Form } from "formik";
import * as Yup from "yup";
import Link from "next/link";
import { Button } from "@/components/atoms/Button";
import { FormField } from "@/components/molecules/FormField";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/atoms/Card";
import { useAppDispatch, useAppSelector } from "@/lib/redux/hooks";
import { signupRequest } from "@/lib/redux/slices/authSlice";

const validationSchema = Yup.object({
  first_name: Yup.string().required("First name is required"),
  last_name: Yup.string().required("Last name is required"),
  email: Yup.string().email("Invalid email address").required("Email is required"),
  password: Yup.string().min(6, "Password must be at least 6 characters").required("Password is required"),
  confirm_password: Yup.string()
    .oneOf([Yup.ref("password")], "Passwords must match")
    .required("Please confirm your password"),
});

export default function SignupPage() {
  const dispatch = useAppDispatch();
  const { loading, error } = useAppSelector((state) => state.auth);

  const handleSubmit = async (values: any, { setSubmitting, setErrors }: any) => {
    try {
      const { confirm_password, ...signupData } = values;
      dispatch(signupRequest(signupData));
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
          <CardTitle className="text-2xl">Create Account</CardTitle>
          <CardDescription>Sign up for your pharmacy account</CardDescription>
        </CardHeader>
        <CardContent>
          <Formik
            initialValues={{
              first_name: "",
              last_name: "",
              email: "",
              password: "",
              confirm_password: "",
            }}
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

                <div className="grid grid-cols-2 gap-4">
                  <FormField name="first_name" label="First Name" placeholder="John" />
                  <FormField name="last_name" label="Last Name" placeholder="Doe" />
                </div>

                <FormField name="email" label="Email" type="email" placeholder="john@example.com" />

                <FormField name="password" label="Password" type="password" placeholder="••••••••" />

                <FormField
                  name="confirm_password"
                  label="Confirm Password"
                  type="password"
                  placeholder="••••••••"
                />

                <Button type="submit" disabled={!isValid || loading || isSubmitting} className="w-full">
                  {loading || isSubmitting ? "Creating account..." : "Create Account"}
                </Button>

                <div className="text-center text-sm text-muted-foreground">
                  Already have an account?{" "}
                  <Link href="/login" className="text-primary hover:underline font-medium">
                    Sign in
                  </Link>
                </div>
              </Form>
            )}
          </Formik>
        </CardContent>
      </Card>
    </div>
  );
}

