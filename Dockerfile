FROM alpine:latest

# 维护者
MAINTAINER Ken

# 切换到容器下的目标目录
WORKDIR /

# 将本地编译好的可执行文件复制到容器中
COPY ./bin/go-example go-example

# 将需要加载的本地配置文件copy到容器下的/dist目录
COPY ./views   ./views
COPY ./assets/jenkins ./assets/jenkins

# 向容器添加卷
VOLUME ["/logs", "/nacos"]

# 容器运行时执行命令
ENTRYPOINT ["/go-example"]


