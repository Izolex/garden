import React from 'react'
import i18next, { Locale } from '../../locale'
import Images from '../Images'
import { Flag } from './flags.styled'

interface FlagsProps {
  lang: Locale
}

const Flags: React.FC<FlagsProps> = ({ lang }) => {
  switch (lang) {
    case Locale.CS:
      return (
        <Flag onClick={() => i18next.changeLanguage(Locale.EN)}>
          <Images.Flags.EN />
        </Flag>
      )
    case Locale.EN:
      return (
        <Flag onClick={() => i18next.changeLanguage(Locale.CS)}>
          <Images.Flags.CZ />
        </Flag>
      )
    default:
      throw new Error(`Undefined locale "${lang}"`)
  }
}

export default Flags
