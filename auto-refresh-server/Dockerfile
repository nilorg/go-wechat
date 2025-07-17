#build stage
FROM golang:1.18-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o ./bin/app ./auto-refresh-server/main.go

#final stage
FROM alpine:latest
WORKDIR /workspace

COPY --from=builder /src/bin/app .
ENTRYPOINT ./app