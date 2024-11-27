import { type Result, type HandlerResult, type HandlerDataArray } from '@/types/Websocket'
import { GetAddress } from '@/wailsjs/go/main/Server'
import Emitter from '@/utils/Emitter'
import { ElNotification } from 'element-plus'

const connection = () => {
  GetAddress().then((address) => {
    const webSocket = new WebSocket('ws://' + address)

    webSocket.onopen = () => {}

    webSocket.onmessage = (event) => {
      try {
        const result:Result = JSON.parse(event.data)
        if (result.code == 'OK') {
          switch (result.handler) {
            case 'data':
            {
              const data: HandlerDataArray = JSON.parse(JSON.stringify(result.data))
              Emitter.emit('change-data', data)
            }
              break
            case 'result':
            {
              const data: HandlerResult = JSON.parse(JSON.stringify(result.data))
              Emitter.emit('change-result', data)
            }
              break
          }
        } else {
          ElNotification({
            type: 'error',
            title: result.code,
            message: result.message,
          })
        }
      } catch (e) {
        console.error((e as Error).message)
      }
    }

    webSocket.onerror = () => {}

    webSocket.onclose = () => {
      setTimeout(() => {
        connection()
      }, 10000)
    }
  })
}

export default connection
