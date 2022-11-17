import { css } from 'styled-components'
import type { SimpleInterpolation } from 'styled-components'

export const textStroke = ({ color, width }: { color: string; width: string }): SimpleInterpolation => css`
  -webkit-text-fill-color: ${color};
  -webkit-text-stroke-color: black;
  -webkit-text-stroke-width: ${width};
`
