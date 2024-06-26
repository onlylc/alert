package runtime

import "net/http"

type Application struct {
	engine http.Handler
}

// NewConfig 默认值
func NewConfig() *Application {
	return &Application{}
}

// SetEngine 设置路由引擎
func (e *Application) SetEngine(engine http.Handler) {
	e.engine = engine
}

// GetEngine 获取路由引擎
func (e *Application) GetEngine() http.Handler {
	return e.engine
}
