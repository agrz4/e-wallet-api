package route

import (
	"e-wallet-api-go/internal/handler"
	"e-wallet-api-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func (r *Router) User(route *gin.RouterGroup, h *handler.Handler) {
	route.GET("/profiles", middleware.AuthMiddleware(r.jwtService, r.userService), h.Profile)
}
