package api

import (
	"github.com/gin-gonic/gin"
	"scoi-web/internal/service"
)

type UnAuthHandler struct {
	service service.Service
}

func NewUnAuthHandler(service service.Service) *UnAuthHandler {
	return &UnAuthHandler{}
}

func (u *UnAuthHandler) Methods(c *gin.Context) {}
func (u *UnAuthHandler) Process(c *gin.Context) {}
