import express from 'express'
import { graphqlHTTP } from 'express-graphql'
import {
  GraphQLSchema,
  GraphQLString,
  GraphQLObjectType,
  GraphQLNonNull,
  GraphQLInt,
} from 'graphql'

let count = 0
export function Server({ context }) {
  const query = new GraphQLObjectType({
    name: 'Query',
    fields: () => ({
      count: {
        type: GraphQLString,
        resolve: async (root, args, { publish }) => {
          // to publish messages to a stream:
          count++
          const pa = publish({ channel: 'channels.count', event: { count } })
          return count
        },
      },
      domain: {
        type: GraphQLString,
        args: {
          name: {
            description: 'the domain',
            type: new GraphQLNonNull(GraphQLString),
          },
        },
        resolve: async (root, { name }, { publish }) => {
          // to publish messages to a stream:
          const pa = publish({
            channel: 'channels.domains',
            event: { domain: name },
          })
          return name
        },
      },
    }),
  })

  let server = express()

  server.use(
    '/',
    graphqlHTTP({
      schema: new GraphQLSchema({ query }),
      graphiql: true,
      context,
    }),
  )
  return server
}
