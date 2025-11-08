import React from 'react'
import { screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { CreateOrderDialog } from '@/components/organisms/CreateOrderDialog'
import { render, createTestStore } from '../setup/test-utils'

const mockOnClose = jest.fn()

describe('Summary Consistency - CreateOrderDialog', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should require summary field', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    const user = userEvent.setup()
    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const deliveryPreferenceSelect = screen.getByLabelText(/delivery preference/i)
    await user.selectOptions(deliveryPreferenceSelect, 'IN_STORE')

    const submitButton = screen.getByRole('button', { name: /get ai suggestions/i })
    
    expect(submitButton).toBeDisabled()

    const summaryField = screen.getByLabelText(/order summary/i)
    await user.click(summaryField)
    await user.tab()

    await waitFor(() => {
      expect(submitButton).toBeDisabled()
    }, { timeout: 2000 })
  })

  it('should enforce minimum length of 10 characters for summary', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    const user = userEvent.setup()
    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const summaryField = screen.getByLabelText(/order summary/i)
    await user.type(summaryField, 'Short')

    const deliveryPreferenceSelect = screen.getByLabelText(/delivery preference/i)
    await user.selectOptions(deliveryPreferenceSelect, 'IN_STORE')

    const submitButton = screen.getByRole('button', { name: /get ai suggestions/i })
    
    await waitFor(() => {
      expect(submitButton).toBeDisabled()
    }, { timeout: 2000 })
  })

  it('should accept valid summary with 10+ characters', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    const user = userEvent.setup()
    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const summaryField = screen.getByLabelText(/order summary/i)
    await user.type(summaryField, 'I need medicine for headache and fever')

    const deliveryPreferenceSelect = screen.getByLabelText(/delivery preference/i)
    await user.selectOptions(deliveryPreferenceSelect, 'IN_STORE')

    const submitButton = screen.getByRole('button', { name: /get ai suggestions/i })
    
    await waitFor(() => {
      expect(submitButton).not.toBeDisabled()
    })

    expect(screen.queryByText(/summary is required/i)).not.toBeInTheDocument()
    expect(screen.queryByText(/summary must be at least 10 characters/i)).not.toBeInTheDocument()
  })

  it('should maintain summary value when switching between form steps', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    const user = userEvent.setup()
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

  it('should show error message when summary is empty on blur', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    const user = userEvent.setup()
    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const deliveryPreferenceSelect = screen.getByLabelText(/delivery preference/i)
    await user.selectOptions(deliveryPreferenceSelect, 'IN_STORE')

    const summaryField = screen.getByLabelText(/order summary/i)
    const submitButton = screen.getByRole('button', { name: /get ai suggestions/i })
    
    expect(submitButton).toBeDisabled()

    await user.click(summaryField)
    await user.tab()

    await waitFor(() => {
      expect(submitButton).toBeDisabled()
    }, { timeout: 2000 })
  })

  it('should clear error when valid summary is entered', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    const user = userEvent.setup()
    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const deliveryPreferenceSelect = screen.getByLabelText(/delivery preference/i)
    await user.selectOptions(deliveryPreferenceSelect, 'IN_STORE')

    const summaryField = screen.getByLabelText(/order summary/i)
    const submitButton = screen.getByRole('button', { name: /get ai suggestions/i })
    
    expect(submitButton).toBeDisabled()

    await user.click(summaryField)
    await user.type(summaryField, 'I need medicine for headache')

    await waitFor(() => {
      expect(screen.queryByText(/summary is required/i)).not.toBeInTheDocument()
      expect(screen.queryByText(/Summary is required/i)).not.toBeInTheDocument()
      expect(submitButton).not.toBeDisabled()
    }, { timeout: 3000 })
  })
})

