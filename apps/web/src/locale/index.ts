// eslint-disable-next-line import/named
import i18next, { i18n } from 'i18next'
import { initReactI18next } from 'react-i18next'
import csTranslation from './cs.json'
import enTranslation from './en.json'

const resources = {
  cs: {
    translation: csTranslation,
  },
  en: {
    translation: enTranslation,
  },
}

i18next.use(initReactI18next).init({
  resources,
  lng: 'cs',
  fallbackLng: ['cs', 'en'],
  interpolation: {
    escapeValue: false,
  },
  simplifyPluralSuffix: true,
  returnNull: false,
  react: {
    useSuspense: false,
  },
})

export enum Locale {
  CS = 'cs',
  EN = 'en',
}

interface i18 extends i18n {
  readonly language: Locale
}

export default i18next as i18
