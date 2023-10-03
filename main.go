package main

import (
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

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	routes.StickerRoutes(r)

	r.Run(":8080")
}
