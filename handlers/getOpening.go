package handlers

import (
	"net/http"

	g "github.com/gin-gonic/gin"
)

// GetOpeningHandler handles the GET /api/v1/opening route
func GetOpeningHandler(ctx *g.Context) {
	ctx.JSON(http.StatusOK, g.H{
		"message": "Get Opening soon!",
	})
}