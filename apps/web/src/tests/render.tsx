// eslint-disable-next-line import/named
import { render as rtlRender, RenderOptions, RenderResult } from '@testing-library/react'
import React, { ReactElement } from 'react'
import { ThemeProvider } from 'styled-components'
import { I18nextProvider } from 'react-i18next'
import GlobalTheme from '../constants/GlobalTheme'
import i18next from '../locale'

const AllProviders: React.JSXElementConstructor<{ children: React.ReactElement }> = ({ children }: { children: React.ReactElement }) => (
  <I18nextProvider i18n={i18next}>
    <ThemeProvider theme={GlobalTheme}>{children}</ThemeProvider>
  </I18nextProvider>
)

const render = (ui: ReactElement, options?: RenderOptions): RenderResult => rtlRender(ui, { wrapper: AllProviders, ...options })

export { render }
