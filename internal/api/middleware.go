package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) getAccessWithToken(ctx *gin.Context) {
	authHeaderValue := ctx.GetHeader("Authorization")
	if authHeaderValue == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	parts := strings.Split(authHeaderValue, " ")
	if len(parts) > 1 && parts[0] == "Bearer" {
		userId, err := h.getUserIdFromToken(parts[1])
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("userId", userId)
	}
	ctx.Next()
}

func (h *Handler) getUserIdFromToken(tokenString string) (string, error) {

	draft, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(h.jwtKey), nil
		},
	)

	if err != nil {
		return "", err
	}

	if draft.Valid {
		id := draft.Claims.(*jwt.StandardClaims).Subject
		return id, nil
	}

	return "", err
}

func (h *Handler) GetUserID(ctx *gin.Context) (int, error) {
	authHeaderValue := ctx.GetHeader("Authorization")
	if authHeaderValue == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return 0, nil
	}
	parts := strings.Split(authHeaderValue, " ")
	var id int
	if len(parts) > 1 && parts[0] == "Bearer" {
		userId, err := h.getUserIdFromToken(parts[1])
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return 0, err
		}
		id, err = strconv.Atoi(userId)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, fmt.Errorf("error convert string to int"))
			return 0, err
		}
	}
	return id, nil
}
