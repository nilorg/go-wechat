#!/bin/bash
docker build -f auto-refresh-server/Dockerfile -t nilorg/wechat-auto-refresh-server:dev .
docker push nilorg/wechat-auto-refresh-server:dev

# docker build -f gateway/Dockerfile -t nilorg/wechat-gateway:dev .
# docker push nilorg/wechat-gateway:dev

# docker build -f proxy/Dockerfile -t nilorg/wechat-proxy:dev .
# docker push nilorg/wechat-proxy:dev