package handlers

import (
	"net/http"

	g "github.com/gin-gonic/gin"
)

// DeleteOpeningHandler handles the DELETE /api/v1/opening route
func DeleteOpeningHandler(ctx *g.Context) {
	ctx.JSON(http.StatusOK, g.H{
		"message": "Delete Opening soon!",
	})
}