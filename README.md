# gin-api

> 目前只适用于http项目支持

### 架构
- corn 异步定时任务
- base64Captcha 灵活验证码
- viper 配置服务
- casbin 权限服务
- redis 缓存服务
- rabbit 队列服务

### 快速开始

1. 下载并安装golang,[下载地址](https://go.dev/dl/)->选择合适的版本
2. 切换国内golang镜像源
   - windows
```shell
 SETX GO111MODULE on    
 go env -w GOPROXY=https://goproxy.cn,direct
 SETX GOPROXY https://goproxy.cn,direct
```
- linux

```shell
  echo "export GO111MODULE=on" >> ~/.profile
  echo "export GOPROXY=https://goproxy.cn" >> ~/.profile
  source ~/.profile
```

3. 配置`config.yaml`文件中的数据库或其他环境
4. 运行项目,初次运行会拉取项目依赖，请等待拉取完成后运行
```shell
go run ./main.go
```
5. `build`为二进制运行,执行完成后，根目录会出现`dpj-admin-api`的二进制文件，
```shell
go build -o main-linux main.go
```
6. 部署与守护进程运行，将二进制文件与config.yaml放到服务器后，使用**Supervisor**守护进程运行，
```text
  
```


### 开发
 #### 配置读取
```golang
config.Get("配置项" ,"默认值")
```

#### 中间件
```golang
r.User(cors)
```

|中间件| 作用域 |
|----|----|
| cors| 跨域放行中间件|
| casbin| 权限校验中间件|
| jwt| token校验中间件|
| error | 异常处理中间件|

#### 二进制打包
```shell
go build -o main-linux main.go
```
