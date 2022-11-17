import React from 'react'
import type { ChartData } from 'chart.js'
import { Col, Container, Row } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'
import { Raspberry } from '../../modules/raspberry/model'
import { PeripheryValue } from '../../modules/periphery/model'
import theme from '../../constants/GlobalTheme'
import Chart from './chart'
import { ChartContainer, Name, RaspberryID, Hr } from './styled'

interface Props {
  raspberry: Raspberry
}

const createChartData = (values: PeripheryValue[]): ChartData<'line'> => {
  const labels: string[] = []
  const data: number[] = []

  values.forEach((val) => {
    labels.push(new Date(val.DateTime).getHours().toString() + ':00')
    data.push(val.Value)
  })

  return {
    labels,
    datasets: [
      {
        data,
        borderColor: theme.color.orange,
        backgroundColor: theme.color.orange,
      },
    ],
  }
}

const RaspberryView: React.FC<Props> = ({ raspberry }) => {
  const { t } = useTranslation()

  return (
    <div>
      <Name>{raspberry.Name}</Name>
      <RaspberryID>raspberry #{raspberry.ID}</RaspberryID>
      <Hr />

      <Container fluid="md">
        <Row>
          {raspberry.PeripheryList.filter((periphery) => periphery.IsMeasurable).map((periphery) => {
            return (
              <Col key={`${raspberry.ID}x${periphery.ID}`} md={6} sm={12}>
                <ChartContainer>
                  <h3>{t(`periphery.${periphery.Name}.value`)}</h3>
                  <Chart data={createChartData(periphery.ValueList)} />
                </ChartContainer>
              </Col>
            )
          })}
        </Row>
      </Container>
    </div>
  )
}

export default RaspberryView
