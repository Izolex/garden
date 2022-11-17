import 'bootstrap/dist/css/bootstrap.min.css'
import React from 'react'
import * as Sentry from '@sentry/browser'
import { withProfiler } from '@sentry/react/dist/profiler'
import { ErrorBoundary } from '@sentry/react/dist/errorboundary'
import { Integrations } from '@sentry/tracing'
import { ThemeProvider } from 'styled-components'
import { I18nextProvider } from 'react-i18next'
import { ApolloClient, InMemoryCache, ApolloProvider } from '@apollo/client'
import Index from './routes'
import i18next from './locale'
import GlobalTheme from './constants/GlobalTheme'
import { GlobalStyles } from './constants/GlobalStyle'
import DefaultLayout from './components/DefaultLayout'
import ErrorBoundaryScreen from './components/ErrorBoundaryScreen'

if (process.env.NODE_ENV === 'development') {
  window.Config = {
    API_URL: process.env.API_URL as string,
    SENTRY_DSN: process.env.SENTRY_DSN as string,
  }
}

Sentry.init({
  dsn: window.Config.SENTRY_DSN,
  release: process.env.VERSION,
  integrations: [new Integrations.BrowserTracing()],
  tracesSampleRate: 0.05,
})

const apolloClient = new ApolloClient({
  uri: window.Config.API_URL + '/graphql',
  cache: new InMemoryCache(),
})

const App: React.FC = () => {
  return (
    <ErrorBoundary fallback={() => <ErrorBoundaryScreen />}>
      <ApolloProvider client={apolloClient}>
        <I18nextProvider i18n={i18next}>
          <ThemeProvider theme={GlobalTheme}>
            <GlobalStyles />
            <DefaultLayout>
              <Index />
            </DefaultLayout>
          </ThemeProvider>
        </I18nextProvider>
      </ApolloProvider>
    </ErrorBoundary>
  )
}

export default withProfiler(App)
