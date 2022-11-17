export interface Periphery {
  ID: number
  Name: string
  IsMeasurable: boolean
  ValueList: PeripheryValue[]
}

export interface PeripheryValue {
  Value: number
  DateTime: string
}
