package router

import (
	"alert/sdk"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	var r *gin.Engine

	h := sdk.Runtime.GetEngine()
	if h == nil {
		h = gin.New()
		sdk.Runtime.SetEngine(h)
	}

	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	InitSysRouter(r)
}
