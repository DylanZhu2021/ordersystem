package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"ordersystem/internal/controller"
	adminCtrl "ordersystem/internal/controller/admin"
	"ordersystem/internal/middleware"
	"ordersystem/utility"
)

var Main = gcmd.Command{
	Name:  "main",
	Usage: "main",
	Brief: "餐厅点单系统后端服务",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		s := g.Server()

		// ==================== 用户端接口 ====================
		s.Group("/api/v1", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.CORS)

			// 健康检查
			group.GET("/health", func(r *ghttp.Request) {
				utility.Success(r, g.Map{"status": "ok"})
			})

			// 公开接口（无需登录）
			group.Group("/auth", func(group *ghttp.RouterGroup) {
				group.POST("/wechat-login", controller.Auth.WechatLogin)
			})
			group.Group("/categories", func(group *ghttp.RouterGroup) {
				group.GET("/", controller.Category.List)
			})
			group.Group("/products", func(group *ghttp.RouterGroup) {
				group.GET("/", controller.Product.List)
				group.GET("/search", controller.Product.Search)
				group.GET("/:id", controller.Product.Detail)
			})
			group.Group("/reviews", func(group *ghttp.RouterGroup) {
				group.GET("/", controller.Review.List)
			})

			// 需要登录的接口
			group.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.JWTAuth)

				// 用户信息
				group.GET("/user/info", controller.Auth.GetUserInfo)
				group.PUT("/user/update", controller.Auth.UpdateUser)

				// 地址管理
				group.GET("/addresses", controller.Address.List)
				group.POST("/addresses", controller.Address.Create)
				group.PUT("/addresses", controller.Address.Update)
				group.DELETE("/addresses/:id", controller.Address.Delete)

				// 购物车
				group.GET("/cart", controller.Cart.List)
				group.POST("/cart", controller.Cart.Add)
				group.PUT("/cart", controller.Cart.Update)
				group.DELETE("/cart/item", controller.Cart.Delete)
				group.DELETE("/cart/clear", controller.Cart.Clear)

				// 订单
				group.POST("/orders", controller.Order.Create)
				group.GET("/orders", controller.Order.List)
				group.GET("/orders/:id", controller.Order.Detail)
				group.POST("/orders/:id/cancel", controller.Order.Cancel)
				group.POST("/orders/:id/refund", controller.Order.Refund)

				// 支付
				group.POST("/payments/simulate", controller.Payment.SimulatePay)

				// 优惠券
				group.GET("/coupons", controller.Coupon.List)
				group.POST("/coupons/claim", controller.Coupon.Claim)
				group.GET("/coupons/my", controller.Coupon.MyCoupons)

				// 评价
				group.POST("/reviews", controller.Review.Create)

				// 收藏
				group.GET("/favorites", controller.Favorite.List)
				group.POST("/favorites", controller.Favorite.Add)
				group.DELETE("/favorites", controller.Favorite.Delete)
			})
		})

		// ==================== 管理端接口 ====================
		s.Group("/api/admin", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.CORS)

			// 管理员登录（公开）
			group.POST("/auth/login", adminCtrl.Auth.Login)

			// 需要管理员登录的接口
			group.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.AdminAuth)

				group.GET("/auth/info", adminCtrl.Auth.Info)

				// 商品管理
				group.POST("/products", adminCtrl.Product.Create)
				group.PUT("/products", adminCtrl.Product.Update)
				group.DELETE("/products/:id", adminCtrl.Product.Delete)
				group.PUT("/products/status", adminCtrl.Product.UpdateStatus)

				// 分类管理
				group.POST("/categories", adminCtrl.Category.Create)
				group.PUT("/categories", adminCtrl.Category.Update)
				group.DELETE("/categories/:id", adminCtrl.Category.Delete)

				// 订单管理
				group.GET("/orders", adminCtrl.Order.List)
				group.PUT("/orders/status", adminCtrl.Order.UpdateStatus)
				group.POST("/orders/refund", adminCtrl.Order.Refund)

				// 优惠券管理
				group.POST("/coupons", adminCtrl.Coupon.Create)
				group.PUT("/coupons", adminCtrl.Coupon.Update)
				group.DELETE("/coupons/:id", adminCtrl.Coupon.Delete)

				// 数据统计
				group.GET("/stats/dashboard", adminCtrl.Stats.Dashboard)
				group.GET("/stats/sales", adminCtrl.Stats.SalesStats)
				group.GET("/stats/hot-products", adminCtrl.Stats.HotProducts)

				// 用户管理
				group.GET("/users", adminCtrl.User.List)
			})
		})

		s.Run()
		return nil
	},
}
