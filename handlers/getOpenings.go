package handlers

import (
	"net/http"

	g "github.com/gin-gonic/gin"
)

// GetOpeningsHandler handles the GET /api/v1/openings route
func GetOpeningsHandler(ctx *g.Context) {
	ctx.JSON(http.StatusOK, g.H{
		"message": "Get Openings soon!",
	})
}