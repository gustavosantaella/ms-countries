package controllers

import (
	"net/http"
	"project/src/modules/currencies/app/services"

	"github.com/gin-gonic/gin"
)

func CurrencyConversion(gin *gin.Context) {
	base := gin.Query("base")
	qoute := gin.Query("qoute")
	amount := gin.Query("amount")

	response, err := services.CurrencyConversion(base, qoute, amount)
	if err != nil && response == nil {
		gin.JSON(http.StatusBadRequest, err)
		return

	}
	gin.JSON(http.StatusOK, response)
}
