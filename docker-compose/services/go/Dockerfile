FROM golang:1.18-bullseye as build

# 容器环境变量添加，会覆盖默认的变量值
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作区
WORKDIR /data

# 把全部文件添加到/go/release目录
ADD ../ .

# 编译：把cmd/main.go编译成可执行的二进制文件，命名为app
RUN go build main.go

# 运行：使用scratch作为基础镜像
FROM scratch as prod

# 在build阶段复制时区到
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# 在build阶段复制可执行的go二进制文件app
COPY --from=build /main /

COPY ./config.example.yaml /
# 启动服务
CMD ./main serve