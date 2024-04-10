<p align="center">
  <img width="160px" src="http://cdn.nine1120.cn/logo-i.png" alt="sunsgne">
</p>

## 🐬 gin-fast-api

> 支持`http-server`,`rabbitMq`,`async-task`服务同时运行

### 基础架构 🥃
- ✅ corn 异步定时任务
- ✅ base64Captcha 验证码
- ✅ go_captcha 行为验证码
- ✅ viper 配置发现服务
- ✅ casbin 权限认证服务
- ✅ redis 缓存服务
- ✅ rabbit 队列服务
- ✅ http  客户端服务
- ✅ rabbit 消息队列
- ✅ go-resty HTTP客户端

### 快速开始 🏫

1. 下载并安装golang,[下载地址](https://go.dev/dl/)->选择合适的版本
2. 切换国内golang镜像源
   ```shell
    # windows 
    SETX GO111MODULE on    
    go env -w GOPROXY=https://goproxy.cn,direct
    SETX GOPROXY https://goproxy.cn,direct
    
    # linux
    echo "export GO111MODULE=on" >> ~/.profile
    echo "export GOPROXY=https://goproxy.cn" >> ~/.profile
    source ~/.profile
   ```
3. 配置`config.yaml`文件中的数据库或其他环境

4. 运行项目,初次运行会拉取项目依赖，请等待拉取完成后运行
   ```shell
   go run ./main.go
   ```
5. 配置打包环境为Linux
   ```shell
   set GOOS=linux
   or
   go env -w GOOS=linux
   ```

6. `build`为二进制运行,执行完成后，根目录会出现`dpj-admin-api`的二进制文件，
   ```shell
   set GOOS=linux
   
   go build -o dpj-admin-api main.go
   ```
7. 部署与守护进程运行，将二进制文件与config.yaml放到服务器后，使用**Supervisor**守护进程运行，
   ```text
   # 新建一个应用并设置一个名称，这里设置为 
   [program:dpj-admin-api]
   # 设置命令在指定的目录内执行
   directory=/www/wwwroot/dpj-admin.cqzln.top/
   # 这里为您要管理的项目的启动命令
   command=./dpj-admin-api
   # 以哪个用户来运行该进程
   user=root
   # supervisor 启动时自动该应用
   autostart=true
   # 进程退出后自动重启进程
   autorestart=true
   # 进程持续运行多久才认为是启动成功
   startsecs=1
   # 重试次数
   startretries=3
   # stderr 日志输出位置
   #stderr_logfile=/www/wwwroot/dpj-admin.cqzln.top/runtime/stderr.log
   # stdout 日志输出位置
   #stdout_logfile=/www/wwwroot/dpj-admin.cqzln.top/runtime/stdout.log
   ```


### 开发 🧑‍💻
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

#### 二进制打包 📦
```shell
go build -o main-linux main.go
```

#### 消息队列 🐇

##### 队列投递 🗳
```go
    // 创建 RabbitMQ 实例时将使用新的连接配置
	rabbitmq, _ := client.NewRabbitMQ("queueName")
	// 其他操作...
	rabbitmq.PublishSimple("Hello, RabbitMQ!")
	// 最后别忘了关闭连接
	defer rabbitmq.Destroy()
```

##### 队列消费 🆓
```go
	go func() {

		rabbitmq, _ := client.NewRabbitMQ("queueName")

		rabbitmq.ConsumeSimple(func(msg amqp.Delivery) {
			log.Printf("接收到消费数据: %s", msg.Body)
		})
	}()

```
