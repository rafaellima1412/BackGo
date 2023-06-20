package router

import (
	"backgo/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	basePath := "/api/v1"
	handler.InitializeHandler()

	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		// v1.POST("/opening", handler.CreateOpeningHandler)
		// v1.DELETE("/opening", handler.DeleteOpeningHandler)
		// v1.PUT("/opening", handler.UpdateOpeningHandler)
		// v1.GET("/openings", handler.ListOpeningHandler)
	}

}
