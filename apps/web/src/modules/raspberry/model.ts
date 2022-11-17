import { Periphery } from '../periphery/model'

export interface Raspberry {
  ID: number
  Name: string
  IsActive: boolean
  PeripheryList: Periphery[]
}
