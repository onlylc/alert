package router

import (
	"alert/apis"

	"github.com/gin-gonic/gin"
)

// 路由示例
func InitSysRouter(r *gin.Engine) *gin.Engine {
	g := r.Group("")

	// 无需认证的路由
	sysBaseRouter(g)

	return r
}

// 无需认证的路由示例
func sysBaseRouter(r *gin.RouterGroup) {

	v1 := r.Group("/api/v1")
	{
		v1.POST("/", apis.PostApi)
		// Refresh time can be longer than token timeout
		v1.GET("/", apis.GetApi)
	}

}
