package router

import (
	"backgo/handler"

	"backgo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {

	basePath := "/api/v1"
	handler.InitializeHandler()
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server BackGo server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "BackGo.swagger.io"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		// v1.POST("/opening", handler.CreateOpeningHandler)
		// v1.DELETE("/opening", handler.DeleteOpeningHandler)
		// v1.PUT("/opening", handler.UpdateOpeningHandler)
		// v1.GET("/openings", handler.ListOpeningHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
