FROM golang:1.19 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main
RUN mkdir publish  \
    && cp main publish  \
    && cp -r conf publish


FROM busybox:1.28.4

WORKDIR /app

COPY --from=builder /app/publish .
# 指定运行时环境变量
ENV GIN_MODE=release
EXPOSE 3000
ENTRYPOINT ["./main"]

##mysql ,redis 是golang 中不同服务，如果单个golang 服务我用dockerFile 容器启动就行，但是项目中有mysql和redis等服务，
##需要多服务启动那就用到docker-compose.yml