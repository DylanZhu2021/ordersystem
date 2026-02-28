# 餐厅点单微信小程序系统

基于 Go + GoFrame v2 + PostgreSQL + Redis 的餐厅点单系统后端，配套微信小程序前端，采用前后端分离 RESTful 架构。

## 技术栈

| 层级 | 技术 |
|------|------|
| 后端框架 | GoFrame v2.7.4 |
| 语言 | Go 1.22 |
| 数据库 | PostgreSQL 16 |
| 缓存 | Redis 7 |
| 认证 | JWT (golang-jwt/v5) |
| 前端 | 微信小程序原生开发 |
| 容器化 | Docker Compose |

## 项目结构

```
ordersystem/                          # 后端项目
├── main.go                           # 入口文件
├── go.mod / go.sum
├── manifest/
│   ├── config/config.yaml            # 应用配置（数据库/Redis/JWT/微信）
│   ├── sql/001_init_schema.sql       # 数据库建表 + 种子数据
│   └── docker/docker-compose.yaml    # PostgreSQL + Redis 容器编排
├── api/v1/                           # 请求/响应结构体定义
│   ├── user.go                       # 用户相关
│   ├── product.go                    # 商品相关
│   ├── cart.go                       # 购物车
│   ├── order.go                      # 订单
│   ├── misc.go                       # 优惠券/评价/收藏/地址
│   └── admin/                        # 管理端
│       ├── admin.go                  # 管理员认证/商品管理
│       └── manage.go                 # 订单/优惠券/统计/分类/用户管理
├── internal/
│   ├── cmd/cmd.go                    # 服务启动 + 路由注册
│   ├── controller/                   # 控制器层（参数解析 → 调用 Service）
│   │   ├── auth.go                   # 用户认证
│   │   ├── product.go                # 商品/分类
│   │   ├── cart.go                   # 购物车
│   │   ├── order.go                  # 订单/支付
│   │   ├── misc.go                   # 优惠券/评价/收藏/地址
│   │   └── admin/admin.go           # 管理端全部控制器
│   ├── service/                      # 业务接口定义（依赖注入）
│   │   ├── service.go                # 用户端接口
│   │   └── admin.go                  # 管理端接口
│   ├── logic/                        # 业务逻辑实现
│   │   ├── auth/auth.go              # 微信登录 + 自动注册
│   │   ├── product/product.go        # 商品列表/详情/搜索
│   │   ├── category/category.go      # 分类列表
│   │   ├── cart/cart.go              # Redis 购物车
│   │   ├── order/order.go            # 下单（库存安全）+ 订单管理
│   │   ├── payment/payment.go        # 模拟支付 + 积分发放
│   │   ├── coupon/coupon.go          # 优惠券领取/使用
│   │   ├── review/review.go          # 评价
│   │   ├── favorite/favorite.go      # 收藏
│   │   ├── address/address.go        # 地址管理
│   │   └── admin/                    # 管理端逻辑
│   │       ├── auth.go               # 管理员登录（bcrypt）
│   │       ├── product.go            # 商品 CRUD
│   │       ├── order.go              # 订单管理/退款
│   │       └── manage.go             # 优惠券/分类/统计/用户管理
│   ├── model/entity/                 # 数据库实体结构体
│   ├── middleware/middleware.go       # JWT认证/CORS/日志/限流/响应包装
│   ├── consts/                       # 常量定义
│   │   ├── consts.go                 # 订单状态/配送方式/优惠券类型/会员等级
│   │   ├── error_code.go             # 业务错误码
│   │   └── cache_key.go              # Redis Key 模板
│   └── packed/packed.go              # GoFrame 资源打包
└── utility/                          # 工具函数
    ├── response.go                   # 统一 JSON 响应（Success/Error/SuccessPage）
    ├── jwt.go                        # JWT 生成/解析
    ├── wechat.go                     # 微信 code2session
    ├── snowflake.go                  # 订单号/支付号生成
    └── idempotent.go                 # Redis 幂等性检查

ordersystem-miniapp/                  # 前端小程序项目
├── app.js / app.json / app.wxss     # 全局配置
├── project.config.json               # 微信开发者工具配置
├── utils/
│   ├── request.js                    # HTTP 请求封装（自动带 JWT）
│   ├── api.js                        # 全部 API 接口定义
│   └── util.js                       # 工具函数
└── pages/                            # 13 个页面
    ├── index/                        # 首页（左分类 + 右商品列表）
    ├── detail/                       # 商品详情 + 搜索
    ├── cart/                         # 购物车
    ├── order-confirm/                # 下单确认
    ├── order-list/                   # 订单列表
    ├── order-detail/                 # 订单详情
    ├── mine/                         # 个人中心
    ├── login/                        # 微信登录
    ├── address/ + address-edit/      # 地址管理
    ├── coupon/                       # 优惠券
    ├── favorite/                     # 收藏
    └── review/                       # 评价
```

## 系统架构设计

```
┌──────────────────────────────────────────────────────────┐
│                   微信小程序 (前端)                         │
│         index / detail / cart / order / mine              │
└──────────────────────┬───────────────────────────────────┘
                       │ HTTPS / JSON
┌──────────────────────▼───────────────────────────────────┐
│                  GoFrame HTTP Server (:8080)              │
│  ┌─────────────────────────────────────────────────────┐ │
│  │ Middleware: CORS → ResponseHandler → RequestLog      │ │
│  │            → JWTAuth (用户) / AdminAuth (管理员)      │ │
│  └─────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────┐ │
│  │ Controller: 参数校验 → 调用 Service                   │ │
│  └─────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────┐ │
│  │ Service (接口) → Logic (实现): 核心业务逻辑            │ │
│  └─────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────┐ │
│  │ ORM (g.DB) + Raw SQL: 数据访问                       │ │
│  └─────────────────────────────────────────────────────┘ │
└──────────┬──────────────────────────────┬────────────────┘
           │                              │
    ┌──────▼──────┐                ┌──────▼──────┐
    │ PostgreSQL  │                │    Redis    │
    │  16-alpine  │                │  7-alpine   │
    │  :5432      │                │  :6379      │
    │             │                │             │
    │ 15张业务表   │                │ 购物车 HASH  │
    │ 索引+约束    │                │ 幂等性 SETNX │
    │ 事务保证     │                │ 缓存 Key     │
    └─────────────┘                └─────────────┘
```

### 分层职责

| 层级 | 目录 | 职责 |
|------|------|------|
| API 定义 | `api/v1/` | 请求/响应结构体，参数校验规则 |
| 控制器 | `internal/controller/` | 解析参数，调用 Service，返回响应 |
| 服务接口 | `internal/service/` | 定义业务接口（依赖注入） |
| 业务逻辑 | `internal/logic/` | 实现业务逻辑，通过 `init()` 注册到 Service |
| 中间件 | `internal/middleware/` | JWT 认证、CORS、日志、限流、响应包装 |
| 数据实体 | `internal/model/entity/` | 数据库表对应的 Go 结构体 |
| 工具函数 | `utility/` | JWT、微信API、雪花ID、幂等、响应封装 |

---

## 快速启动

### 环境要求

- Go 1.22+
- Docker Desktop（运行 PostgreSQL + Redis）
- 微信开发者工具（运行小程序前端）

### 1. 启动数据库和 Redis

```bash
cd manifest/docker
docker-compose up -d
```

首次启动时，`manifest/sql/001_init_schema.sql` 会自动执行，完成：
- 创建全部 15 张业务表 + 索引
- 插入默认角色（店长、员工）
- 插入默认管理员（admin / admin123）
- 插入默认商品分类（热销推荐、主食、小吃、饮品、甜点）

验证数据库是否就绪：

```bash
docker exec -it ordersystem-pg psql -U ordersystem -c "\dt"
```

### 2. 启动后端服务

```bash
# Windows 环境设置 Go 路径（如果默认 GOROOT 不对）
set GOROOT=C:\Program Files\go1.22.4.windows-amd64\go
set PATH=%GOROOT%\bin;%PATH%

# 安装依赖（首次）
go mod tidy

# 启动服务
go run main.go
```

服务启动后输出：

```
  SERVER  | DOMAIN  | ADDRESS | METHOD |         ROUTE
----------|---------|---------|--------|------------------------
  default | default | :8080   | ALL    | /api/v1/*
  default | default | :8080   | ALL    | /api/admin/*
```

### 3. 启动前端小程序

1. 打开微信开发者工具
2. 导入项目 → 选择 `ordersystem-miniapp` 目录
3. AppID 填测试号或你自己的
4. 进入「详情 → 本地设置」→ 勾选「不校验合法域名」
5. 编译运行

### 4. 验证接口

```bash
# 获取分类列表（无需认证）
curl http://localhost:8080/api/v1/categories

# 管理员登录
curl -X POST http://localhost:8080/api/admin/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'

# 用返回的 token 添加商品
curl -X POST http://localhost:8080/api/admin/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "categoryId": 2,
    "name": "宫保鸡丁",
    "description": "经典川菜，鸡肉嫩滑，花生酥脆",
    "price": 38.00,
    "originalPrice": 48.00,
    "stock": 100,
    "status": 1
  }'
```

---

## 配置说明

配置文件位于 `manifest/config/config.yaml`：

```yaml
server:
  address: ":8080"              # 服务端口
  openapiPath: "/api.json"      # OpenAPI 文档路径
  swaggerPath: "/swagger"       # Swagger UI 路径

database:
  default:
    type: "pgsql"               # 数据库类型
    host: "127.0.0.1"
    port: "5432"
    user: "ordersystem"
    pass: "ordersystem123"
    name: "ordersystem"
    maxIdle: 10                 # 最大空闲连接
    maxOpen: 100                # 最大打开连接

redis:
  default:
    address: "127.0.0.1:6379"
    db: 0

app:
  jwt:
    secret: "your-jwt-secret-change-in-production"  # ⚠️ 生产环境必须修改
    expire: 86400               # 用户 token 有效期（24小时）
    adminExpire: 7200           # 管理员 token 有效期（2小时）
  wechat:
    appId: "wx_your_app_id"     # ⚠️ 替换为真实 AppID
    appSecret: "your_app_secret" # ⚠️ 替换为真实 AppSecret
  rateLimit:
    requests: 100               # 每窗口最大请求数
    window: 60                  # 窗口时间（秒）
  points:
    orderRate: 1                # 每消费1元获得1积分
    reviewReward: 10            # 每次评价奖励10积分
  membership:                   # 会员等级积分门槛
    silverThreshold: 1000
    goldThreshold: 5000
    platinumThreshold: 20000
```

### 生产环境必改项

| 配置项 | 说明 |
|--------|------|
| `app.jwt.secret` | JWT 签名密钥，必须改为随机强密码 |
| `app.wechat.appId` | 微信小程序 AppID |
| `app.wechat.appSecret` | 微信小程序 AppSecret |
| `database.default.pass` | 数据库密码 |
| `server.address` | 生产环境建议通过 Nginx 反向代理 |

### 前端配置

小程序的后端地址在 `ordersystem-miniapp/app.js` 中：

```javascript
globalData: {
  baseUrl: 'http://localhost:8080'  // 改为你的后端地址
}
```

---

## 数据库设计

### ER 关系概览

```
user ──1:N──> address
user ──1:N──> order ──1:N──> order_item ──N:1──> product
user ──1:N──> user_coupon ──N:1──> coupon
user ──1:N──> favorite ──N:1──> product
user ──1:N──> review ──N:1──> product
user ──1:N──> points_log
order ──1:N──> payment_log
product ──N:1──> category
product ──1:N──> product_spec
admin_user ──N:1──> role
```

### 核心表说明

| 表名 | 说明 | 关键字段 |
|------|------|----------|
| `user` | 用户表 | openid(唯一), points, total_points, member_level |
| `product` | 商品表 | stock(库存), sales(销量), is_hot, is_recommend, status |
| `product_spec` | 商品规格 | spec_name(口味/大小), price_diff(价格差) |
| `order` | 订单表 | order_no(唯一), status(状态机), idempotency_key(幂等) |
| `order_item` | 订单明细 | 商品快照（下单时冻结价格和名称） |
| `coupon` | 优惠券模板 | type(满减/折扣/无门槛), claimed_count(原子计数) |
| `user_coupon` | 用户优惠券 | status(未用/已用/过期) |
| `payment_log` | 支付记录 | transaction_no(唯一), status |
| `review` | 评价 | rating(1-5), reply(商家回复) |
| `points_log` | 积分流水 | change(正=获得/负=消费), balance |

### 订单状态流转

```
                    ┌──── 超时/用户取消 ────→ 已取消(5)
                    │
待支付(0) ──支付──→ 已支付(1) ──接单──→ 制作中(2) ──完成──→ 待取餐(3) ──确认──→ 已完成(4)
                    │
                    └──── 申请退款 ────→ 已退款(6)
```

---

## 接口设计

### 统一响应格式

```json
{
  "code": 0,
  "message": "success",
  "data": { ... }
}
```

### 错误码

| 错误码 | 含义 |
|--------|------|
| 0 | 成功 |
| 1001 | 参数错误 |
| 1002 | 未登录 / Token 过期 |
| 1003 | 无权限 |
| 1004 | 资源不存在 |
| 1005 | 重复操作 |
| 2001 | 库存不足 |
| 2002 | 下单失败 |
| 2003 | 支付失败 |
| 2004 | 优惠券不可用 |
| 2005 | 优惠券已领取 |
| 3001 | 请求过于频繁 |
| 4001 | 微信认证失败 |
| 5000 | 服务器内部错误 |

### 用户端接口 `/api/v1`

```
POST   /auth/wechat-login          微信登录
GET    /user/info                   获取用户信息
PUT    /user/update                 更新用户信息

GET    /categories                  分类列表（无需认证）
GET    /products                    商品列表（无需认证）
GET    /products/:id                商品详情（无需认证）
GET    /products/search             搜索商品（无需认证）

GET    /cart                        购物车列表
POST   /cart                        加入购物车
PUT    /cart/:productId             更新数量
DELETE /cart/:productId             删除商品
DELETE /cart/clear                  清空购物车

POST   /orders                      创建订单
GET    /orders                      订单列表
GET    /orders/:id                  订单详情
POST   /orders/:id/cancel           取消订单
POST   /orders/:id/refund           申请退款

POST   /payments/pay                模拟支付

GET    /coupons                     可领取优惠券
POST   /coupons/:id/claim           领取优惠券
GET    /coupons/my                  我的优惠券

POST   /reviews                     创建评价
GET    /reviews                     评价列表

POST   /favorites                   添加收藏
DELETE /favorites/:productId        取消收藏
GET    /favorites                   收藏列表

GET    /addresses                   地址列表
POST   /addresses                   创建地址
PUT    /addresses/:id               更新地址
DELETE /addresses/:id               删除地址
```

### 管理端接口 `/api/admin`

```
POST   /auth/login                  管理员登录
GET    /auth/info                   管理员信息

POST   /products                    创建商品
PUT    /products/:id                更新商品
DELETE /products/:id                删除商品
PUT    /products/:id/status         上架/下架

POST   /categories                  创建分类
PUT    /categories/:id              更新分类
DELETE /categories/:id              删除分类

GET    /orders                      订单列表
PUT    /orders/:id/status           更新订单状态
POST   /orders/:id/refund           处理退款

POST   /coupons                     创建优惠券
PUT    /coupons/:id                 更新优惠券
DELETE /coupons/:id                 删除优惠券

GET    /stats/dashboard             数据看板
GET    /stats/sales                 销售统计
GET    /stats/hot-products          热销排行

GET    /users                       用户列表
```

---

## 核心业务流程

### 下单流程（含并发安全）

```
用户点击「提交订单」
       │
       ▼
┌─ 1. 幂等性检查 ─────────────────────────────────────────┐
│  Redis SETNX idempotent:{key} → 5分钟过期                │
│  重复请求直接拒绝                                         │
└──────────────────────────────────────────────────────────┘
       │
       ▼
┌─ 2. 读取购物车 ─────────────────────────────────────────┐
│  从 Redis HASH cart:{userId} 获取商品列表                 │
│  校验购物车非空                                           │
└──────────────────────────────────────────────────────────┘
       │
       ▼
┌─ 3. 开启数据库事务 ─────────────────────────────────────┐
│                                                          │
│  ┌─ 3a. 原子扣减库存（防超卖核心）──────────────────┐    │
│  │  UPDATE product                                  │    │
│  │  SET stock = stock - ?, sales = sales + ?        │    │
│  │  WHERE id = ? AND stock >= ? AND status = 1      │    │
│  │                                                  │    │
│  │  affected == 0 → 回滚事务，返回「库存不足」       │    │
│  └──────────────────────────────────────────────────┘    │
│                                                          │
│  ┌─ 3b. 快照商品信息 ──────────────────────────────┐    │
│  │  记录下单时的商品名称、价格、规格                  │    │
│  │  后续商品改价不影响已有订单                        │    │
│  └──────────────────────────────────────────────────┘    │
│                                                          │
│  ┌─ 3c. 应用优惠券（如有）─────────────────────────┐    │
│  │  校验优惠券有效期、最低消费                       │    │
│  │  计算折扣金额                                     │    │
│  │  标记 user_coupon 为已使用                        │    │
│  └──────────────────────────────────────────────────┘    │
│                                                          │
│  ┌─ 3d. 写入订单 + 订单明细 ──────────────────────┐    │
│  │  生成雪花订单号 ORD + 时间戳 + 序列号             │    │
│  │  INSERT order + INSERT order_item                │    │
│  └──────────────────────────────────────────────────┘    │
│                                                          │
│  提交事务                                                │
└──────────────────────────────────────────────────────────┘
       │
       ▼
┌─ 4. 清空购物车 ─────────────────────────────────────────┐
│  删除 Redis HASH cart:{userId}                           │
└──────────────────────────────────────────────────────────┘
       │
       ▼
  返回订单信息（orderId, orderNo, payAmount）
```

### 库存防超卖方案

采用数据库行级原子操作，核心 SQL：

```sql
UPDATE product
SET stock = stock - $1, sales = sales + $1, updated_at = NOW()
WHERE id = $2 AND stock >= $1 AND status = 1
```

- `WHERE stock >= $1` 保证不会扣成负数
- 整个操作在事务内，任一商品库存不足则全部回滚
- 取消订单时自动恢复库存

### 优惠券防超领方案

```sql
-- 原子扣减可领数量
UPDATE coupon
SET claimed_count = claimed_count + 1
WHERE id = $1 AND claimed_count < total_count AND status = 1
  AND start_time <= NOW() AND end_time >= NOW()
```

事务内还会检查 per_user_limit，超限则回滚扣减。

### 订单号生成

使用雪花算法变体，格式：`ORD` + 年月日时分秒(14位) + 原子序列号(6位)

```go
// 示例: ORD20260228143052000001
func GenerateOrderNo() string {
    seq := atomic.AddInt64(&orderSeq, 1) % 1000000
    return "ORD" + time.Now().Format("20060102150405") + fmt.Sprintf("%06d", seq)
}
```

### 购物车 Redis 存储

```
Key:   cart:{userId}          # Redis HASH
Field: {productId}:{specId}   # 商品+规格组合
Value: JSON { productId, specId, quantity }
TTL:   30天
```

---

## 认证机制

### 用户端：微信登录

```
小程序 wx.login() → code
       │
       ▼
后端 POST /api/v1/auth/wechat-login { code }
       │
       ▼
调用微信 API: https://api.weixin.qq.com/sns/jscode2session
       │
       ▼
获取 openid → INSERT ... ON CONFLICT(openid) DO UPDATE（自动注册）
       │
       ▼
生成 JWT Token（有效期 24 小时）→ 返回给小程序
```

### 管理端：账号密码

```
POST /api/admin/auth/login { username, password }
       │
       ▼
查询 admin_user → bcrypt.CompareHashAndPassword 校验密码
       │
       ▼
生成 JWT Token（有效期 2 小时，userType = "admin"）
```

### JWT 结构

```go
type Claims struct {
    UserId   int64  `json:"userId"`
    UserType string `json:"userType"`  // "user" 或 "admin"
    jwt.RegisteredClaims
}
```

中间件通过 `UserType` 区分用户端和管理端请求。

---

## Redis 缓存策略

| Key 模式 | 用途 | TTL |
|----------|------|-----|
| `cart:{userId}` | 购物车 HASH | 30天 |
| `idempotent:{key}` | 下单幂等性 | 5分钟 |
| `product:{id}` | 商品详情缓存 | 5分钟 |
| `categories` | 分类列表缓存 | 30分钟 |
| `rate_limit:{ip}` | IP 限流计数 | 1分钟 |

---

## 部署说明

### 开发环境

```bash
# 1. 启动基础设施
cd manifest/docker && docker-compose up -d

# 2. 启动后端
go run main.go

# 3. 微信开发者工具导入 ordersystem-miniapp
```

### 生产环境

```bash
# 编译
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ordersystem main.go

# 运行
./ordersystem
```

建议配合 Nginx 反向代理：

```nginx
server {
    listen 443 ssl;
    server_name api.yourrestaurant.com;

    ssl_certificate     /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;

    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

小程序上线前需要在微信公众平台配置合法域名。

---

## 默认账号

| 角色 | 用户名 | 密码 | 说明 |
|------|--------|------|------|
| 管理员 | admin | admin123 | 拥有全部权限（店长角色） |

用户端通过微信授权自动注册，无需手动创建账号。
