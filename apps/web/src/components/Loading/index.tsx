import Spinner from 'react-bootstrap/Spinner'
import React from 'react'
import theme from '../../constants/GlobalTheme'

const Loading: React.FC = () => {
  return <Spinner animation="border" style={{ color: theme.color.orange }} />
}

export default Loading
