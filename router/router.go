package router

import (
	g "github.com/gin-gonic/gin"
)

// Init initializes the router
func Init () {
	r := g.Default();
	r.GET("/ping", func(c *g.Context) {
		c.JSON(200, g.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}