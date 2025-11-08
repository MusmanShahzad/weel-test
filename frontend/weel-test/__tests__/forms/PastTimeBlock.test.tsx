import React from 'react'
import { screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { CreateOrderDialog } from '@/components/organisms/CreateOrderDialog'
import { render, createTestStore } from '../setup/test-utils'

const mockOnClose = jest.fn()

describe('Past Time Block - CreateOrderDialog', () => {
  beforeEach(() => {
    jest.clearAllMocks()
    jest.useFakeTimers()
  })

  afterEach(() => {
    jest.runOnlyPendingTimers()
    jest.useRealTimers()
  })

  it('should validate that summary is consistent across form steps', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    const user = userEvent.setup({ delay: null })
    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const summaryText = 'I need medicine for headache and fever'
    const summaryField = screen.getByLabelText(/order summary/i)
    await user.type(summaryField, summaryText)

    const deliveryPreferenceSelect = screen.getByLabelText(/delivery preference/i)
    await user.selectOptions(deliveryPreferenceSelect, 'IN_STORE')

    const submitButton = screen.getByRole('button', { name: /get ai suggestions/i })
    await user.click(submitButton)

    await waitFor(() => {
      expect(screen.getByText(/ai suggested products/i)).toBeInTheDocument()
    })

    const backButton = screen.getByRole('button', { name: /back/i })
    await user.click(backButton)

    await waitFor(() => {
      const summaryFieldAfterBack = screen.getByLabelText(/order summary/i)
      expect(summaryFieldAfterBack).toHaveValue(summaryText)
    })
  })

  it('should maintain form state consistency when navigating back', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    const user = userEvent.setup({ delay: null })
    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const summaryText = 'I need medicine for headache and fever'
    const addressText = '123 Main St, New York, NY 10001'

    const summaryField = screen.getByLabelText(/order summary/i)
    await user.type(summaryField, summaryText)

    const deliveryPreferenceSelect = screen.getByLabelText(/delivery preference/i)
    await user.selectOptions(deliveryPreferenceSelect, 'DELIVERY')

    await waitFor(() => {
      const addressField = screen.getByLabelText(/delivery address/i)
      expect(addressField).toBeInTheDocument()
    })

    const addressField = screen.getByLabelText(/delivery address/i)
    await user.type(addressField, addressText)

    const submitButton = screen.getByRole('button', { name: /get ai suggestions/i })
    await user.click(submitButton)

    await waitFor(() => {
      expect(screen.getByText(/ai suggested products/i)).toBeInTheDocument()
    })

    const backButton = screen.getByRole('button', { name: /back/i })
    await user.click(backButton)

    await waitFor(() => {
      expect(screen.getByLabelText(/order summary/i)).toHaveValue(summaryText)
      expect(screen.getByLabelText(/delivery address/i)).toHaveValue(addressText)
      expect(screen.getByLabelText(/delivery preference/i)).toHaveValue('DELIVERY')
    })
  })

  it('should preserve summary when toggling delivery preference', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    const user = userEvent.setup({ delay: null })
    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const summaryText = 'I need medicine for headache and fever'
    const summaryField = screen.getByLabelText(/order summary/i)
    await user.type(summaryField, summaryText)

    const deliveryPreferenceSelect = screen.getByLabelText(/delivery preference/i)
    
    await user.selectOptions(deliveryPreferenceSelect, 'IN_STORE')
    await user.selectOptions(deliveryPreferenceSelect, 'DELIVERY')
    await user.selectOptions(deliveryPreferenceSelect, 'CURBSIDE')
    await user.selectOptions(deliveryPreferenceSelect, 'IN_STORE')

    await waitFor(() => {
      expect(summaryField).toHaveValue(summaryText)
    })
  })
})

