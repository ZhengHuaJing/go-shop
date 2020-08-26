package initialize

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/zhenghuajing/fresh_shop/docs"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/middleware"
	"github.com/zhenghuajing/fresh_shop/model"
	"github.com/zhenghuajing/fresh_shop/pkg/upload"
	"github.com/zhenghuajing/fresh_shop/pkg/util"
	v1 "github.com/zhenghuajing/fresh_shop/router/api/v1"
	"github.com/zhenghuajing/fresh_shop/service/api_service"
	"net/http"
	"strings"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	{
		// API在线文档
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 无需认证权限路由
	apiNotAuthV1 := r.Group("/api/v1")
	{
		// 上传文件
		apiNotAuthV1.POST("/upload", v1.UploadImage)
		// 图片访问
		apiNotAuthV1.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

		// 生成验证码
		apiNotAuthV1.GET("/captcha", v1.GenerateCaptcha)
		// 发送邮箱验证码
		apiNotAuthV1.POST("/verify_code", v1.GenerateVerifyCode)

		// 用户登录
		apiNotAuthV1.POST("/login", v1.Login)
		// 用户注册
		apiNotAuthV1.POST("/users", v1.AddUser)

		// 商品详情
		apiNotAuthV1.GET("/products/:id", v1.GetProduct)
		// 查询所有商品
		apiNotAuthV1.GET("/products", v1.GetAllProduct)

		// 商品分类详情
		apiNotAuthV1.GET("/categories/:id", v1.GetCategory)
		// 查询所有商品分类
		apiNotAuthV1.GET("/categories", v1.GetAllCategory)

		// 支付成功回调
		apiNotAuthV1.POST("/pay", v1.Pay)
	}

	// 需要认证权限路由
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.UserAuth())
	apiV1.Use(middleware.Casbin())
	{
		// 修改用户
		apiV1.PUT("/users/:id", v1.UpdateUser)
		// 用户详情
		apiV1.GET("/users/:id", middleware.IsSelfOperate(), v1.GetUser)
		// 删除用户
		apiV1.DELETE("/users/:id", v1.DeleteUser)
		// 用户列表
		apiV1.GET("/users", v1.GetAllUser)

		// API详情
		apiV1.GET("/apis/:id", v1.GetApi)
		// 查询所有API
		apiV1.GET("/apis", v1.GetAllApi)

		// 为用户添加角色
		apiV1.POST("/user/role", v1.AddRoleForUser)
		// 删除角色
		apiV1.DELETE("/roles/:role_name", v1.DeleteRole)
		// 添加/修改角色
		apiV1.POST("/roles", v1.UpdateRole)
		// 查询所有角色
		apiV1.GET("/roles", v1.GetAllRole)

		// 添加商品
		apiV1.POST("/products", v1.AddProduct)
		// 删除商品
		apiV1.DELETE("/products/:id", v1.DeleteProduct)
		// 修改商品
		apiV1.PUT("/products/:id", v1.UpdateProduct)

		// 添加商品分类
		apiV1.POST("/categories", v1.AddCategory)
		// 删除商品分类
		apiV1.DELETE("/categories/:id", v1.DeleteCategory)
		// 修改商品分类
		apiV1.PUT("/categories/:id", v1.UpdateCategory)

		// 添加订单
		apiV1.POST("/orders", v1.AddOrderInfo)
		// 删除订单
		apiV1.DELETE("/orders/:id", v1.DeleteOrderInfo)
		// 修改订单
		apiV1.PUT("/orders/:id", v1.UpdateOrderInfo)
		// 订单详情
		apiV1.GET("/orders/:id", v1.GetOrderInfo)
		// 查询所有订单
		apiV1.GET("/orders", v1.GetAllOrderInfo)

		// 添加分类广告
		apiV1.POST("/ads", v1.AddAd)
		// 删除分类广告
		apiV1.DELETE("/ads/:id", v1.DeleteAd)
		// 修改分类广告
		apiV1.PUT("/ads/:id", v1.UpdateAd)
		// 分类广告详情
		apiV1.GET("/ads/:id", v1.GetAd)
		// 查询所有分类广告
		apiV1.GET("/ads", v1.GetAllAd)

		// 添加用户收货地址
		apiV1.POST("/addresses", v1.AddAddress)
		// 删除用户收货地址
		apiV1.DELETE("/addresses/:id", v1.DeleteAddress)
		// 修改用户收货地址
		apiV1.PUT("/addresses/:id", v1.UpdateAddress)
		// 用户收货地址详情
		apiV1.GET("/addresses/:id", v1.GetAddress)
		// 查询所有用户收货地址
		apiV1.GET("/addresses", v1.GetAllAddress)

		// 添加品牌
		apiV1.POST("/brands", v1.AddBrand)
		// 删除品牌
		apiV1.DELETE("/brands/:id", v1.DeleteBrand)
		// 修改品牌
		apiV1.PUT("/brands/:id", v1.UpdateBrand)
		// 品牌详情
		apiV1.GET("/brands/:id", v1.GetBrand)
		// 查询所有品牌
		apiV1.GET("/brands", v1.GetAllBrand)

		// 添加用户收藏
		apiV1.POST("/collects", v1.AddCollect)
		// 删除用户收藏
		apiV1.DELETE("/collects/:id", v1.DeleteCollect)
		// 修改用户收藏
		apiV1.PUT("/collects/:id", v1.UpdateCollect)
		// 用户收藏详情
		apiV1.GET("/collects/:id", v1.GetCollect)
		// 查询所有用户收藏
		apiV1.GET("/collects", v1.GetAllCollect)

		// 添加热搜词
		apiV1.POST("/hot_searchs", v1.AddHotSearch)
		// 删除热搜词
		apiV1.DELETE("/hot_searchs/:id", v1.DeleteHotSearch)
		// 修改热搜词
		apiV1.PUT("/hot_searchs/:id", v1.UpdateHotSearch)
		// 热搜词详情
		apiV1.GET("/hot_searchs/:id", v1.GetHotSearch)
		// 查询所有热搜词
		apiV1.GET("/hot_searchs", v1.GetAllHotSearch)

		// 添加首页轮播
		apiV1.POST("/index_banners", v1.AddIndexBanner)
		// 删除首页轮播
		apiV1.DELETE("/index_banners/:id", v1.DeleteIndexBanner)
		// 修改首页轮播
		apiV1.PUT("/index_banners/:id", v1.UpdateIndexBanner)
		// 首页轮播详情
		apiV1.GET("/index_banners/:id", v1.GetIndexBanner)
		// 查询所有首页轮播
		apiV1.GET("/index_banners", v1.GetAllIndexBanner)

		// 添加用户留言
		apiV1.POST("/leaving_messages", v1.AddLeavingMessage)
		// 删除用户留言
		apiV1.DELETE("/leaving_messages/:id", v1.DeleteLeavingMessage)
		// 修改用户留言
		apiV1.PUT("/leaving_messages/:id", v1.UpdateLeavingMessage)
		// 用户留言详情
		apiV1.GET("/leaving_messages/:id", v1.GetLeavingMessage)
		// 查询所有用户留言
		apiV1.GET("/leaving_messages", v1.GetAllLeavingMessage)

		// 添加购物车
		apiV1.POST("/shopping_carts", v1.AddShoppingCart)
		// 删除购物车
		apiV1.DELETE("/shopping_carts/:id", v1.DeleteShoppingCart)
		// 修改购物车
		apiV1.PUT("/shopping_carts/:id", v1.UpdateShoppingCart)
		// 购物车详情
		apiV1.GET("/shopping_carts/:id", v1.GetShoppingCart)
		// 查询所有购物车
		apiV1.GET("/shopping_carts", v1.GetAllShoppingCart)
	}

	apiDocsMigrate()

	return r
}

// 将api接口信息自动同步到数据库apis表中
func apiDocsMigrate() {
	apiDocs := util.JsonFileToMap(global.Config.Casbin.ApiJsonFilePath)
	apiModel := model.Api{}

	for k, v := range apiDocs {
		for k2, v2 := range v {
			apiModel.Path = k
			apiModel.Description = v2["description"].(string)
			apiModel.ApiTag = v2["tags"].([]interface{})[0].(string)
			apiModel.Method = strings.ToUpper(k2)

			api_service.AddApi(apiModel)
		}
	}
}
