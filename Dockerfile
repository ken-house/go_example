FROM golang:1.18-alpine3.16

# 维护者
MAINTAINER Ken

# 设置环境变量
ENV GO111MODULE=on CGO_ENABLE=0 GOOS=linux GO_ARCH=amd64 GOPROXY=https://goproxy.cn,direct

# 将本地项目文件拷贝到容器当前目录
COPY . /build

# 切换容器的目录到
WORKDIR /build

# 执行go build生成可执行文件example-http-server
RUN go build -o example-http-server main.go


# 切换到容器下的目标目录
WORKDIR /dist

# 将需要加载的本地配置文件copy到容器下的/dist目录
COPY ./configs ./configs
COPY ./assets  ./assets
COPY ./views   ./views

# 向容器添加卷
VOLUME ["/dist/configs","/dist/nacos","/dist/logs"]

# 将容器内的可执行文件拷贝到容器的当前目录
RUN cp /build/example-http-server .

# 暴露端口
EXPOSE 8080

# 启动容器时运行项目，并指定参数http
ENTRYPOINT ["/dist/example-http-server","http"]


