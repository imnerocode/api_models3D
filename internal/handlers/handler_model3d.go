package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imnerocode/apis/api_models3D/internal/models"
	"github.com/imnerocode/apis/api_models3D/internal/services"
)

type HandlerModel3D struct {
	model3dService *services.ServiceModel3D
}

func NewHandlerModel3D(model3dService *services.ServiceModel3D) *HandlerModel3D {
	return &HandlerModel3D{model3dService: model3dService}
}

func (h *HandlerModel3D) PostModel3D(c *gin.Context) {
	var model3d *models.Model3D

	if err := c.ShouldBindJSON(&model3d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data request"})
		return
	}

	idModel, err := h.model3dService.PostModel3D(model3d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internat server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Model posted", "id_model": idModel})
}

func (h *HandlerModel3D) GetModel3D(c *gin.Context) {

	idModel := c.Query("model_id")
	if idModel == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Model 3D not found"})
		return
	}

	model3d, err := h.model3dService.GetModel3D(idModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": model3d})
}

func (h *HandlerModel3D) DeleteModel3D(c *gin.Context) {
	idModel := c.Query("model_id")
	if idModel == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Model 3D not found"})
		return
	}

	isDeleted, err := h.model3dService.DeleteModel3D(idModel)
	if err != nil || !isDeleted {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot delete model 3D"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Model deleted successfully"})
}

func (h *HandlerModel3D) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid format data"})
		return
	}
	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot open file"})
		return
	}

	defer f.Close()

	data := make([]byte, file.Size)
	_, err = f.Read(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot read file"})
		return
	}

	fileId, err := h.model3dService.UploadFile(file.Filename, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot upload file"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "File uploaded successfully", "file_id": fileId})
}

func (h *HandlerModel3D) DownloadFile(c *gin.Context) {
	// Get file name from parameters
	filename := c.Query("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Filename is required"})
		return
	}

	// Call the service to download the file
	data, err := h.model3dService.DownloadFile(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot download file"})
		return
	}

	// Set headers for the response
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Type", "application/octet-stream")
	c.Data(http.StatusOK, "application/octet-stream", data)
}
