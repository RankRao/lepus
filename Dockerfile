FROM golang:alpine AS builder
LABEL MAINTAINER sheaven <sheaven@qq.com>

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

RUN set -xe \
    && sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories \
    && apk update --no-cache \
    && apk add --no-cache gcc musl-dev

WORKDIR /build/lepus

COPY . .
RUN set -xe \
    && go mod download \
    && chmod 755 ./build.sh && ./build.sh

FROM alpine
LABEL MAINTAINER sheaven <sheaven@qq.com>

WORKDIR /app/lepus

RUN set -xe \
    && sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories \
    && apk update --no-cache \
    && apk add --no-cache ca-certificates tzdata supervisor mysql-client
ENV TZ Asia/Shanghai

COPY docker/supervisor/supervisord.conf /etc/supervisor/supervisord.conf
COPY docker/supervisor/env.conf /etc/supervisor/conf.d/env.conf
COPY --from=builder /build/lepus/bin /app/lepus/bin
COPY etc/alarm.example.ini /app/lepus/etc/alarm.ini
COPY etc/config.example.ini /app/lepus/etc/config.ini
COPY etc/proxy.example.ini /app/lepus/etc/proxy.ini
COPY init_table.sql  /app/lepus/init_table.sql
COPY init_data.sql  /app/lepus/init_data.sql
COPY entrypoint.sh /app/lepus/entrypoint.sh

ENV DEBUG=false
RUN chmod 755 /app/lepus/entrypoint.sh

EXPOSE 8080

CMD ["/app/lepus/entrypoint.sh"]