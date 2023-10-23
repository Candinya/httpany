package routers

import (
	"github.com/gin-gonic/gin"
	"httpany/handlers"
)

func R(e *gin.Engine) {
	e.NoRoute(handlers.Any)
}
