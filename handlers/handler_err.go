package handlers

import (
	"net/http"

	"github.com/0xatanda/goFeedNews/helper"
	"github.com/gin-gonic/gin"
)

func HandlerError(c *gin.Context) {
	helper.RespondWithError(c.Writer, http.StatusBadRequest, "Something went wrong")
}
