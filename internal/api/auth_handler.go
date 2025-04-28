package api

import (
	"github.com/gin-gonic/gin"
	"scoi-web/internal/service"
)

type AuthHandler struct {
	service service.Service
}

func NewAuthHandler(service service.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) SignIN(c *gin.Context)       {}
func (h *AuthHandler) SignUP(c *gin.Context)       {}
func (h *AuthHandler) RefreshToken(c *gin.Context) {}
