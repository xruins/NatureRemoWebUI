# build nature-remo-web-ui 
FROM node:carbon as webui-builder

ARG VERSION=2
ARG TARGETPLATFORM
ARG BUILDPLATFORM

WORKDIR /usr/src/nature-remo-web-ui

COPY ./nature-remo-web-ui/package*.json .
RUN --mount=type=cache,target=/tmp/npm_cache npm_config_cache=/tmp/npm_cache npm install
COPY ./nature-remo-web-ui/ .
RUN npm run build

# build nature-remo-api-server
FROM golang:1.19 as server-builder

WORKDIR /app

COPY ./nature-remo-api-server/ ./
COPY --from=webui-builder /usr/src/nature-remo-web-ui/build ./pkg/server/build

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=$(echo ${TARGETPLATFORM} | cut -d'/' -f1) \
    GOARCH=$(echo ${TARGETPLATFORM} | cut -d'/' -f2) \
    GOARM=$(echo ${TARGETPLATFORM} | cut -d'/' -f3 | cut -c2) \
    CGO_ENABLED=0 \
    go build -ldflags="-w -s" -o /app/nature-remo-api-server

# build runtime image
FROM --platform=$TARGETPLATFORM gcr.io/distroless/static-debian11:debug-nonroot
LABEL org.opencontainers.image.source = "https://github.com/xruins/NatureRemoWebUI"

COPY --from=server-builder --chmod=777 /app/nature-remo-api-server /usr/local/bin/nature-remo-api-server

EXPOSE 3000

WORKDIR /usr/src
ENTRYPOINT ["nature-remo-api-server", "server"]
