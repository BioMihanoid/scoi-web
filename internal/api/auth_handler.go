package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scoi-web/internal/service"
	"strconv"
	"time"
)

type AuthHandler struct {
	services service.Service
}

type signInInput struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=4,max=64"`
}

type tokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

const (
	timeLife = 15
)

func NewAuthHandler(services service.Service) *AuthHandler {
	return &AuthHandler{services: services}
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var input signInInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("error invalid input body"))
		return
	}
	// TODO do check email validate
	id, err := h.services.SignIn(service.SignInInput{Email: input.Email, Password: input.Password})
	if err != nil {
		// TODO do check user exits
		c.JSON(http.StatusInternalServerError, fmt.Errorf("error create user"))
		return
	}
	jwtToken, err := h.services.NewJWT(strconv.Itoa(id), time.Minute*timeLife)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	refreshToken, err := h.services.NewRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  jwtToken,
		RefreshToken: refreshToken,
	})
}
func (h *AuthHandler) SignUp(c *gin.Context) {
	var input signInInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("error invalid input body"))
		return
	}

	id, err := h.services.SignUp(service.SignInInput{Email: input.Email, Password: input.Password})
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("error get user"))
		return
	}
	jwtToken, err := h.services.NewJWT(strconv.Itoa(id), time.Minute*timeLife)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	refreshToken, err := h.services.NewRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  jwtToken,
		RefreshToken: refreshToken,
	})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	id, err := h.GetUserID(c)
	if err != nil || id == 0 {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	jwtToken, err := h.services.NewJWT(strconv.Itoa(id), time.Minute*timeLife)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	refreshToken, err := h.services.NewRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  jwtToken,
		RefreshToken: refreshToken,
	})
}
