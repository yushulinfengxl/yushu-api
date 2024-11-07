# 使用 Ubuntu 20.04 作为基础镜像
FROM ubuntu:latest

# 防止在安装包时出现交互式提示
ENV DEBIAN_FRONTEND=noninteractive

# 创建文件夹
RUN mkdir -p /store
RUN mkdir -p /store/env

# 更新软件包列表并安装 wget、curl 和 sudo
RUN apt-get update && \
    apt-get install -y wget && \
    apt-get install -y curl && \
    apt-get install -y sudo && \
    apt-get install -y openssl && \
    apt-get clean && \
    rm -rf /var/zone/apt/lists/*

# 设置环境变量
ENV NODEJS_HOME=/store/env/nodejs
ENV NODEJS_PATH=/store/env/nodejs/bin
ENV PATH=$NODEJS_PATH:$PATH

# 将文件从主机复制到容器内 COPY 只复制本地文件，而 ADD 还可以复制
# COPY app.js /app/

# 远程 clone URL 并解压文件指定文件夹
#ADD https://dev.ilin.eu.org/iblog/nodejs.tar.gz /store/env
## 解压nodejs
#RUN tar -xvzf /store/env/nodejs.tar.gz -C /store/env

# 设置 root 用户的默认密码为 123456
RUN echo 'root:123456' | chpasswd

# 声明容器将监听的端口
# 监听ssh端口
EXPOSE 22
# 监听bt面板端口
EXPOSE 8888
# 监听项目1面板端口
EXPOSE 8881
# 监听项目2面板端口
EXPOSE 8882
# 监听项目3面板端口
EXPOSE 8883
# 监听项目4面板端口
EXPOSE 8884

## 创建一个名为 "app" 的用户，并设置其 UID 为 1000
#RUN useradd -u 1000 -ms /bin/bash app
## 将 /store/env 目录的所有文件所有者更改为 app 用户
#RUN chown -R app:app /store/env
## 切换到 app 用户
#USER app
## 设置工作目录为 /store/env
#WORKDIR /store/env

# 设置工作目录为 /app。在运行容器时，所有后续命令都将在这个目录下执行。这一步是可选的，可以根据需要进行更改
WORKDIR /app
# 挂载仓储卷
# VOLUME /xxx path

# 将默认用户设置为 root
USER root

# 设置标签
LABEL xiaolin_ubuntu=latest

# 设置容器启动时执行的命令
CMD ["/bin/bash"]

#docker run -d -p 8080:80 --name my_container docker2
# docker run --privileged -it --name new-docker2 -p 13022:22 -p 13888:8888 --name new-docker2 -d my-docker2:test /bin/bash
# docker command
# docker build --tag xiaolin-ubuntu:latest http://154.12.63.133:8999/v1/api/docker/mirror-x
# docker run -it --name xiaolin-ubuntu -p 18888:8888 -p 18881:8881 -p 18882:8882 -p 18883:8883 -p 18884:8884 --name xiaolin-ubuntu -d xiaolin-ubuntu:latest /bin/bash