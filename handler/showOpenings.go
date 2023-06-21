package handler

import (
	"backgo/config"
	"backgo/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1

//	@Summary		Show Opening
//	@Description	Show Opening job opening
//	@Tags			Openings
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"Opening indetification"
//	@Success		200	{object}	ShowOpeningReponse
//	@Failure		500	{object}	ErrorReponse
//	@Router			/opening [get]

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
	// Se db for um ponteiro nulo, o erro ocorrerá ao chamar db
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
