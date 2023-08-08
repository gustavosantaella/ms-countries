package router

import (
	"project/src/modules/countries/infra/controllers"
	currencies "project/src/modules/currencies/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Apirouter(engine *gin.Engine) {

	engine.GET("/countries", controllers.GetAll)
	engine.GET("/countries/:iso3", controllers.GetOneByIso3)
	engine.GET("/currencies/conversion", currencies.CurrencyConversion)
	engine.GET("/currencies", currencies.AllCurrencies)
}
