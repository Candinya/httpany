package inits

import (
	"github.com/gin-gonic/gin"
	"httpany/routers"
)

type Option func(*gin.Engine)

var options []Option

// Include : Register routers
func include(opts ...Option) {
	options = append(options, opts...)
}

func ginInit(middleware ...gin.HandlerFunc) *gin.Engine {

	engine := gin.Default()

	for _, mid := range middleware {
		engine.Use(mid)
	}

	for _, opt := range options {
		opt(engine)
	}

	return engine
}

func WebEngine() *gin.Engine {
	include(routers.R)

	return ginInit()
}
