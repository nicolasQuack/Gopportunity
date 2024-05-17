package router

import (
	g "github.com/gin-gonic/gin"
)

// Init initializes the router
func Init () {
	// Create a new router
	r := g.Default();
	
	initRoutes(r)

	// Run the router on port 8080
	r.Run(":8080")
}