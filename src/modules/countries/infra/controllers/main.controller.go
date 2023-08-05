package controllers

import (
	"net/http"
	"project/src/modules/countries/app/services"

	"github.com/gin-gonic/gin"
)

func GetAll(gin *gin.Context) {

	countries := services.GetAll()

	gin.JSON(http.StatusOK, countries)
}

func GetOneByIso3(gin *gin.Context) {

	iso3, err := gin.Params.Get("iso3")
	if err != true {

	}
	country := services.GetOneByIso3(iso3)
	gin.JSON(http.StatusOK, country)

}
