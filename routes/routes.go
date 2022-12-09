package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"to-do-list/api"
	_ "to-do-list/docs" // 这里需要引入本地已生成文档
	"to-do-list/middleware"
)

// 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default() //生成了一个WSGI应用程序实例
	store := cookie.NewStore([]byte("something-very-secret"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 开启swag
	r.Use(sessions.Sessions("mysession", store))
	///跨域，(是解决不同浏览器对接口请求的限;如果直接服务端请求,是不会触发跨域限制的)
	r.Use(middleware.Cors())
	///分组
	v1 := r.Group("api/v1")
	{
		///ping 请求
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		///再次分一组
		authed := v1.Group("/") //需要登陆保护
		authed.Use(middleware.JWT())
		//authed.Use(middleware.NetLog())
		{
			//任务操作
			authed.GET("tasks", api.ListTasks)
			authed.POST("task", api.CreateTask)
			authed.GET("task/:id", api.ShowTask)
			authed.DELETE("task/:id", api.DeleteTask)
			authed.PUT("task/:id", api.UpdateTask)
			authed.POST("search", api.SearchTasks)
		}
	}
	return r
}
