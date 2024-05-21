package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type MongoRouter struct {
	router *Router
}

func NewMongoRouter(router *Router) {
	m := MongoRouter{
		router: router,
	}
	baseUri := "/mongo"

	m.router.GET(baseUri+"/health", m.health)

	m.router.GET(baseUri+"/bucket", nil)              // 장바구니에 대한 정보
	m.router.GET(baseUri+"/content", nil)             // 상품 정보를 조회
	m.router.GET(baseUri+"/user-bucket-history", nil) // 유저의 구매 이력 정보
}

func (m *MongoRouter) health(c *gin.Context) {
	time.Sleep(3 * time.Second)
	fmt.Println("MongoRouter health")
}
