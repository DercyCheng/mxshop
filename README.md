# Shop 电商微服务系统

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version" />
  <img src="https://img.shields.io/badge/Docker-Compose-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
  <img src="https://img.shields.io/badge/gRPC-Microservices-2396F3?style=for-the-badge&logo=grpc&logoColor=white" alt="gRPC" />
  <img src="https://img.shields.io/badge/MySQL-Database-4479A1?style=for-the-badge&logo=mysql&logoColor=white" alt="MySQL" />
  <img src="https://img.shields.io/badge/Vue.js-3.x-4FC08D?style=for-the-badge&logo=vue.js&logoColor=white" alt="Vue.js" />
  <img src="https://img.shields.io/badge/Wire-DI-00BFFF?style=for-the-badge&logo=go&logoColor=white" alt="Wire DI" />
</p>

<p align="center">
  Shop 是一个基于 Go 语言和微服务架构的完整电商解决方案，采用前后端分离设计，集成了多种现代技术栈与最佳实践。
</p>

## 📑 目录

- [系统概述](#-系统概述)
- [架构设计](#-架构设计)
- [核心功能](#-核心功能)
- [技术栈](#-技术栈)
- [项目结构](#-项目结构)
- [快速开始](#-快速开始)
- [接口文档](#-接口文档)
- [性能测试](#-性能测试)
- [开发指南](#-开发指南)
- [系统文档](#-系统文档)
- [常见问题](#-常见问题)
- [贡献指南](#-贡献指南)
- [许可证](#-许可证)

## 🚀 系统概述

Shop 是一个面向中小型企业的电商系统解决方案，实现了用户、商品、订单、库存和支付等电商核心业务流程。系统基于微服务架构设计，各服务独立部署，确保了系统的高可用性、可扩展性和容错性。

### 业务流程图

```mermaid
graph TD
    A[用户浏览商品] --> B[添加购物车]
    B --> C[下单]
    C --> D{库存校验}
    D -->|库存充足| E[库存锁定]
    E --> F[创建订单]
    F --> G[支付]
    G -->|支付成功| H[发货]
    G -->|支付失败| I[取消订单]
    I --> J[释放库存]
    D -->|库存不足| K[下单失败]
```

## 🏗 架构设计

Shop 采用经典的微服务分层架构，确保了系统的模块化和可维护性。

### 系统架构图

![架构图](./image/README/1744869775075.png)

### 分层设计

| 层级         | 职责             | 组件                                                           |
| ------------ | ---------------- | -------------------------------------------------------------- |
| 基础设施层   | 提供基础服务支持 | MySQL、Redis、ElasticSearch、RocketMQ、Consul、Nacos           |
| 服务层 (SRV) | 实现核心业务逻辑 | 用户服务、商品服务、库存服务、订单服务、用户操作服务           |
| API 层 (Web) | 提供 HTTP 接口   | 用户 API、商品 API、订单 API、库存 API、用户操作 API、OSS 服务 |
| 前端层       | 用户界面展示     | 管理后台 (Vue 3 + Element Plus)                                |

### 服务通信

```mermaid
sequenceDiagram
    participant Client
    participant API Gateway
    participant User Service
    participant Order Service
    participant Inventory Service
    participant MQ
  
    Client->>API Gateway: 创建订单请求
    API Gateway->>Order Service: gRPC: 创建订单
    Order Service->>Inventory Service: gRPC: 锁定库存
    Inventory Service-->>Order Service: 库存锁定结果
    Order Service->>MQ: 发送订单创建事件
    Order Service-->>API Gateway: 返回订单结果
    API Gateway-->>Client: 订单创建成功
    MQ->>Inventory Service: 消费订单事件
```

## 🔥 核心功能

Shop 实现了电商系统所需的全部核心功能：

### 模块功能对比

| 功能模块 | 主要特性                   | 技术亮点                              |
| -------- | -------------------------- | ------------------------------------- |
| 用户服务 | 注册、登录、鉴权、个人中心 | JWT认证、RBAC权限模型、手机验证码登录 |
| 商品服务 | 商品管理、分类、品牌、属性 | ES全文检索、多级分类、规格管理        |
| 库存服务 | 库存管理、库存锁定/释放    | 分布式锁、乐观并发控制、库存预警      |
| 订单服务 | 购物车、订单管理、支付集成 | 分布式事务、状态机、超时取消          |
| 用户操作 | 收藏、地址管理、消息       | 地址结构化、收藏同步                  |
| OSS服务  | 文件上传、图片处理         | 对象存储、图片压缩、水印              |

## 💻 技术栈

Shop 采用现代化技术栈，兼顾性能与开发效率：

### 后端技术栈

```mermaid
graph TD
    Go[Go 语言] --> GRPC[gRPC]
    Go --> Gin[Gin Web 框架]
    Go --> GORM[GORM]
    Go --> Wire[Wire 依赖注入]
  
    subgraph 数据存储
        MySQL[(MySQL)]
        Redis[(Redis)]
        ES[(ElasticSearch)]
    end
  
    subgraph 中间件
        Consul[Consul 服务发现]
        Nacos[Nacos 配置中心]
        RocketMQ[RocketMQ 消息队列]
        Jaeger[Jaeger 链路追踪]
    end
  
    GRPC --> Consul
    GRPC --> Nacos
    Gin --> GRPC
    GORM --> MySQL
    Go --> Redis
    Go --> ES
    Go --> RocketMQ
    Go --> Jaeger
    Wire --> Go
```

### 前端技术栈

- **框架**: Vue 3 + TypeScript
- **组件库**: Element Plus
- **状态管理**: Pinia
- **路由**: Vue Router
- **HTTP 客户端**: Axios
- **构建工具**: Vite

## 📁 项目结构

```
shop/
├── docker-compose.yml  # Docker 部署配置
├── Dockerfile          # Docker 构建文件
├── README.md           # 项目说明
├── doc/                # 详细文档
├── image/              # 图片资源
├── scripts/            # 部署脚本
├── shop_api/           # API 层
│   ├── goods_web/      # 商品 API
│   ├── order_web/      # 订单 API
│   ├── oss_web/        # 对象存储 API
│   ├── user_web/       # 用户 API
│   └── userop_web/     # 用户操作 API
├── shop_srv/           # 服务层
│   ├── goods_srv/      # 商品服务
│   ├── inventory_srv/  # 库存服务
│   ├── order_srv/      # 订单服务
│   ├── user_srv/       # 用户服务
│   └── userop_srv/     # 用户操作服务
└── shop_stress/        # 压力测试工具
```

## 🚀 快速开始

### 前置要求

- Docker 和 Docker Compose
- Go 1.16+
- MySQL 8.0+
- 其他依赖组件（可通过 Docker Compose 自动部署）

### 使用 Docker Compose 一键部署

```bash
# 克隆仓库
git clone https://github.com/username/shop.git
cd shop

# 启动所有服务
python run.py
```

### 本地开发环境设置

```bash
# 启动基础设施
docker-compose up -d mysql nacos consul jaeger rocketmq

# 启动服务层
cd scripts
./start.sh start_srv

# 启动 API 层
./start.sh start_api
```

详细的部署文档请参考 [环境搭建指南](./doc/环境搭建.md)。

## 📝 接口文档

API 文档通过 Swagger UI 提供，启动服务后可访问：

- 用户服务: http://localhost:8021/swagger/index.html
- 商品服务: http://localhost:8022/swagger/index.html
- 订单服务: http://localhost:8023/swagger/index.html
- 用户操作: http://localhost:8024/swagger/index.html
- OSS 服务: http://localhost:8025/swagger/index.html

## 📊 性能测试

Shop 提供了完整的性能测试工具 (shop_stress)，支持对各微服务进行压力测试：

```bash
cd shop_stress
python stress_test.py -s user -d 30 -c 200 -t 8
```

详细的测试报告和使用方法请参考 [压力测试文档](./shop_stress/README.md)。

### 性能测试结果

| 服务名称 | QPS (1000并发) | 平均响应时间 | P99响应时间 |
| -------- | -------------- | ------------ | ----------- |
| 用户服务 | 5,000+         | < 20ms       | < 50ms      |
| 商品服务 | 3,000+         | < 30ms       | < 70ms      |
| 订单服务 | 2,000+         | < 40ms       | < 90ms      |
| 库存服务 | 4,000+         | < 25ms       | < 60ms      |

## 🔧 开发指南

### 添加新服务

1. 在 shop_srv 目录下创建新的服务目录
2. 编写 proto 文件定义服务接口
3. 生成 gRPC 代码
4. 定义领域接口
5. 实现具体的接口实现类
6. 使用 Wire 配置依赖注入
7. 在 shop_api 目录下创建对应的 API 层
8. 注册到服务发现

### 面向接口编程

Shop 采用面向接口编程的设计理念，主要优势包括：

- **松耦合**：通过接口将系统组件解耦，便于修改和扩展
- **可测试性**：便于编写单元测试，可以轻松模拟依赖项
- **代码复用**：接口允许多种实现方式，提高代码复用性
- **灵活性**：可以轻松替换具体实现，而不影响依赖该接口的代码

详细开发指南请参考 [开发文档](./doc/开发指南.md)。

## 📖 系统文档

Shop 项目提供了详细的系统文档，帮助开发者更好地理解和扩展系统：

### 架构文档

- [系统架构与数据流图](./doc/系统架构与数据流图.md) - 详细的系统组件关系图和数据流转过程
- [架构决策记录(ADR)](./doc/架构决策记录.md) - 记录重要架构决策的原因和影响

### API文档

- [API文档规范](./doc/API文档规范.md) - API开发和文档编写规范，包括Swagger使用指南
- Swagger UI接口文档（启动相应服务后访问）:
  - 用户服务: http://localhost:8021/swagger/index.html
  - 商品服务: http://localhost:8022/swagger/index.html
  - 订单服务: http://localhost:8023/swagger/index.html
  - 用户操作: http://localhost:8024/swagger/index.html
  - OSS 服务: http://localhost:8025/swagger/index.html

### 微服务文档

- [用户服务](./doc/用户服务.md) - 用户注册、登录、认证等功能
- [商品服务](./doc/商品服务.md) - 商品管理、分类、品牌等功能
- [订单服务](./doc/订单服务.md) - 订单处理、支付集成等
- [库存服务](./doc/库存服务.md) - 库存管理、锁定释放等
- [用户操作服务](./doc/用户操作服务.md) - 用户收藏、地址管理等

### 开发与部署

- [环境搭建](./doc/环境搭建.md) - 开发和生产环境的搭建指南
- [开发者指南](./doc/开发者指南.md) - 详细的代码规范、最佳实践和示例

## ❓ 常见问题

**Q: 如何修改服务配置？**

A: 所有服务配置都存储在 Nacos 配置中心，可以通过 Nacos 控制台 (http://localhost:8848/nacos) 进行修改。

**Q: 如何监控服务状态？**

A: 系统集成了 Jaeger 链路追踪，可以通过 Jaeger UI (http://localhost:16686) 查看服务调用情况。

**Q: 如何进行系统扩展？**

A: 得益于微服务架构，您可以轻松扩展或替换任何服务模块，只需确保遵循现有的接口定义。

更多问题请查阅 [常见问题解答](./doc/FAQ.md)。

## 👥 贡献指南

欢迎为 Shop 做出贡献！无论是提交 bug 报告、功能建议还是代码贡献，我们都非常感谢。

### 如何贡献

1. Fork 本仓库
2. 创建您的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交您的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 打开一个 Pull Request

### 代码风格

- Go 代码遵循 [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- 使用 `gofmt` 格式化代码
- 所有 API 必须有文档注释

## 📄 许可证

该项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 📮 联系我们

- 项目负责人: [dercy](mailto:dercyc@example.com)
- 项目仓库: [GitHub](https://github.com/username/shop)
- 问题反馈: [GitHub Issues](https://github.com/username/shop/issues)
