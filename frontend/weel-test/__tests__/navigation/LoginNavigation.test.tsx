import React from 'react'
import { screen, waitFor } from '@testing-library/react'
import { useRouter } from 'next/navigation'
import Home from '@/app/page'
import { render } from '../setup/test-utils'

jest.mock('next/navigation', () => ({
  useRouter: jest.fn(),
}))

const mockRouter = {
  replace: jest.fn(),
  push: jest.fn(),
  prefetch: jest.fn(),
}

describe('Login Navigation', () => {
  beforeEach(() => {
    jest.clearAllMocks()
    ;(useRouter as jest.Mock).mockReturnValue(mockRouter)
    localStorage.clear()
    if (window.localStorage.getItem) {
      (window.localStorage.getItem as jest.Mock).mockClear()
    }
  })

  it('should redirect to dashboard when user is logged in', async () => {
    localStorage.setItem('token', 'valid-token')

    render(<Home />)

    await waitFor(() => {
      expect(mockRouter.replace).toHaveBeenCalled()
      const lastCall = mockRouter.replace.mock.calls[mockRouter.replace.mock.calls.length - 1]
      expect(lastCall[0]).toBe('/dashboard')
    }, { timeout: 3000 })
  })

  it('should redirect to login when user is not logged in', async () => {
    localStorage.removeItem('token')

    render(<Home />)

    await waitFor(() => {
      expect(mockRouter.replace).toHaveBeenCalledWith('/login')
    })
  })

  it('should show loading spinner during navigation', () => {
    render(<Home />)

    expect(screen.getByRole('status')).toBeInTheDocument()
  })

  it('should check localStorage on mount', async () => {
    localStorage.clear()
    localStorage.setItem('token', 'test-token')
    
    const getItemSpy = window.localStorage.getItem

    render(<Home />)

    await waitFor(() => {
      expect(getItemSpy).toHaveBeenCalledWith('token')
    }, { timeout: 2000 })
  })
})

