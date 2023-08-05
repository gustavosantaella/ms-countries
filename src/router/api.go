package router

import (
	"project/src/modules/countries/infra/controllers"

	"github.com/gin-gonic/gin"
)

func Apirouter(engine *gin.Engine) {

	engine.GET("/countries", controllers.GetAll)
	engine.GET("/countries/:iso3", controllers.GetOneByIso3)
}
