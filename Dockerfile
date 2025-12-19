# ===========================
# 构建阶段 (Builder)
# ===========================
FROM golang:1.25.5-alpine3.23 AS builder

# 基础依赖
RUN apk --no-cache add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app

# 直接复制本地编译好的文件
COPY blog-app .
RUN chmod +x blog-app
COPY config ./config
# COPY static ./static  <-- 如果有

EXPOSE 8080
CMD ["./blog-app"]