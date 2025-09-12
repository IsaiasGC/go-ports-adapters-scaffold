FROM golang:1.24 AS builder

WORKDIR /app

ADD . .

RUN go mod download && \
    go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o build ./cmd/api
RUN go clean -modcache

FROM alpine

ARG version

WORKDIR /app

RUN apk add --no-cache tzdata
RUN apk add --no-cache ca-certificates

ENV VERSION=$version
ENV TZ="America/Mexico_City"

COPY --from=builder /app/build /app/build

EXPOSE 8080

CMD ["./build"]
