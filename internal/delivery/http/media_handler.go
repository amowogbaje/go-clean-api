package http

import (
	"net/http"
	"go_clean_api/internal/domain"
	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	MediaUsecase domain.MediaUsecase
}

func NewMediaHandler(r *gin.Engine, us domain.MediaUsecase) {
	handler := &MediaHandler{
		MediaUsecase: us,
	}

	api := r.Group("/api")
	{
		api.GET("/", handler.Welcome)
		api.POST("/upload/", handler.UploadMedia)
		api.GET("/medias", handler.GetMedias)
	}
}

// 1. Welcome Route
func (h *MediaHandler) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome to using Go",
	})
}

// 2. Upload Media
func (h *MediaHandler) UploadMedia(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// Save file to the mapped /uploads folder
	destination := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, destination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Process business logic and save to DB
	media, err := h.MediaUsecase.ProcessUpload(file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"data": gin.H{
			"link": media.URL,
		},
	})
}

// 3. Get Medias
func (h *MediaHandler) GetMedias(c *gin.Context) {
	medias, err := h.MediaUsecase.GetAllMedia()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return empty array instead of null if no records
	if medias == nil {
		medias = []domain.Media{}
	}

	c.JSON(http.StatusOK, medias)
}