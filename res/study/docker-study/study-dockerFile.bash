docker save -o /backup/myImage.tar(/打包路径) myImage:1.0(镜像id/名字)
docker save -o <output.tar> <image1> <image2> ... //多个镜像

# 结合Dockerfile文件使用命令构建镜像
docker build -t iblog:latest .

# 使用docker远程构建镜像
# http://154.12.63.133:8999/v1/api/docker/mirror-x
docker build --tag qwq:latest https://dev.ilin.eu.org/docker/image.php
docker build --tag xiaolin-ubuntu:latest http://154.12.63.133:8999/v1/api/docker/mirror-x

docker run -it -p 8999:8080 qwq:latest /bin/bash
docker run -it -p 8999:9000 -p 8993:3000 -p 8995:5000 --name qwq -d qwq:latest /bin/bash
docker run -it -p 10234:22 --name xiaolin-ubuntu -d xiaolin-ubuntu:latest /bin/bash
# -p(port) --name(container name) -d[不进入伪终端，但保持终端运行，但需要配合attach]
docker attach (container id / name)

# 这样就创建了一个名为 my_volume 的仓储卷。
docker volume create my_volume
# 在这个例子中，my_volume 仓储卷会挂载到容器内的 /path/in/container 路径上。
docker run -v my_volume:/path/in/container my_image
# 这将列出所有已创建的 Docker 仓储卷，包括它们的名称、驱动程序、状态等信息
docker volume ls


# qq bot docker commands
docker run -d \
-e ACCOUNT=909770874 \
-e WS_ENABLE=true \
-e WS_URLS='["ws://172.17.0.1:2536/OneBotv11"]' \
-e MESSAGE_POST_FORMAT=array \
--name napcat \
--net=bridge \
--mac-address="02:42:ac:11:00:02" \
--restart=always \
-v /qq:/root/.config/QQ \
-v /qq/napcat/config:/usr/src/app/napcat/config \
docker-hub.ixiaolin.com/mlikiowa/napcat-docker:latest

docker run -d \
  -e ACCOUNT=909770874 \
  -e WSR_ENABLE=true \
  -e WS_URLS='["ws://172.17.0.6:2536/OneBotv11"]' \
  -e MESSAGE_POST_FORMAT=array \
  --name napcat \
  --net=bridge \
  --mac-address="02:42:ac:11:00:02" \
  --restart=always \
  -v /qq:/root/.config/QQ \
  -v /qq/napcat/config:/usr/src/app/napcat/config \
  db2780d957b8