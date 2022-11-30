FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
ENV GOCACHE /app/cache
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY service/message/rpc/etc /app/etc
COPY /var/lib/docker/volumes/soft2_backend_go-cache/_data /app/cache
COPY cache/message/rpc /app/cache
RUN go build -ldflags="-s -w" -o /app/message service/message/rpc/message.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
COPY --from=builder /app/cache /app/cache
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/cache /app/cache
COPY --from=builder /app/message /app/message
COPY --from=builder /app/etc /app/etc

CMD ["./message", "-f", "etc/message.yaml"]
