export interface Result {
  handler: string,
  code: string,
  message: string,
  data: any,
}

interface HandlerData {
  type: string,
  x: number,
  y: number,
}
export type HandlerDataArray = Array<HandlerData>

export interface HandlerResult {
  result: number,
}
