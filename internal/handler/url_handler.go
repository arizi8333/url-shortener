package handler

import (
	"net/http"
	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(s *service.URLService) *URLHandler {
	return &URLHandler{service: s}
}

func (h *URLHandler) Shorten(c *gin.Context) {
	var req struct {
		URL         string `json:"url"`
		CustomAlias string `json:"custom_alias"`
		TTLMinutes  int    `json:"ttl_minutes"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := h.service.CreateShortURL(
		req.URL,
		req.CustomAlias,
		req.TTLMinutes,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_url":  "/" + url.ShortCode,
		"expired_at": url.ExpiredAt,
	})
}

func (h *URLHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	userAgent := c.Request.UserAgent()
	ip := c.ClientIP()

	url, err := h.service.GetOriginalURL(code, userAgent, ip)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(302, url.OriginalURL)
}

func (h *URLHandler) Stats(c *gin.Context) {
	code := c.Param("code")

	count, err := h.service.GetStats(code)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	c.JSON(200, gin.H{
		"clicks": count,
	})
}
