import styled from 'styled-components'
import { textStroke } from '../constants/mixins'

export const H1 = styled.h1`
  padding-top: 50px;
  color: ${({ theme }) => theme.color.orange};
  font-size: 60px;
  font-weight: bold;
  ${(props) => textStroke({ color: props.theme.color.orange, width: '2px' })}

  @media (max-width: 390px) {
    font-size: 55px;
  }

  @media (max-width: 340px) {
    font-size: 45px;
  }

  svg {
    width: 50px !important;
    height: 50px;
    animation: rotating 2s linear infinite;
  }

  @keyframes rotating {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }
`
