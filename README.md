# ZNJ[智能记] 中小团队协作工具

#### 🎉 简介

可以简单替代 tower 的基于文档的团队协作工具

- 支持待办事项管理
- 支持项目任务管理
- 支持 markdown，流程图文档

#### 🚀 源码安装

```bash
git clone https://github.com/itmisx/znj
cd znj/web_src
yarn install
yarn build
cd ../
go mod tidy
GOOS=linux GOARCH=amd64 go build -o znj main.go
# 可以自定义配置 /etc/znj.yaml
./znj
```

#### 🚀 Docker

```bash
git clone https://github.com/itmisx/znj
cd znj/docker
# 自定义配置文件
# 编辑docker-compose.yaml
# volumes:
#  - ./etc/znj.yaml:/etc/znj.yaml
docker-compose up -d
```

#### 🎉 截图

![1](doc/screenshot/img-1.jpg)

#### LICENSE
