package main

import (
	"log"
	"web/handlers"
	"web/middleware"

	_ "web/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 用户管理系统 API
// @version 1.0
// @description 基于 Gin 框架的用户管理系统
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	// 配置Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 公开路由
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// 需要认证的路由组
	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.POST("/logout", handlers.Logout)
		authGroup.PUT("/user/password", handlers.ChangePassword)
		authGroup.GET("/user/info", handlers.GetUserInfo)
		authGroup.PUT("/user/info", handlers.UpdateUserInfo)
		authGroup.GET("/users", handlers.GetUsers)
		authGroup.GET("/auth/check", handlers.CheckAuth)
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "用户管理系统运行正常",
		})
	})

	log.Println("服务器启动在 :8080 端口")
	log.Println("Swagger文档: http://localhost:8080/swagger/index.html")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
