import styled from 'styled-components'
import { textStroke } from '../../constants/mixins'

export const Name = styled.h2`
  margin-bottom: 0;
  color: var(--orage);
  font-size: 40px;
  ${(props) => textStroke({ color: props.theme.color.orange, width: '2px' })}
`

export const RaspberryID = styled.small`
  font-size: 11px;
`

export const ChartContainer = styled.div`
  h3 {
    ${(props) => textStroke({ color: props.theme.color.pink, width: '1px' })}
  }
`

export const Hr = styled.hr`
  color: var(--orange);
`
