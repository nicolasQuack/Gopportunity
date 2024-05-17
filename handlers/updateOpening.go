package handlers

import (
	"net/http"

	g "github.com/gin-gonic/gin"
)

// UpdateOpeningHandler handles the PUT /api/v1/opening route
func UpdateOpeningHandler(ctx *g.Context) {
	ctx.JSON(http.StatusOK, g.H{
		"message": "Put Opening soon!",
	})
}