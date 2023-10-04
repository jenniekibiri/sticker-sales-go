package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jenniekibiri/go-stickers/cmd"
)

func StickerRoutes(r *gin.Engine) {
	r.GET("/stickers", cmd.GetStickers)
	r.POST("/stickers", cmd.CreateSticker)
	r.GET("/stickers/:id", cmd.GetSticker)
	r.DELETE("/sticker/:id", cmd.DeleteSticker)
	r.PUT("/sticker/:id", cmd.UpdateSticker)
	r.POST("/order", cmd.CreateOrder)
}
