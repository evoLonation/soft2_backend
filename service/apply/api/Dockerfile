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
COPY service/apply/api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/apply service/apply/api/apply.go
WORKDIR /app
CMD ["./apply", "-f", "etc/apply-api.yaml"]
