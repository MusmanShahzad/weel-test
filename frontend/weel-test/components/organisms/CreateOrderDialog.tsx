"use client";

import React, { useState } from "react";
import { Formik, Form } from "formik";
import * as Yup from "yup";
import { Button } from "../atoms/Button";
import { FormField } from "../molecules/FormField";
import { useAppDispatch, useAppSelector } from "@/lib/redux/hooks";
import {
  getAISuggestionsRequest,
  createOrderRequest,
  clearAISuggestions,
} from "@/lib/redux/slices/ordersSlice";
import { AISuggestedProduct } from "@/types";

interface CreateOrderDialogProps {
  onClose: () => void;
}

const validationSchema = Yup.object({
  summary: Yup.string()
    .min(10, "Summary must be at least 10 characters")
    .required("Summary is required"),
  delivery_preference: Yup.string()
    .oneOf(["IN_STORE", "DELIVERY", "CURBSIDE"], "Invalid delivery preference")
    .required("Delivery preference is required"),
  delivery_address: Yup.string().when("delivery_preference", {
    is: "DELIVERY",
    then: (schema) => schema.required("Address is required for delivery"),
    otherwise: (schema) => schema.notRequired(),
  }),
  postal_code: Yup.string(),
});

export const CreateOrderDialog: React.FC<CreateOrderDialogProps> = ({ onClose }) => {
  const dispatch = useAppDispatch();
  const { aiSuggestions, loading } = useAppSelector((state) => state.orders);
  const { flags } = useAppSelector((state) => state.featureFlags);
  const [selectedProducts, setSelectedProducts] = useState<AISuggestedProduct[]>([]);
  const [step, setStep] = useState<"form" | "suggestions">("form");
  
  const isAIEnabled = flags.ai_suggestions === true;

  const handleGetSuggestions = (values: any) => {
    dispatch(
      getAISuggestionsRequest({
        summary: values.summary,
        delivery_address: values.delivery_address || undefined,
      })
    );
    setStep("suggestions");
  };

  const handleCreateOrder = (values: any) => {
    dispatch(
      createOrderRequest({
        summary: values.summary,
        delivery_preference: values.delivery_preference,
        delivery_address: values.delivery_address || undefined,
        postal_code: values.postal_code || undefined,
        selected_products: selectedProducts.length > 0 ? selectedProducts : undefined,
      })
    );
    dispatch(clearAISuggestions());
    onClose();
  };

  const toggleProduct = (product: AISuggestedProduct) => {
    setSelectedProducts((prev) => {
      const exists = prev.find((p) => p.name === product.name);
      if (exists) {
        return prev.filter((p) => p.name !== product.name);
      }
      return [...prev, product];
    });
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/50">
      <div className="bg-white rounded-lg shadow-lg w-full max-w-2xl max-h-[90vh] overflow-y-auto m-4">
        <div className="p-6">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-2xl font-bold">
              {step === "form" ? "Create New Order" : "Select Products"}
            </h2>
            <button
              onClick={onClose}
              className="text-gray-400 hover:text-gray-600"
            >
              ×
            </button>
          </div>

          <Formik
            initialValues={{
              summary: "",
              delivery_preference: "",
              delivery_address: "",
              postal_code: "",
            }}
            validationSchema={validationSchema}
            onSubmit={step === "form" ? handleGetSuggestions : handleCreateOrder}
            validateOnChange={true}
            validateOnBlur={true}
          >
            {({ values, isValid, handleSubmit: formikHandleSubmit }) => (
              <Form className="space-y-4" onSubmit={(e) => {
                e.preventDefault();
                formikHandleSubmit(e);
              }}>
                {step === "form" && (
                  <>
                    <FormField
                      name="summary"
                      label="Order Summary"
                      as="textarea"
                      placeholder="I need medicine for headache and fever..."
                    />

                    <FormField
                      name="delivery_preference"
                      label="Delivery Preference"
                      as="select"
                      options={[
                        { value: "IN_STORE", label: "In-Store Pickup" },
                        { value: "DELIVERY", label: "Home Delivery" },
                        { value: "CURBSIDE", label: "Curbside Pickup" },
                      ]}
                    />

                    {(values.delivery_preference === "DELIVERY" ||
                      values.delivery_preference === "CURBSIDE") && (
                      <FormField
                        name="delivery_address"
                        label="Delivery Address"
                        placeholder="123 Main St, New York, NY 10001"
                      />
                    )}

                    <FormField
                      name="postal_code"
                      label="Postal Code (Optional)"
                      placeholder="10001"
                    />

                    <div className="flex gap-2 pt-4">
                      {isAIEnabled ? (
                        <>
                          <Button
                            type="submit"
                            disabled={!isValid || loading}
                            className="flex-1"
                          >
                            {loading ? "Getting Suggestions..." : "Get AI Suggestions"}
                          </Button>
                          <Button
                            type="button"
                            variant="outline"
                            onClick={() => handleCreateOrder(values)}
                            disabled={!isValid}
                          >
                            Skip & Create
                          </Button>
                        </>
                      ) : (
                        <Button
                          type="button"
                          onClick={() => handleCreateOrder(values)}
                          disabled={!isValid}
                          className="flex-1"
                        >
                          Create Order
                        </Button>
                      )}
                    </div>
                    
                    {!isAIEnabled && (
                      <div className="text-xs text-center text-muted-foreground bg-yellow-50 p-2 rounded">
                        AI suggestions are currently disabled
                      </div>
                    )}
                  </>
                )}

                {step === "suggestions" && (
                  <>
                    <div className="space-y-2">
                      <h3 className="font-semibold">AI Suggested Products:</h3>
                      {loading && <p>Loading suggestions...</p>}
                      {!loading && aiSuggestions.length === 0 && (
                        <p className="text-muted-foreground">
                          No pharmacy products found in your request.
                        </p>
                      )}
                      {!loading && aiSuggestions.length > 0 && (
                        <div className="space-y-2">
                          {aiSuggestions.map((product, index) => {
                            const isSelected = selectedProducts.some(
                              (p) => p.name === product.name
                            );
                            return (
                              <div
                                key={index}
                                className={`p-4 border rounded-lg cursor-pointer transition-colors ${
                                  isSelected
                                    ? "border-primary bg-primary/5"
                                    : "border-gray-200 hover:border-gray-300"
                                }`}
                                onClick={() => toggleProduct(product)}
                              >
                                <div className="flex justify-between items-start">
                                  <div className="flex-1">
                                    <h4 className="font-medium">{product.name}</h4>
                                    <p className="text-sm text-muted-foreground">
                                      {product.reason}
                                    </p>
                                  </div>
                                  <div className="text-right ml-4">
                                    <p className="font-semibold">${product.price}</p>
                                    <p className="text-sm text-muted-foreground">
                                      Qty: {product.quantity}
                                    </p>
                                  </div>
                                </div>
                              </div>
                            );
                          })}
                        </div>
                      )}
                    </div>

                    <div className="flex gap-2 pt-4">
                      <Button
                        type="button"
                        variant="outline"
                        onClick={() => setStep("form")}
                      >
                        Back
                      </Button>
                      <Button
                        type="submit"
                        disabled={loading}
                        className="flex-1"
                      >
                        Create Order ({selectedProducts.length} selected)
                      </Button>
                    </div>
                  </>
                )}
              </Form>
            )}
          </Formik>
        </div>
      </div>
    </div>
  );
};

