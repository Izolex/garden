import React from 'react'
import { render, cleanup } from '@testing-library/react'
import App from './App'

jest.mock('./routes', () => () => <div>mocked-component</div>)

afterEach(cleanup)

describe('App.tsx', () => {
  it('Should start App.tsx with mocked index', () => {
    const { getByText } = render(<App />)

    const elem = getByText('mocked-component')

    expect(elem).toBeTruthy()
  })
})
