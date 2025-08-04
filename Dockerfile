FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o genkeys

FROM alpine
COPY --from=builder /app/genkeys /usr/local/bin/genkeys
ENTRYPOINT ["genkeys"]
