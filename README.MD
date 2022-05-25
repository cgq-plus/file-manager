# file-manager

#### 介绍
基于gin + Vue的文件管理系统

#### 软件架构
- 后端基于Gin框架
- 前端基于Vue + Element UI


#### 后端安装教程

1. 依赖于Gin框架
2. 克隆仓库
3. `go mod tidy` 后
4. 启动运行
5. 或者直接使用可执行文件

#### 前端安装教程

1. 进入file-managerui目录
2. `yarn install`
3. 启动运行 `yarn run dev`

### 后端部署
1. 编译打包
   1. 编译Linux amd64平台 `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/file_manager_linux_amd64  -ldflags "-s -w" main.go`
   2. 编译Windows amd64平台 `CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./out/file_manager_windows_amd64.exe  -ldflags "-s -w" main.go`
   3. 编译Darwin amd64平台 `CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./out/file_manager_darwin_amd64  -ldflags "-s -w" main.go`
   4. 编译Darwin arm64平台 `CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./out/file_manager_darwin_arm64  -ldflags "-s -w" main.go`
2. 编辑配置文件`conf/config.yaml`
3. 直接运行启动`./file_manager`

### 前端部署教程
1. 打包编译`yarn run build:prod`
2. 上传到Nginx
3. 修改Nginx反向代理(代理后端地址)
```
location /prod-api/{
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header REMOTE-HOST $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_connect_timeout 300;
    proxy_send_timeout 300;
    proxy_read_timeout 600;
    proxy_buffer_size 256k;
    proxy_buffers 4 256k;
    proxy_busy_buffers_size 256k;
    proxy_temp_file_write_size 256k;
    proxy_next_upstream error timeout invalid_header http_500 http_502 http_503 http_504 http_404;
    proxy_max_temp_file_size 128m;
    client_max_body_size 50m;
    proxy_pass http://127.0.0.1:8080/;
}
```