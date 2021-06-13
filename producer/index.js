import { connect, StringCodec, JSONCodec, Empty, AckPolicy } from 'nats'
import { Server } from './src/Server.js'
const { PORT = 3000, HOST = '0.0.0.0' } = process.env

;(async () => {
  const nc = await connect({ url: 'localhost:4222' })
  const jsm = await nc.jetstreamManager()

  const streams = await jsm.streams.list().next()
  await jsm.streams.add({ name: 'channels', subjects: ['channels.*'] })

  // create a jetstream client:
  const js = nc.jetstream()
  const jc = JSONCodec()
  const sc = StringCodec()
  const publish = (event) => {
    const encoded = jc.encode(event)
    console.log({ publishing: event, encoded })
    nc.publish('channels.count', encoded)
  }

  const server = Server({ context: { publish } })

  server.listen({ port: PORT, host: HOST }, () =>
    console.log(`🚀 Tracker listening on ${HOST}:${PORT}`),
  )

  process.on('SIGTERM', () => process.exit(0))
  process.on('SIGINT', () => process.exit(0))
})()
