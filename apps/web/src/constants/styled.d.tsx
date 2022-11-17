import 'styled-components'
import theme from './GlobalTheme'

type ThemeType = typeof theme

declare module 'styled-components' {
  // eslint-disable-next-line
  export interface DefaultTheme extends ThemeType {}
}
