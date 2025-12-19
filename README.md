## FlowBlog Server

### 该项目为博客系统的后端，使用go语言，框架gin编写

#### 前端项目：[AArabbit/FLOW-BLOG-VUE: 使用vue3,typescript,vite](https://github.com/AArabbit/FLOW-BLOG-VUE)

#### 有详细的API文档(站点可F12查看接口)：[API文档](https://rabbitwebsite.top/codex/3)

#### 本地运行方式：

1. 修改项目根目录 `config/.env.dev` 文件，根据注释修改mysql与redis配置

2. 如果你想自动建数据库表，打开`main.go`，将自动建表下的注释打开

3. 修改 `config/config.yml`里面的`email`配置，`github`配置

4. 终端运行命令

   ```bash
   go mod tidy
   go build main.go
   ```

看到终端 `Listening and serving HTTP on :8080` 就ok



#### 服务器部署：

1. 在以上基础上再次修改`config/config.yml`
2. 修改mysql配置，host不用改，docker进行管理，修改`user`与`password`
3. redis同理
4. 确认email与github配置
5. 确认`docker-compose.yml`里的配置，重点mysql的密码
6. 如果想要自动建表，打开`main.go`，将自动建表下的注释打开
7. 打开根目录下`pack.sh`，逐一执行里面的命令
8. 根目录生成linux的二进制文件`blog-app`
9. 把`Dockerfile`，`docker-compose.yml`，`blog-app`上传到服务器目录
10. 在服务器上传目录下(与`blog-app`同级)，新建config文件夹，将项目里`config`目录下的`ip2region_v4.xdb`，`config.yml` 放到服务器config文件夹里
11. 确认服务器docker与go环境安装正确，服务器终端进入上传目录，以ubnutn为例，执行`sudo docker compose build -d --build`，执行`sudo docker compose down`可停止并删除











