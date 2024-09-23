# 使用 Go 作为基础镜像
FROM golang:1.22.3 AS builder

# 设置工作目录
WORKDIR /app

# 复制 Go 模块文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用并生成静态二进制文件
RUN CGO_ENABLED=0 GOOS=linux go build -o main . && chmod +x main && ls -l

# 使用轻量级的镜像
FROM alpine:latest

# 将构建的二进制文件复制到新镜像中
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/saved_text.txt .
# 提供静态文件支持
COPY static/ ./static/


# 运行应用
CMD ["./main"]