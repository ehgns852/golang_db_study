package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Router) GET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return r.engin.GET(path, handler...)
}

func (r *Router) POST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return r.engin.GET(path, handler...)
}

func (r *Router) PUT(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return r.engin.GET(path, handler...)
}

func (r *Router) DELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return r.engin.GET(path, handler...)
}

func (r *Router) ResponseOK(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}

func (r *Router) ResponseErr(c *gin.Context, err ...interface{}) {
	c.JSON(http.StatusInternalServerError, err)
}
