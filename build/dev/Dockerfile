# Build stage
FROM golang:1.19-alpine3.16 AS builder
RUN apk add --no-cache git gcc musl-dev
ARG BITBUCKET_TOKEN
WORKDIR /app
COPY ../.. .
RUN git config --global --add url."https://x-token-auth:${BITBUCKET_TOKEN}@github.com/nguyenkhoa0721".insteadOf "https://github.com/nguyenkhoa0721"
RUN go mod download
RUN go build -o main -tags musl cmd/multichain-explorer/main.go

# Run stage
FROM alpine:3.16 as runner
WORKDIR /app
COPY --from=builder /app/main .
COPY ../../config.yml .

EXPOSE 8000
CMD [ "/app/main" ]