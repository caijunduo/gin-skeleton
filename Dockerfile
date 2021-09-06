# 构建应用程序
FROM golang:1.15.2-alpine3.12 AS builder

# 1、编译
WORKDIR /build
RUN adduser -u 10001 -D go-runner

ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main

# 不用操作系统，直接运行可执行文件，几乎不占用额外的空间
#FROM scratch

# 最小化操作系统，大概占用 10MB 空间 ,便于 /bin/sh 调试
FROM alpine:3.12 AS final

RUN apk add -U tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 2、运行
WORKDIR /app
COPY --from=builder /build/main /app/
#COPY --from=builder /build/config /app/config
COPY --from=builder /etc/passwd /etc/passwd
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER go-runner
ENTRYPOINT ["./main"]
