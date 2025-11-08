import React from 'react'
import { render, RenderOptions } from '@testing-library/react'
import { Provider } from 'react-redux'
import { configureStore } from '@reduxjs/toolkit'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import authReducer from '@/lib/redux/slices/authSlice'
import ordersReducer from '@/lib/redux/slices/ordersSlice'
import featureFlagsReducer from '@/lib/redux/slices/featureFlagsSlice'

const createTestStore = (initialState = {}) => {
  return configureStore({
    reducer: {
      auth: authReducer,
      orders: ordersReducer,
      featureFlags: featureFlagsReducer,
    },
    preloadedState: initialState,
  })
}

const createTestQueryClient = () => {
  return new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  })
}

interface AllTheProvidersProps {
  children: React.ReactNode
  store?: ReturnType<typeof createTestStore>
  queryClient?: ReturnType<typeof createTestQueryClient>
}

const AllTheProviders: React.FC<AllTheProvidersProps> = ({
  children,
  store = createTestStore(),
  queryClient = createTestQueryClient(),
}) => {
  return (
    <Provider store={store}>
      <QueryClientProvider client={queryClient}>
        {children}
      </QueryClientProvider>
    </Provider>
  )
}

interface CustomRenderOptions extends Omit<RenderOptions, 'wrapper'> {
  store?: ReturnType<typeof createTestStore>
  queryClient?: ReturnType<typeof createTestQueryClient>
}

const customRender = (
  ui: React.ReactElement,
  options?: CustomRenderOptions
) => {
  const { store, queryClient, ...renderOptions } = options || {}
  return render(ui, {
    wrapper: (props) => (
      <AllTheProviders store={store} queryClient={queryClient} {...props} />
    ),
    ...renderOptions,
  })
}

export * from '@testing-library/react'
export { customRender as render, createTestStore, createTestQueryClient }

