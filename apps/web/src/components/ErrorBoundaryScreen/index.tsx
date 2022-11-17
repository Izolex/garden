import React from 'react'
import { ErrorDiv } from './styled'

const ErrorBoundaryScreen: React.FC = () => {
  return (
    <ErrorDiv>
      <h1>Fucking fuck, something really bad happened :/</h1>
    </ErrorDiv>
  )
}

export default ErrorBoundaryScreen
