-- ============================================================
-- 餐厅点单系统 - 数据库初始化脚本
-- PostgreSQL 16
-- ============================================================

-- 用户表
CREATE TABLE IF NOT EXISTS "user" (
    id            BIGSERIAL PRIMARY KEY,
    openid        VARCHAR(64) UNIQUE NOT NULL,
    union_id      VARCHAR(64),
    nickname      VARCHAR(64) DEFAULT '',
    avatar_url    VARCHAR(512) DEFAULT '',
    phone         VARCHAR(20) DEFAULT '',
    gender        SMALLINT DEFAULT 0,
    points        INT DEFAULT 0,
    total_points  INT DEFAULT 0,
    member_level  SMALLINT DEFAULT 0,
    status        SMALLINT DEFAULT 1,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_user_phone ON "user"(phone);

-- 用户地址表
CREATE TABLE IF NOT EXISTS address (
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT NOT NULL,
    name        VARCHAR(50) NOT NULL,
    phone       VARCHAR(20) NOT NULL,
    province    VARCHAR(50) DEFAULT '',
    city        VARCHAR(50) DEFAULT '',
    district    VARCHAR(50) DEFAULT '',
    detail      VARCHAR(200) NOT NULL,
    is_default  SMALLINT DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_address_user ON address(user_id);

-- 管理员表
CREATE TABLE IF NOT EXISTS admin_user (
    id            BIGSERIAL PRIMARY KEY,
    username      VARCHAR(64) UNIQUE NOT NULL,
    password_hash VARCHAR(128) NOT NULL,
    real_name     VARCHAR(64) DEFAULT '',
    phone         VARCHAR(20) DEFAULT '',
    role_id       INT DEFAULT 0,
    status        SMALLINT DEFAULT 1,
    last_login_at TIMESTAMPTZ,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 角色表
CREATE TABLE IF NOT EXISTS role (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(64) UNIQUE NOT NULL,
    description VARCHAR(256) DEFAULT '',
    permissions JSONB NOT NULL DEFAULT '[]',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 商品分类表
CREATE TABLE IF NOT EXISTS category (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(64) NOT NULL,
    icon_url   VARCHAR(512) DEFAULT '',
    sort_order INT DEFAULT 0,
    status     SMALLINT DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 商品表
CREATE TABLE IF NOT EXISTS product (
    id             BIGSERIAL PRIMARY KEY,
    category_id    INT NOT NULL,
    name           VARCHAR(128) NOT NULL,
    description    TEXT DEFAULT '',
    price          NUMERIC(10,2) NOT NULL,
    original_price NUMERIC(10,2),
    image_url      VARCHAR(512) DEFAULT '',
    images         JSONB DEFAULT '[]',
    stock          INT NOT NULL DEFAULT 0,
    sales          INT DEFAULT 0,
    is_hot         SMALLINT DEFAULT 0,
    is_recommend   SMALLINT DEFAULT 0,
    status         SMALLINT DEFAULT 1,
    sort_order     INT DEFAULT 0,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_product_category ON product(category_id);
CREATE INDEX IF NOT EXISTS idx_product_status ON product(status);
CREATE INDEX IF NOT EXISTS idx_product_sales ON product(sales DESC);

-- 商品规格表
CREATE TABLE IF NOT EXISTS product_spec (
    id          BIGSERIAL PRIMARY KEY,
    product_id  BIGINT NOT NULL,
    spec_name   VARCHAR(50) NOT NULL,
    spec_value  VARCHAR(50) NOT NULL,
    price_diff  NUMERIC(10,2) DEFAULT 0,
    stock       INT DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_spec_product ON product_spec(product_id);

-- 收藏表
CREATE TABLE IF NOT EXISTS favorite (
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, product_id)
);

-- 订单表
CREATE TABLE IF NOT EXISTS "order" (
    id              BIGSERIAL PRIMARY KEY,
    order_no        VARCHAR(32) UNIQUE NOT NULL,
    user_id         BIGINT NOT NULL,
    total_amount    NUMERIC(10,2) NOT NULL,
    discount_amount NUMERIC(10,2) DEFAULT 0,
    pay_amount      NUMERIC(10,2) NOT NULL,
    coupon_id       BIGINT DEFAULT 0,
    delivery_type   SMALLINT NOT NULL DEFAULT 1,
    table_no        VARCHAR(20) DEFAULT '',
    address_id      BIGINT DEFAULT 0,
    contact_name    VARCHAR(50) DEFAULT '',
    contact_phone   VARCHAR(20) DEFAULT '',
    remark          VARCHAR(256) DEFAULT '',
    status          SMALLINT NOT NULL DEFAULT 0,
    idempotency_key VARCHAR(64) UNIQUE,
    paid_at         TIMESTAMPTZ,
    completed_at    TIMESTAMPTZ,
    cancelled_at    TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_order_user ON "order"(user_id);
CREATE INDEX IF NOT EXISTS idx_order_status ON "order"(status);
CREATE INDEX IF NOT EXISTS idx_order_created ON "order"(created_at DESC);

-- 订单明细表
CREATE TABLE IF NOT EXISTS order_item (
    id            BIGSERIAL PRIMARY KEY,
    order_id      BIGINT NOT NULL,
    product_id    BIGINT NOT NULL,
    product_name  VARCHAR(128) NOT NULL,
    product_image VARCHAR(512) DEFAULT '',
    spec_id       BIGINT DEFAULT 0,
    spec_info     VARCHAR(200) DEFAULT '',
    price         NUMERIC(10,2) NOT NULL,
    quantity      INT NOT NULL,
    total_amount  NUMERIC(10,2) NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_order_item_order ON order_item(order_id);

-- 优惠券表
CREATE TABLE IF NOT EXISTS coupon (
    id             BIGSERIAL PRIMARY KEY,
    name           VARCHAR(128) NOT NULL,
    type           SMALLINT NOT NULL,
    discount_value NUMERIC(10,2) NOT NULL,
    min_amount     NUMERIC(10,2) DEFAULT 0,
    total_count    INT NOT NULL,
    claimed_count  INT DEFAULT 0,
    per_user_limit INT DEFAULT 1,
    start_time     TIMESTAMPTZ NOT NULL,
    end_time       TIMESTAMPTZ NOT NULL,
    status         SMALLINT DEFAULT 1,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 用户优惠券表
CREATE TABLE IF NOT EXISTS user_coupon (
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT NOT NULL,
    coupon_id  BIGINT NOT NULL,
    order_id   BIGINT DEFAULT 0,
    status     SMALLINT DEFAULT 0,
    claimed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    used_at    TIMESTAMPTZ
);
CREATE INDEX IF NOT EXISTS idx_ucoupon_user ON user_coupon(user_id);
CREATE INDEX IF NOT EXISTS idx_ucoupon_coupon ON user_coupon(coupon_id);

-- 支付记录表
CREATE TABLE IF NOT EXISTS payment_log (
    id             BIGSERIAL PRIMARY KEY,
    order_id       BIGINT NOT NULL,
    transaction_no VARCHAR(64) UNIQUE NOT NULL,
    amount         NUMERIC(10,2) NOT NULL,
    method         VARCHAR(32) DEFAULT 'wechat_simulated',
    status         SMALLINT NOT NULL DEFAULT 0,
    notify_data    TEXT DEFAULT '',
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_payment_order ON payment_log(order_id);

-- 评价表
CREATE TABLE IF NOT EXISTS review (
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    order_id   BIGINT NOT NULL,
    rating     SMALLINT NOT NULL CHECK (rating BETWEEN 1 AND 5),
    content    TEXT DEFAULT '',
    images     JSONB DEFAULT '[]',
    reply      TEXT DEFAULT '',
    status     SMALLINT DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_review_product ON review(product_id);
CREATE INDEX IF NOT EXISTS idx_review_user ON review(user_id);

-- 积分记录表
CREATE TABLE IF NOT EXISTS points_log (
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT NOT NULL,
    change     INT NOT NULL,
    balance    INT NOT NULL,
    type       SMALLINT NOT NULL,
    ref_id     BIGINT DEFAULT 0,
    remark     VARCHAR(256) DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_points_user ON points_log(user_id);

-- 操作日志表
CREATE TABLE IF NOT EXISTS operation_log (
    id         BIGSERIAL PRIMARY KEY,
    admin_id   BIGINT DEFAULT 0,
    module     VARCHAR(50) DEFAULT '',
    action     VARCHAR(50) DEFAULT '',
    content    TEXT DEFAULT '',
    ip         VARCHAR(50) DEFAULT '',
    user_agent TEXT DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_oplog_admin ON operation_log(admin_id);
CREATE INDEX IF NOT EXISTS idx_oplog_created ON operation_log(created_at DESC);

-- ============================================================
-- 初始数据
-- ============================================================
INSERT INTO role (name, description, permissions) VALUES
('店长', '店铺管理员，拥有全部权限', '["*"]'),
('员工', '普通员工', '["order:view","product:view","stats:view"]')
ON CONFLICT (name) DO NOTHING;

INSERT INTO admin_user (username, password_hash, real_name, role_id) VALUES
('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '系统管理员', 1)
ON CONFLICT (username) DO NOTHING;

INSERT INTO category (name, sort_order) VALUES
('热销推荐', 1), ('主食', 2), ('小吃', 3), ('饮品', 4), ('甜点', 5)
ON CONFLICT DO NOTHING;
