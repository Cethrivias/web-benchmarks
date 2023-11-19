import Fastify from 'fastify'

const fastify = Fastify({
    logger: false
})

fastify.get('/', (_request, reply) => {
    reply.send({ hello: 'world' })
})

fastify.listen({ port: 3000, host: "0.0.0.0" }, (err, address) => {
    if (err) throw err
    console.log(`Server is now listening on ${address}`)
})
