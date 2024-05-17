package router

import (
	"context"
	"fmt"
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

	NewMongoRouter(r)

	return r, r.engin.Run(config.ServerInfo.Port)
}

func requestTimeOutMiddleWare(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 타임아웃 시간이 지나면 ctx가 자동으로 취소됨
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel() // 함수가 끝날 때 cancel 호출하여 리소스 정리

		// 새로운 타임아웃 컨텍스트를 요청에 설정
		c.Request = c.Request.WithContext(ctx)
		done := make(chan struct{})

		// 요청을 처리할 고루틴 생성
		go func() {
			defer close(done) // 고루틴이 끝나면 done 채널 닫기
			c.Next()
		}()

		// select 문을 통해 작업 완료 또는 타임아웃 대기
		select {
		case <-done:
			// 요청이 완료되면 여기로 옴
			fmt.Println("요청 완료")
		case <-ctx.Done():
			// 타임아웃이 발생하면 여기로 옴
			if ctx.Err() == context.DeadlineExceeded {
				c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{"error": "Request Time out"})
				return
			}
		}
	}
}
