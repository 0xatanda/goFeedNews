package handlers

import (
	"net/http"

	"github.com/0xatanda/goFeedNews/helper"
	"github.com/gin-gonic/gin"
)

func HandlerReadiness(c *gin.Context) {
	helper.RespondWithJSON(c.Writer, http.StatusOK, struct{}{})
}
