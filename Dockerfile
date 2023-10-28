ARG GO_VERSION="1.21.3-bullseye"

FROM docker.io/golang:${GO_VERSION} AS build

ARG ARCHITECTURE="amd64"
ARG PLATFORM="linux"

ARG COMMIT_ID
ARG BUILD_TIMESTAMP
ARG BUILD_VERSION

RUN set -eux \
    && apt-get update \
    && apt-get clean \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/*
