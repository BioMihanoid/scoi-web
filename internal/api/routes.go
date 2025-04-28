package api

import (
	"github.com/gin-gonic/gin"
	"scoi-web/internal/service"
)

type Handler struct {
	services service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: *services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	unAuthHandler := NewUnAuthHandler(h.services)
	router.GET("/methods", unAuthHandler.Methods)
	router.POST("/process", unAuthHandler.Process)

	authHandler := NewAuthHandler(h.services)
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", authHandler.SignIN)
		auth.POST("/sign-up", authHandler.SignUP)
		auth.POST("/refresh-token", authHandler.RefreshToken)
	}

	userHandler := NewUserHandler(h.services)
	users := router.Group("/users")
	{
		users.POST("/process", userHandler.Process)
		users.GET("/images", userHandler.Images)
		users.GET("/images/:id", userHandler.ImagesById)
		users.DELETE("/images/:id", userHandler.DeleteImagesById)
	}

	return router
}
