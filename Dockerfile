# 使用Alpine Linux作为基础镜像
FROM alpine:latest

# 安装SQLite
RUN apk add --no-cache sqlite

# 设置工作目录
WORKDIR /db

# 保持容器运行
ENTRYPOINT ["tail", "-f", "/dev/null"]
