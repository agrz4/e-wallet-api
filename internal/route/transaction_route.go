package route

import (
	"e-wallet-api-go/internal/handler"
	"e-wallet-api-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func (r *Router) Transaction(route *gin.RouterGroup, h *handler.Handler) {
	route.Use(middleware.AuthMiddleware(r.jwtService, r.userService))
	route.GET("/transactions", h.GetTransactions)
	route.POST("/top-up", h.TopUp)
	route.POST("/transfer", h.Transfer)
}
