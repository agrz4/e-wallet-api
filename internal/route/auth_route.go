package route

import "github.com/gin-gonic/gin"

func (r *Router) Auth(route *gin.RouterGroup, h *handler.Handler) {
	route.POST("/sign-up", h.Register)
	route.POST("/sign-in", h.Login)
	route.POST("/forgot-password", h.ForgotPassword)
	route.POST("/reset-password", h.ResetPassword)
}
