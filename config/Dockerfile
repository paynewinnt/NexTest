# 使用官方Go镜像作为构建环境
FROM golang:1.16-alpine as builder

# 设置工作目录
WORKDIR /app

# 复制Go模块依赖文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o nextest .

# 使用alpine镜像开始新的构建阶段
# alpine被选择是因为它小巧但同时包含了常用的工具和库
FROM alpine:latest

# 为了安全性和便利性，安装ca-certificates
RUN apk --no-cache add ca-certificates

# 从builder阶段复制构建的二进制文件
COPY --from=builder /app/nextest /nextest

# 暴露端口（根据应用配置调整）
EXPOSE 8080

# 运行应用程序
CMD ["/nextest"]

