#!/bin/bash
docker build -f auto-refresh-server/Dockerfile -t nilorg/naas-token-server:dev .
# docker push nilorg/naas-token-server:dev