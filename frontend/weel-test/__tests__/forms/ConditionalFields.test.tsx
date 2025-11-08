import React from 'react'
import { screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { CreateOrderDialog } from '@/components/organisms/CreateOrderDialog'
import { render, createTestStore } from '../setup/test-utils'

const mockOnClose = jest.fn()

describe('Conditional Fields - CreateOrderDialog', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should not show delivery address field when IN_STORE is selected', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const deliveryPreferenceSelect = screen.getByLabelText(/delivery preference/i)
    await userEvent.selectOptions(deliveryPreferenceSelect, 'IN_STORE')

    await waitFor(() => {
      expect(screen.queryByLabelText(/delivery address/i)).not.toBeInTheDocument()
    })
  })

  it('should show delivery address field when DELIVERY is selected', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const deliveryPreferenceSelect = screen.getByLabelText(/delivery preference/i)
    await userEvent.selectOptions(deliveryPreferenceSelect, 'DELIVERY')

    await waitFor(() => {
      expect(screen.getByLabelText(/delivery address/i)).toBeInTheDocument()
    })
  })

  it('should show delivery address field when CURBSIDE is selected', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const deliveryPreferenceSelect = screen.getByLabelText(/delivery preference/i)
    await userEvent.selectOptions(deliveryPreferenceSelect, 'CURBSIDE')

    await waitFor(() => {
      expect(screen.getByLabelText(/delivery address/i)).toBeInTheDocument()
    })
  })

  it('should require delivery address when DELIVERY is selected', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    const user = userEvent.setup()
    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const summaryField = screen.getByLabelText(/order summary/i)
    await user.type(summaryField, 'I need medicine for headache and fever')

    const deliveryPreferenceSelect = screen.getByLabelText(/delivery preference/i)
    await user.selectOptions(deliveryPreferenceSelect, 'DELIVERY')

    await waitFor(() => {
      expect(screen.getByLabelText(/delivery address/i)).toBeInTheDocument()
    })

    const submitButton = screen.getByRole('button', { name: /get ai suggestions/i })
    
    expect(submitButton).toBeDisabled()

    const addressField = screen.getByLabelText(/delivery address/i)
    await user.click(addressField)
    await user.tab()

    await waitFor(() => {
      expect(submitButton).toBeDisabled()
    }, { timeout: 2000 })
  })

  it('should not require delivery address when IN_STORE is selected', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    const user = userEvent.setup()
    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const summaryField = screen.getByLabelText(/order summary/i)
    await user.type(summaryField, 'I need medicine for headache and fever')

    const deliveryPreferenceSelect = screen.getByRole('combobox', { name: /delivery preference/i }) || screen.getByLabelText(/delivery preference/i)
    await user.selectOptions(deliveryPreferenceSelect, 'IN_STORE')

    const submitButton = screen.getByRole('button', { name: /get ai suggestions/i })
    await user.click(submitButton)

    await waitFor(() => {
      expect(screen.queryByText(/address is required/i)).not.toBeInTheDocument()
    })
  })

  it('should hide delivery address when switching from DELIVERY to IN_STORE', async () => {
    const store = createTestStore({
      featureFlags: { flags: { ai_suggestions: true }, loading: false },
      orders: { aiSuggestions: [], loading: false },
    })

    const user = userEvent.setup()
    render(<CreateOrderDialog onClose={mockOnClose} />, { store })

    const deliveryPreferenceSelect = screen.getByRole('combobox', { name: /delivery preference/i }) || screen.getByLabelText(/delivery preference/i)
    
    await user.selectOptions(deliveryPreferenceSelect, 'DELIVERY')
    await waitFor(() => {
      expect(screen.getByLabelText(/delivery address/i)).toBeInTheDocument()
    })

    await user.selectOptions(deliveryPreferenceSelect, 'IN_STORE')
    await waitFor(() => {
      expect(screen.queryByLabelText(/delivery address/i)).not.toBeInTheDocument()
    })
  })
})

