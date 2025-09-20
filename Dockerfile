FROM golang:1.25.1 AS builder

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /bin/gelection .
COPY db/migrations /bin/db/migrations

FROM debian:bookworm-20250908

COPY --from=builder /bin/gelection /bin/gelection
COPY --from=builder /bin/db/migrations /bin/db/migrations

CMD ["/bin/gelection"]
