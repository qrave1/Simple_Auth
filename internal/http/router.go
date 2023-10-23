package http

import (
	"github.com/gin-gonic/gin"
	"rchir7/internal/http/handlers"
)

type Router struct {
	R *gin.Engine
	h *handlers.Handler
}

func NewRouter(h *handlers.Handler) *Router {
	return &Router{
		R: func() *gin.Engine {
			r := gin.Default()
			r.POST("/SignUp", h.SignUp)
			r.GET("/SignIn", h.SignIn)
			return r
		}(),
		h: h,
	}
}
