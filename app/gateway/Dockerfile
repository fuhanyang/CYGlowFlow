FROM golang:1.24-alpine AS builder

WORKDIR /app

# 构建参数
ARG BUILD_ENV=dev
ENV ENVIRONMENT=$BUILD_ENV

# 复制go模块文件
COPY app/gateway/go.mod app/gateway/go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 根据环境优化构建
RUN cd app/gateway && \
    if [ "$ENVIRONMENT" = "prod" ]; then \
        CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o main . ; \
    else \
        CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main . ; \
    fi

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=builder /app/app/gateway/main .

# 根据环境复制对应的配置文件
RUN if [ "$BUILD_ENV" = "test" ]; then \
        cp -r app/gateway/conf/test/* ./conf/ ; \
    elif [ "$BUILD_ENV" = "prod" ]; then \
        cp -r app/gateway/conf/online/* ./conf/ ; \
    else \
        cp -r app/gateway/conf/dev/* ./conf/ ; \
    fi

EXPOSE 8888

# 设置环境变量
ENV ENVIRONMENT=$BUILD_ENV

CMD ["./main"]
