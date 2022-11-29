FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
ENV GOCACHE /root/.cache/go-build
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

VOLUME ./cache:/root/.cache/go-build
WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY service/paper/api/etc /app/etc
RUN cp service/paper/api/main.go /root/.cache/go-build
RUN go build -ldflags="-s -w" -o /app/main service/paper/api/main.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/etc /app/etc

CMD ["./main", "-f", "etc/paper-api.yaml"]
