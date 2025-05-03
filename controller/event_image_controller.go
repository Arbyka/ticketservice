package controller

import (
	"net/http"
	"strconv"
	"ticketing/entity"
	"ticketing/service"

	"github.com/gin-gonic/gin"
)

type EventImageController struct {
	service service.EventImageService
}

func NewEventImageController(service service.EventImageService) *EventImageController {
	return &EventImageController{service}
}

func (c *EventImageController) Upload(ctx *gin.Context) {
	// Ambil event_id dari form
	eventIDStr := ctx.PostForm("event_id")
	if eventIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "event_id is required"})
		return
	}
	
	// Ambil file dari form-data
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	// Simpan file ke folder lokal
	path := "uploads/" + file.Filename
	if err := ctx.SaveUploadedFile(file, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Simpan data ke database
	image := entity.EventImage{
		EventID: parseUint(eventIDStr),
		URL:     "/" + path, // Bisa dibuat full URL jika perlu
	}
	if err := c.service.AddImage(&image); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save event image"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Image uploaded successfully",
		"data":    image,
	})
}

// Helper untuk konversi string ke uint
func parseUint(s string) uint {
	id, _ := strconv.ParseUint(s, 10, 64)
	return uint(id)
}
