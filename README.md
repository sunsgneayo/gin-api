# gin-api

> 目前只适用于http项目支持

### 架构
- corn 异步定时任务
- base64Captcha 灵活验证码
- viper 配置服务
- casbin 权限服务
- redis 缓存服务
- rabbit 队列服务

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
