package api

import (
	"github.com/gin-gonic/gin"
	"scoi-web/internal/service"
)

type Handler struct {
	jwtKey   string
	services service.Service
}

func NewHandler(services *service.Service, jwtKey string) *Handler {
	return &Handler{
		jwtKey:   jwtKey,
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
		auth.POST("/sign-in", authHandler.SignIn)
		auth.POST("/sign-up", authHandler.SignUp)
		auth.POST("/refresh-token", h.RefreshToken)
	}

	userHandler := NewUserHandler(h.services)
	users := router.Group("/users", h.getAccessWithToken)
	{
		users.POST("/process", userHandler.Process)
		users.GET("/images", userHandler.Images)
		users.GET("/images/:id", userHandler.ImagesById)
		users.DELETE("/images/:id", userHandler.DeleteImagesById)
	}

	return router
}
