import React from 'react'
import { useTranslation } from 'react-i18next'
import Images from '../Images'
import Menu from '../Menu'
import { Wrapper, Main, Container, Footer } from './styled'

const triggerErr = () => {
  throw new Error('Palačinky se salámem!!!')
}

const DefaultLayout: React.FC = ({ children }: React.PropsWithChildren<React.ReactNode>) => {
  const { t } = useTranslation()

  return (
    <Wrapper>
      <Menu />
      <Main>
        <Container>{children}</Container>
      </Main>

      <Footer>
        {t('components.DefaultLayout.mungo')} &nbsp;
        <span onClick={triggerErr}>
          <Images.Heart />
        </span>
        &nbsp;
        <a rel="external" target="_blank" href="https://jantuzil.cz">
          Jan Tužil
        </a>
      </Footer>
    </Wrapper>
  )
}

export default DefaultLayout
