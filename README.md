# Eventserver

This is a little project to teach myself Nats Jetstream.

The project is simple, Jetstream itself, 1 consumer in Golang that prints events to stdout, and a producer which comes in the form of a GraphQL API in Node.js.

```
[mike@ouroboros eventserver]$ tree -L 1
.
├── consumer
├── docker-compose.yaml
├── producer
└── README.md

2 directories, 2 files

```

## Running it

`docker-compose up` will start everything, and when it's running you can connect to the producer on `localhost:3000`.

The GraphQL api only supports a single query `{ count }` which fires an event, which the consumer should print.
