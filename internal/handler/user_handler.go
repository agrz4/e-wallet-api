package handler

import (
	"e-wallet-api-go/internal/dto"
	"e-wallet-api-go/internal/model"
	"e-wallet-api-go/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Profile(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	input := &dto.WalletRequestBody{}
	input.UserID = int(user.ID)
	wallet, err := h.walletService.GetWalletByUserId(input)
	if err != nil {
		response := utils.ErrorResponse("show profile failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedUser := dto.FormatUserDetail(user, wallet)
	response := utils.SuccessResponse("show profile success", http.StatusOK, formattedUser)
	c.JSON(http.StatusOK, response)
}
