package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zzlpeter/dawn-go/middlewares"
	"github.com/zzlpeter/dawn-go/routers/api"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/gin-contrib/pprof"

	_ "github.com/zzlpeter/dawn-go/docs"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	pprof.Register(r)
	// 加载捕获panic中间件
	r.Use(middlewares.PanicCatchMiddleware())
	// 加载swagger模块
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//casView := r.Group("/assassin/cas")
	//{
	//	// 登录
	//	casView.GET("/log_in", api.GetLogIn)
	//	// 退出
	//	casView.GET("/log_out")
	//}

	// 加载其他中间件
	r.Use(middlewares.TraceIDMiddleware())

	taskView := r.Group("/assassin/task")
	{
		// 任务列表
		taskView.GET("/tasks", api.GetTasks)
		// 单项任务增加、修改、查询
		taskView.PUT("/task", api.PutTask)
		taskView.GET("/task", api.GetTask)
		taskView.POST("/task", api.PostTask)
	}

	accountView := r.Group("/assassin/account")
	{
		// roles列表
		accountView.GET("/roles", api.GetRoles)
		// 单项role添加、修改、删除
		accountView.POST("/role", api.PostRole)
		accountView.PUT("/role", api.PutRole)
		accountView.DELETE("/role", api.DeleteRole)
		// permissions列表
		accountView.GET("/permissions", api.GetPermissions)
		// 单项permission添加、修改、删除
		accountView.POST("/permission", api.PostPermission)
		accountView.PUT("/permission", api.PutPermission)
		accountView.DELETE("/permission", api.DeletePermission)
	}

	return r
}