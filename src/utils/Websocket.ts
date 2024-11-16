const connection = () => {
  const webSocket = new WebSocket('ws://127.0.0.1:5818?hutool=66813725')

  webSocket.onopen = () => {
    console.log('WebSocket opened')
  }

  webSocket.onmessage = () => {
    console.log('WebSocket message received')
  }

  webSocket.onerror = () => {
    console.log('WebSocket error received')
  }

  webSocket.onclose = () => {
    console.log('WebSocket closed')

    setTimeout(() => {
      connection();
    },  10000)
  }
}

export default connection
