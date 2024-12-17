基于windows docker搭建lamp环境
===
# 整理需要的容器
| 容器名称  | 作用      | 备注                  |
|-------|---------|---------------------|
| nginx | 代理转发    | 配置需要挂载出来，避免每次构建丢失配置 |
| mysql | 存储数据    | 数据需要挂载出来避免丢失        |
| php   | 运行php服务 | 需要通过nginx代理转发       |
# 拉取对应镜像
nginx: docker pull nginx可以拉取最新

New-Item -Path 'd:\' -Name 'minikube' -ItemType Directory -Force
Invoke-WebRequest -OutFile 'd:\minikube\minikube.exe' -Uri 'https://github.com/kubernetes/minikube/releases/latest/download/minikube-windows-amd64.exe' -UseBasicParsing