import React from 'react'
import { useTranslation } from 'react-i18next'
import Images from '../Images'
import i18next from '../../locale'
import { Gitlab, Nav } from './index.styled'
import Flags from './flags'

const Menu: React.FC = () => {
  useTranslation()

  return (
    <Nav>
      <Flags lang={i18next.language} />
      <ul>
        <li>
          <a href={window.Config.API_URL + '/doc/redoc'} rel="external" target="_blank">
            OpenAPI
          </a>
        </li>
        <li>
          <Gitlab href="https://gitlab.honzas.garden" rel="external" target="_blank">
            <Images.Gitlab />
          </Gitlab>
        </li>
        <li>
          <a href={window.Config.API_URL + '/graphql'} rel="external" target="_blank">
            GraphiQL
          </a>
        </li>
      </ul>
    </Nav>
  )
}

export default Menu
