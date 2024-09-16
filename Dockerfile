FROM golang:1.23.1-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download \
    && go build -o railway-volume-dump .

FROM alpine:3.20

WORKDIR /app
ENV NODE_ENV=production

COPY --from=builder /app/railway-volume-dump .

ENTRYPOINT ["/app/railway-volume-dump"]
