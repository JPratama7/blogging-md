FROM golang:1-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./blogging-md"]