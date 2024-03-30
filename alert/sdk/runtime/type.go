package runtime

import "net/http"

type Runtime interface {

	// SetEngine 使用的路由
	SetEngine(engine http.Handler)
	GetEngine() http.Handler
}
