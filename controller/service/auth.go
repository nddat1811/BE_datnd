//go:generate mockgen -source auth.go -destination ../testdata/mock_service/auth_gen.go
package service

import (
	"BE_datnd/data"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	GenerateToken(useInfo data.User) (map[string]string, error)
	ValidateToken(ctx *gin.Context, tokenString string) error
	GetUserId(ctx *gin.Context, tokenString string) (int, error) 
}