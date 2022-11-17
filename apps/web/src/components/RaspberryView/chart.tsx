import React from 'react'
import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, Tooltip } from 'chart.js'
import type { ChartData, ChartOptions } from 'chart.js'
import { Line } from 'react-chartjs-2'

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Tooltip)

export const options: ChartOptions<'line'> = {
  responsive: true,
}

interface Props {
  data: ChartData<'line'>
}

const Chart: React.FC<Props> = ({ data }) => {
  return <Line options={options} data={data} />
}

export default Chart
