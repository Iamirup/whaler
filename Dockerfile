# ------------------------------------------- Builder
FROM golang:alpine AS builder

# RUN apk add git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /entrypoint

# ------------------------------------------- Runtime
FROM alpine:latest AS runtime

WORKDIR /app

COPY --from=builder /entrypoint .

COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
