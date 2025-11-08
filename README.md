# Kitex-Nacos 微服务测试项目

这是一个使用 Go + Kitex 框架并通过 Nacos 实现服务注册与发现的微服务测试项目。

## 项目结构

```
kitex-nacos-test/
├── user-service/          # 用户服务
│   ├── main.go
│   └── handler.go
├── order-service/         # 订单服务
│   ├── main.go
│   └── handler.go
├── idl/                   # 接口定义文件
│   ├── user.thrift
│   └── order.thrift
├── common/                # 公共模块
│   └── nacos.go
├── kitex_gen/             # Kitex 生成代码
├── go.mod
└── README.md
```

## 功能说明

1. **user-service**：提供 GetUserInfo(userID) 接口，根据用户ID返回用户信息
2. **order-service**：调用 user-service 的接口来查询用户信息，并创建订单
3. 两个服务都注册到 Nacos 中，实现服务注册与发现

## 环境要求

- Go 1.16+
- Nacos 服务已启动并运行在默认端口 (8848)

## 依赖配置

```bash
go mod tidy
```

## 使用方法

### 1. 启动 user-service

```bash
cd user-service
go run .
```

### 2. 启动 order-service

```bash
cd order-service
go run .
```

### 3. 访问 Nacos 控制台

打开浏览器访问 Nacos 控制台：
- 地址：http://localhost:8848/nacos
- 用户名：nacos
- 密码：nacos

在服务列表中应该能看到注册的两个服务：
- UserService
- OrderService

### 4. 查看日志

服务启动后会在控制台输出日志，可以通过以下方式查看：
- 直接查看终端输出
- 在项目根目录查找生成的 `.log` 文件（如果有配置日志文件）

## 服务调用链路

1. 客户端调用 order-service 的 CreateOrder 接口
2. order-service 通过 Nacos 发现并调用 user-service 的 GetUserInfo 接口
3. user-service 返回用户信息给 order-service
4. order-service 创建订单并返回结果

## 验证服务调用

可以通过 curl 或其他 HTTP 客户端工具来测试服务调用：

```bash
curl -X POST http://localhost:9002/order/create -d '{"userID": 1, "product": "测试商品"}'
```

注意：实际的 API 调用方式取决于 Kitex 生成的客户端代码，这里只是一个示例。

## 项目配置

Nacos 配置在 [common/nacos.go](common/nacos.go) 文件中，默认配置如下：
- Nacos 地址：127.0.0.1
- Nacos 端口：8848

如需修改，请相应调整该文件中的配置。