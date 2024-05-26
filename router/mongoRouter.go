package router

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	pM "go.mongodb.org/mongo-driver/mongo"
	"golang_db_study/service/mongo"
	. "golang_db_study/types"
	. "golang_db_study/types/err"
	"time"
)

type MongoRouter struct {
	router   *Router
	mService *mongo.MService
}

func NewMongoRouter(router *Router, mService *mongo.MService) {
	m := MongoRouter{
		router:   router,
		mService: mService,
	}
	baseUri := "/mongo"

	m.router.GET(baseUri+"/health", m.health)

	m.router.GET(baseUri+"/bucket", m.userBucket)                     // 장바구니에 대한 정보
	m.router.GET(baseUri+"/content", m.content)                       // 상품 정보를 조회
	m.router.GET(baseUri+"/user-bucket-history", m.userBucketHistory) // 유저의 구매 이력 정보
}

func (m *MongoRouter) userBucket(c *gin.Context) {
	var req BucketRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		m.router.ResponseErr(c, ErrMsg(BindingFailed, err))
		return
	} else if r, err := m.mService.GetUserBucket(req.User); err != nil {
		if errors.Is(err, pM.ErrNoDocuments) {
			m.router.ResponseErr(c, ErrMsg(NoDocument, err))
		} else {
			m.router.ResponseErr(c, ErrMsg(ServerErr, err))
		}
		return
	} else {
		m.router.ResponseOK(c, r)
	}
}

func (m *MongoRouter) content(c *gin.Context) {
	var req ContentRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		m.router.ResponseErr(c, ErrMsg(BindingFailed, err))
		return
	} else if r, err := m.mService.GetContent(req.Content); err != nil {
		if errors.Is(err, pM.ErrNoDocuments) {
			m.router.ResponseErr(c, ErrMsg(NoDocument, err))
		} else {
			m.router.ResponseErr(c, ErrMsg(ServerErr, err))
		}
		return
	} else {
		m.router.ResponseOK(c, r)
	}

}

func (m *MongoRouter) userBucketHistory(c *gin.Context) {
	var req BucketRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		m.router.ResponseErr(c, ErrMsg(BindingFailed, err))
		return
	} else if r, err := m.mService.GetContent(req.User); err != nil {
		if errors.Is(err, pM.ErrNoDocuments) {
			m.router.ResponseErr(c, ErrMsg(NoDocument, err))
		} else {
			m.router.ResponseErr(c, ErrMsg(ServerErr, err))
		}
		return
	} else {
		m.router.ResponseOK(c, r)
	}
}

func (m *MongoRouter) health(c *gin.Context) {
	time.Sleep(3 * time.Second)
	fmt.Println("MongoRouter health")
}
