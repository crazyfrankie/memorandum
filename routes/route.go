package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"memorandum/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"memorandum/controller"
	"memorandum/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte("something-very-secret"))
	router.Use(sessions.Sessions("mysession", store))

	SetRoute(router)

	return router
}

func SetRoute(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1group := r.Group("v1")
	{
		userGroup := v1group.Group("/user")
		{
			userGroup.POST("register", controller.RegisterHandler())
			userGroup.POST("login", controller.LoginHandler())
		}

		authed := v1group.Group("/tasks") // 需要登录保护
		authed.Use(middleware.JWT())
		{
			authed.GET("/list", controller.ListHandler())     // 获取分页任务列表
			authed.GET("", controller.TaskSearch())           // 搜索关键词任务
			authed.GET("/:id", controller.ShowTaskHandler())  // 获取单个任务
			authed.POST("", controller.CreateTaskHandler())   // 创建任务
			authed.DELETE("/:id", controller.DeleteHandler()) // 删除任务
			authed.PUT("/:id", controller.TaskUpdate())       // 更新任务

		}
	}
}
