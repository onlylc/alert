package main

import (
	"alert/router"
	"alert/sdk"
	"alert/tool"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

var AppRouters = make([]func(), 0)

func init() {

	//注册路由 fixme 其他应用的路由，在本目录新建文件放在init方法
	AppRouters = append(AppRouters, router.InitRouter)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	// initRouter()

	for _, f := range AppRouters {
		f()
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", tool.GetLocaHonst(), 8000),
		Handler: sdk.Runtime.GetEngine(),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		// 服务连接

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: ", err)
		}

	}()

	fmt.Printf("-  Local:   http://localhost:%d/ \r\n", 8000)
	fmt.Printf("-  Network: http://%s:%d/ \r\n", tool.GetLocaHonst(), 8000)
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", tool.GetCurrentTimeStr())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", tool.GetCurrentTimeStr())

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}

// func initRouter() {
// 	// var r *gin.Engine
// 	h := sdk.Runtime.GetEngine()
// 	if h == nil {
// 		h = gin.New()
// 		sdk.Runtime.SetEngine(h)
// 	}
// 	switch h.(type) {
// 	case *gin.Engine:
// 		r = h.(*gin.Engine)
// 	default:
// 		log.Fatal("not support other engine")
// 		os.Exit(-1)
// 	}

// 	//r.Use(middleware.Metrics())

// }
