package router

import "github.com/gin-gonic/gin"

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
