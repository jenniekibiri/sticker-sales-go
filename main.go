package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jenniekibiri/go-stickers/internal/db"
	"github.com/jenniekibiri/go-stickers/routes"
)

func init() {
	db.LoadEnvVariables()
	db.ConnectToDb()
	db.SyncDb()
}

func main() {
	// set up cors

	r := gin.Default()
	r.Use(cors.Default())

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()

	})

	routes.StickerRoutes(r)

	r.Run(":8080")
}
