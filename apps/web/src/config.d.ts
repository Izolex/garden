interface Config {
  API_URL: string
  VERSION: string
  SENTRY_DSN: string
}

export declare global {
  interface Window {
    Config
  }
  namespace NodeJS {
    interface ProcessEnv extends Config {
      NODE_END: 'development'
    }
  }
}
