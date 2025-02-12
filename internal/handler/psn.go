// internal/handler/psn.go
package handler

import (
	"strconv"

	"github.com/camtrik/ebbilogue-backend/internal/models"
	"github.com/camtrik/ebbilogue-backend/internal/service/psn"
	"github.com/gin-gonic/gin"
)

type PSNHandler struct {
	psnService *psn.PSNService
}

func NewPSNHandler(psnService *psn.PSNService) *PSNHandler {
	return &PSNHandler{
		psnService: psnService,
	}
}

func (h *PSNHandler) parseOptions(c *gin.Context) *models.GetUserTitlesOptions {
	var options *models.GetUserTitlesOptions

	if limitStr := c.Query("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err == nil && limit > 0 {
			if options == nil {
				options = &models.GetUserTitlesOptions{}
			}
			options.Limit = &limit
		}
	}

	if offSetStr := c.Query("offset"); offSetStr != "" {
		offset, err := strconv.Atoi(offSetStr)
		if err == nil && offset > 0 {
			if options == nil {
				options = &models.GetUserTitlesOptions{}
			}
			options.Offset = &offset
		}
	}
	return options
}

func (h *PSNHandler) parseFilter(c *gin.Context) models.TrophyFilter {
	filter := models.TrophyFilter{
		MinProgress: 0,
		Platform:    c.Query("platform"),
		SortBy:      c.Query("sortBy"),
	}

	if progressStr := c.Query("minProgress"); progressStr != "" {
		progress, err := strconv.Atoi(progressStr)
		if err == nil && progress > 0 {
			filter.MinProgress = progress
		}
	}
	return filter
}

// Get User Trophies by accountId
func (h *PSNHandler) GetUserTitles(c *gin.Context) {
	filter := h.parseFilter(c)
	accountId := c.Param("accountId")
	options := h.parseOptions(c)
	resp, err := h.psnService.GetUserTitles(c.Request.Context(), accountId, options, filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}
