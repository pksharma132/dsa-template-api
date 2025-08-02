package handler

import (
	"dsa-template-api/internal/generator"
	"dsa-template-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleTemplate(c *gin.Context) {
	var req models.TemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	template, err := generator.GenerateTemplate(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"language": req.Language,
		"template": template,
	})
}
