# docker-go-wb

1. Start server, bd, nats-streaming:
`docker-compose up --build`
2. Create DB, table and rows
`stratup.sql`
3. To send generated messages into nats: `go run ./publish`
4. View on `http://localhost`
