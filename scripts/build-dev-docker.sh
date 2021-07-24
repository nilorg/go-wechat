#!/bin/bash
docker build -f auto-refresh-server/Dockerfile -t nilorg/wechat-auto-refresh-server:dev .
docker push nilorg/wechat-auto-refresh-server:dev