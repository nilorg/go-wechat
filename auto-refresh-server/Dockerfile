#build stage
FROM nilorg/golang:latest AS builder
WORKDIR /src
COPY . .
RUN go build -o ./bin/app -i ./auto-refresh-server/main.go

#final stage
FROM nilorg/alpine:latest
WORKDIR /workspace

ENV REDIS_ADDR=127.0.0.1:6379
ENV REDIS_PASSWORD=
ENV REDIS_DB=0
ENV WECHAT_APP_ID=
ENV WECHAT_APP_SECRET=
ENV WECHAT_REFRESH_DURATION=3600
ENV REDIS_ACCESS_TOKEN_KEY=
ENV REDIS_JS_API_TICKET_KEY=

COPY --from=builder /src/bin/app .
ENTRYPOINT ./app