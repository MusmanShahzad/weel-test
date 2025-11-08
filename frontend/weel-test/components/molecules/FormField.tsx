"use client";

import React from "react";
import { Field, ErrorMessage, FieldProps } from "formik";
import { Input } from "../atoms/Input";
import { Label } from "../atoms/Label";
import { cn } from "@/lib/utils";

interface FormFieldProps {
  name: string;
  label: string;
  type?: string;
  placeholder?: string;
  disabled?: boolean;
  as?: "input" | "textarea" | "select";
  options?: Array<{ value: string; label: string }>;
  className?: string;
}

export const FormField: React.FC<FormFieldProps> = ({
  name,
  label,
  type = "text",
  placeholder,
  disabled = false,
  as = "input",
  options,
  className,
}) => {
  return (
    <div className={cn("space-y-2", className)}>
      <Label htmlFor={name}>{label}</Label>
      {as === "input" && (
        <Field name={name}>
          {({ field, meta }: FieldProps) => (
            <>
              <Input
                {...field}
                id={name}
                type={type}
                placeholder={placeholder}
                disabled={disabled}
                className={meta.touched && meta.error ? "border-red-500" : ""}
              />
            </>
          )}
        </Field>
      )}
      {as === "textarea" && (
        <Field name={name}>
          {({ field, meta }: FieldProps) => (
            <textarea
              {...field}
              id={name}
              placeholder={placeholder}
              disabled={disabled}
              rows={4}
              className={cn(
                "flex min-h-[80px] w-full rounded-md border border-input bg-transparent px-3 py-2 text-base shadow-sm placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
                meta.touched && meta.error ? "border-red-500" : ""
              )}
            />
          )}
        </Field>
      )}
      {as === "select" && options && (
        <Field name={name}>
          {({ field, meta }: FieldProps) => (
            <select
              {...field}
              id={name}
              disabled={disabled}
              className={cn(
                "flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-sm transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
                meta.touched && meta.error ? "border-red-500" : ""
              )}
            >
              <option value="">Select...</option>
              {options.map((option) => (
                <option key={option.value} value={option.value}>
                  {option.label}
                </option>
              ))}
            </select>
          )}
        </Field>
      )}
      <ErrorMessage name={name}>
        {(msg) => <div className="text-xs text-red-500">{msg}</div>}
      </ErrorMessage>
    </div>
  );
};

