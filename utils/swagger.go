package utils

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// DocInit swagger documentation init
func DocInit(app *gin.Engine) {
	// url := ginSwagger.URL("http://10.10.12.217:8080/swagger/doc.json")
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
