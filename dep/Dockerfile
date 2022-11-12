FROM golang:1.19.2 AS build

RUN set -eux \
    && apt-get update \
    && apt-get upgrade -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /erebus

COPY . .

ARG CGO_ENABLED
ARG VERSION
ARG TIMESTAMP
ARG GOARCH
ARG GOOS

ENV FLAGS "-X 'main.version=${VERSION}' -X 'main.timestamp=${TIMESTAMP}'"

RUN go mod tidy \
    && CGO_ENABLED=${CGO_ENABLED} GOARCH=${GOARCH} GOOS=${GOOS} go build -o /bin/erebus -a -ldflags="${FLAGS}" -v .

FROM alpine:3.16.2 AS release

RUN set -eux \
    && apk update \
    && apk upgrade \
    && rm -rf /var/cache/apk/*

COPY --from=build /bin/erebus /bin/erebus

WORKDIR "/erebus"

EXPOSE 8080

VOLUME ["/erebus"]

ENTRYPOINT ["/bin/erebus"]
