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

// CreateOpeningHandler handles the POST /api/v1/opening route
func CreateOpeningHandler(ctx *g.Context) {
	ctx.JSON(http.StatusOK, g.H{
		"message": "Post Opening soon!",
	})
}

// DeleteOpeningHandler handles the DELETE /api/v1/opening route
func DeleteOpeningHandler(ctx *g.Context) {
	ctx.JSON(http.StatusOK, g.H{
		"message": "Delete Opening soon!",
	})
}

// UpdateOpeningHandler handles the PUT /api/v1/opening route
func UpdateOpeningHandler(ctx *g.Context) {
	ctx.JSON(http.StatusOK, g.H{
		"message": "Put Opening soon!",
	})
}

// GetOpeningsHandler handles the GET /api/v1/openings route
func GetOpeningsHandler(ctx *g.Context) {
	ctx.JSON(http.StatusOK, g.H{
		"message": "Get Openings soon!",
	})
}
