package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scoi-web/internal/service"
)

type UserHandler struct {
	services service.Service
}

func NewUserHandler(services service.Service) *UserHandler {
	return &UserHandler{
		services: services,
	}
}

func (u *UserHandler) Process(c *gin.Context) {}
func (u *UserHandler) Images(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
func (u *UserHandler) ImagesById(c *gin.Context)       {}
func (u *UserHandler) DeleteImagesById(c *gin.Context) {}
