import mitt from 'mitt'
import { type HandlerDataArray, type HandlerResult } from '@/types/Websocket'

type Events = {
  'change-data': HandlerDataArray,
  'change-result': HandlerResult,
}

export default mitt<Events>()
