'use client'

import { useMemo, useState } from 'react'

import { ForgotPasswordContainer } from './forgot-password-container'
import { LoginContainer } from './login-container'
import { RegisterContainer } from './register-container'

export function AuthContainer() {
  const [mode, setMode] = useState<'login' | 'register' | 'forgot-password'>(
    'login'
  )

  const handleForgotPassword = () => {
    setMode('forgot-password')
  }

  const handleSwitchMode = () => {
    setMode((prevMode) => (prevMode === 'login' ? 'register' : 'login'))
  }

  const renderMode = useMemo(() => {
    switch (mode) {
      case 'login':
        return (
          <LoginContainer
            onForgotPassword={handleForgotPassword}
            onSwitchMode={handleSwitchMode}
          />
        )
      case 'register':
        return <RegisterContainer onSwitchMode={handleSwitchMode} />
      case 'forgot-password':
        return <ForgotPasswordContainer onSwitchMode={handleSwitchMode} />
    }
  }, [mode])

  return (
    <div className='flex w-full rounded-xl border p-6 shadow-sm'>
      {renderMode}
    </div>
  )
}
