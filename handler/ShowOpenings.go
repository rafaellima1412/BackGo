package handler

import (
	"backgo/config"
	"backgo/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
	}
	opening := schemas.Opening{}

	db, err := config.InitializeMysql()
	if err != nil {
		logger.Errf("error connecting to the database: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error connecting to the database")
		return
	}
	// Se db for um ponteiro nulo, o erro ocorrerá ao chamar db.Create(&opening)
	if db == nil {
		logger.Err("db object is nil")
		sendError(c, http.StatusInternalServerError, "database object is nil")
		return
	}
	if err := db.First(&opening).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opening witch id: %s not found", id))
		return
	}
	sendSuccess(c, "show-opening", opening)
}
