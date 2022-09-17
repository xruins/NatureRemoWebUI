FROM golang:1.19 as builder

WORKDIR /app

COPY go.* ./
COPY . ./
RUN go build -ldflags="-w -s" -o /app/nature-remo-api-server

FROM gcr.io/distroless/static-debian11:nonroot
COPY --from=builder --chmod=755 /app/nature-remo-api-server /usr/local/bin/nature-remo-api-server
EXPOSE 3001
CMD ["nature-remo-api-server", "localhost", "3001"]
