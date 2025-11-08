import React from 'react'
import { screen, waitFor } from '@testing-library/react'
import { useRouter } from 'next/navigation'
import { AuthGuard } from '@/lib/auth/AuthGuard'
import { render, createTestStore } from '../setup/test-utils'

jest.mock('next/navigation', () => ({
  useRouter: jest.fn(),
}))

const mockRouter = {
  replace: jest.fn(),
  push: jest.fn(),
  prefetch: jest.fn(),
}

describe('AuthGuard', () => {
  beforeEach(() => {
    jest.clearAllMocks()
    ;(useRouter as jest.Mock).mockReturnValue(mockRouter)
    localStorage.clear()
  })

  describe('Protected Routes (requireAuth=true)', () => {
    it('should redirect to login when no token is present', async () => {
      const store = createTestStore({
        auth: { token: null, user: null, loading: false, error: null },
      })

      render(
        <AuthGuard requireAuth={true}>
          <div>Protected Content</div>
        </AuthGuard>,
        { store }
      )

      await waitFor(() => {
        expect(mockRouter.replace).toHaveBeenCalledWith('/login')
      })

      expect(screen.queryByText('Protected Content')).not.toBeInTheDocument()
    })

    it('should show content when token is in Redux store', async () => {
      const store = createTestStore({
        auth: {
          token: 'valid-token',
          user: { id: 1, email: 'test@example.com' },
          loading: false,
          error: null,
        },
      })

      render(
        <AuthGuard requireAuth={true}>
          <div>Protected Content</div>
        </AuthGuard>,
        { store }
      )

      await waitFor(() => {
        expect(screen.getByText('Protected Content')).toBeInTheDocument()
      })

      expect(mockRouter.replace).not.toHaveBeenCalled()
    })

    it('should show content when token is in localStorage', async () => {
      localStorage.setItem('token', 'local-storage-token')
      const store = createTestStore({
        auth: { token: null, user: null, loading: false, error: null },
      })

      render(
        <AuthGuard requireAuth={true}>
          <div>Protected Content</div>
        </AuthGuard>,
        { store }
      )

      await waitFor(() => {
        expect(screen.getByText('Protected Content')).toBeInTheDocument()
      })

      expect(mockRouter.replace).not.toHaveBeenCalled()
    })

    it('should show loading spinner while checking auth', () => {
      const store = createTestStore({
        auth: { token: null, user: null, loading: false, error: null },
      })

      render(
        <AuthGuard requireAuth={true}>
          <div>Protected Content</div>
        </AuthGuard>,
        { store }
      )

      expect(screen.getByRole('status')).toBeInTheDocument()
    })
  })

  describe('Auth Pages (requireAuth=false)', () => {
    it('should redirect to dashboard when token is present', async () => {
      const store = createTestStore({
        auth: {
          token: 'valid-token',
          user: { id: 1, email: 'test@example.com' },
          loading: false,
          error: null,
        },
      })

      render(
        <AuthGuard requireAuth={false}>
          <div>Login Page</div>
        </AuthGuard>,
        { store }
      )

      await waitFor(() => {
        expect(mockRouter.replace).toHaveBeenCalledWith('/dashboard')
      })

      expect(screen.queryByText('Login Page')).not.toBeInTheDocument()
    })

    it('should show login page when no token is present', async () => {
      const store = createTestStore({
        auth: { token: null, user: null, loading: false, error: null },
      })

      render(
        <AuthGuard requireAuth={false}>
          <div>Login Page</div>
        </AuthGuard>,
        { store }
      )

      await waitFor(() => {
        expect(screen.getByText('Login Page')).toBeInTheDocument()
      })

      expect(mockRouter.replace).not.toHaveBeenCalled()
    })

    it('should redirect when token is in localStorage', async () => {
      localStorage.setItem('token', 'local-storage-token')
      const store = createTestStore({
        auth: { token: null, user: null, loading: false, error: null },
      })

      render(
        <AuthGuard requireAuth={false}>
          <div>Login Page</div>
        </AuthGuard>,
        { store }
      )

      await waitFor(() => {
        expect(mockRouter.replace).toHaveBeenCalledWith('/dashboard')
      })
    })
  })
})

