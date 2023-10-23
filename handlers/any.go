package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Any(ctx *gin.Context) {

	body := make(map[string]any)
	if ctx.Request.Method != http.MethodGet {
		// Auto bind body
		_ = ctx.Bind(&body)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"method":    ctx.Request.Method,
		"url":       ctx.Request.URL.String(),
		"path":      ctx.Request.URL.Path,
		"query":     ctx.Request.URL.Query(),
		"ip":        ctx.ClientIP(),
		"headers":   ctx.Request.Header,
		"body":      body,
		"timestamp": time.Now(),
	})

}
