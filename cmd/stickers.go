package cmd

import (
	"encoding/json"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jenniekibiri/go-stickers/internal/db"
	"github.com/jenniekibiri/go-stickers/internal/models"
)

func CreateSticker(c *gin.Context) {
	// get data from the request
	var body struct {
		Title       string   `json:"title" binding:"required"`
		Description string   ` json:"description" binding:"required"`
		Price       float64  `json:"price"  binding:"required"`
		Quantity    int      `json:"quantity"  binding:"required"`
		Images      []string `json:"images" binding:"required"`
		Size        string   `json:"size"  binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
	}
	imagesJson, err := json.Marshal(body.Images)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to marshal images"})
		return
	}

	sticker := &models.Sticker{
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
		Quantity:    body.Quantity,
		Images:      string(imagesJson),
		Size:        body.Size,
	}

	result := db.DB.Create(sticker)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create sticker"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sticker})
}

func GetStickers(c *gin.Context) {
	var stickers []models.Sticker
	result := []gin.H{}

	db.DB.Find(&stickers)
	for _, sticker := range stickers {
		var images []string
		if err := json.Unmarshal([]byte(sticker.Images), &images); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal images"})
			return
		}

		result = append(result, gin.H{
			"id":          sticker.ID,
			"title":       sticker.Title,
			"description": sticker.Description,
			"price":       sticker.Price,
			"quantity":    sticker.Quantity,
			"images":      images,
			"size":        sticker.Size,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetSticker(c *gin.Context) {
	var sticker models.Sticker
	id := c.Param("id")
	result := db.DB.First(&sticker, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sticker not found"})
		return
	}
	var images []string
	if err := json.Unmarshal([]byte(sticker.Images), &images); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal images"})
		return
	}
	stickers := gin.H{
		"id":          sticker.ID,
		"title":       sticker.Title,
		"description": sticker.Description,
		"price":       sticker.Price,
		"quantity":    sticker.Quantity,
		"images":      images,
		"size":        sticker.Size,
	}

	c.JSON(http.StatusOK, gin.H{"data": stickers})
}

func DeleteSticker(c *gin.Context) {
	var sticker models.Sticker
	id := c.Param("id")

	result := db.DB.Delete(&sticker, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sticker not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sticker deleted successfully"})
}


func Checkout(c *gin.Context) {
	//  save the order to the database 
	var body struct {
		Stickers []int `json:"stickers" binding:"required"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
	}
	var stickers []models.Sticker
	result := []gin.H{}

	db.DB.Find(&stickers, body.Stickers)
	for _, sticker := range stickers {
		var images []string
		if err := json.Unmarshal([]byte(sticker.Images), &images); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal images"})
			return
		}

		result = append(result, gin.H{
			"id":          sticker.ID,
			"title":       sticker.Title,
			"description": sticker.Description,
			"price":       sticker.Price,
			"quantity":    sticker.Quantity,
			"images":      images,
			"size":        sticker.Size,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
	




	




