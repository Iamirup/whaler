# ------------------------------------------- Builder
FROM golang:alpine AS builder

# RUN apk add git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target="/root/.cache/go-build" go build -o /entrypoint

# ------------------------------------------- Runtime
FROM alpine:latest AS runtime

WORKDIR /app

COPY --from=builder /entrypoint .

ENTRYPOINT [ "./entrypoint", "migrate", "up", "&&", "./entrypoint", "server"]
