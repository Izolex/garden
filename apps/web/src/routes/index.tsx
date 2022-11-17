import React from 'react'
import { useTranslation } from 'react-i18next'
import { gql, useQuery } from '@apollo/client'
import { subHours } from 'date-fns'
import i18next, { Locale } from '../locale'
import { Raspberry } from '../modules/raspberry/model'
import RaspberryView from '../components/RaspberryView'
import Loading from '../components/Loading'
import { H1 } from './styled'

const toTime = new Date()
const fromTime = subHours(toTime, 12)

const Index: React.FC = () => {
  const { t } = useTranslation()
  const { loading, error, data } = useQuery<{ RaspberryList: Raspberry[] }>(
    gql`
      query ($fromTime: Time!, $toTime: Time!) {
        RaspberryList(fromTime: $fromTime, toTime: $toTime) {
          ID
          Name
          PeripheryList {
            ID
            Name
            IsMeasurable
            ValueList {
              Value
              DateTime
            }
          }
        }
      }
    `,
    {
      variables: {
        fromTime: fromTime.toISOString(),
        toTime: toTime.toISOString(),
      },
    }
  )

  return (
    <>
      <H1>
        {i18next.language === Locale.CS && <>Zahrádečka</>}
        {i18next.language === Locale.EN && <>Garden</>}
        <br />
        🌱💧🌿💦🍀
      </H1>
      <br />
      <br />

      {loading && <Loading />}
      {error && <div>{t('global.error')}</div>}
      {!loading && !error && (
        <div>
          {data?.RaspberryList.length ? (
            <>
              {data?.RaspberryList?.map((rasp) => (
                <RaspberryView key={rasp.ID} raspberry={rasp} />
              ))}
            </>
          ) : (
            <>{t('routes.index.soon')} 🖤</>
          )}
        </div>
      )}
    </>
  )
}

export default Index
