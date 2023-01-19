FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
ENV GOCACHE /build/cache
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY service/file/rpc/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/file service/file/rpc/main.go
WORKDIR /app
CMD ["./file", "-f", "etc/file.yaml"]
