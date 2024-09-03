package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/0xatanda/goFeedNews/config"
	"github.com/0xatanda/goFeedNews/helper"
	db "github.com/0xatanda/goFeedNews/sql/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandlerCreateUser(cfg *config.APIConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		type parameters struct {
			Name string `json:"name"`
		}
		decoder := json.NewDecoder(c.Request.Body)
		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil {
			helper.RespondWithError(c.Writer, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s", err))
			return
		}
		user, err := cfg.DB.CreateUser(c.Request.Context(), db.CreateUserParams{
			ID:        uuid.New(),
			Name:      params.Name,
			CreateAt:  time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})
		if err != nil {
			helper.RespondWithError(c.Writer, http.StatusBadRequest, fmt.Sprintf("Couldn't create user: %s", err))
			return
		}
		helper.RespondWithJSON(c.Writer, http.StatusOK, user)
	}
}
