package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang_db_study/config"
	"net/http"
	"time"
)

type Router struct {
	config *config.Config

	engin *gin.Engine
}

func NewRouter(config *config.Config) (*Router, error) {
	r := &Router{
		config: config,
		engin:  gin.New(),
	}
	r.engin.Use(requestTimeOutMiddleWare(5 * time.Second))

	return r, r.engin.Run(config.ServerInfo.Port)
}

func requestTimeOutMiddleWare(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		done := make(chan struct{})

		go func() {
			defer close(done)
		}()

		select {
		case <-done:
			return
		case <-ctx.Done():
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{"error": "Request Time out"})
		}
	}
}
