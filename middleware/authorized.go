package middleware

import (
	"BE_datnd/controller/repository"
	"BE_datnd/data"
	"BE_datnd/service"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

type Authorized struct{}

func (m *Authorized) TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := service.NewAuthService()

		authorizationHeader := ctx.GetHeader(data.AuthorHeader)

		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				data.NewResponse(data.CodeErrorInvalidToken, data.MessageErrorTokenInvalid, nil))
			return
		}

		fields := strings.Fields(authorizationHeader)

		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				data.NewResponse(data.CodeErrorInvalidToken, data.MessageErrorTokenInvalid, nil))
			return
		}

		if strings.ToLower(fields[0]) != data.Bearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				data.NewResponse(data.CodeErrorInvalidToken, data.MessageErrorTokenInvalid, nil))
			return
		}

		if fields[1] == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				data.NewResponse(data.CodeErrorInvalidToken, data.MessageErrorTokenInvalid, nil))
			return
		}

		// err := auth.ValidateToken(ctx, fields[1])
		// if err != nil {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized,
		// 		data.NewResponse(data.CodeErrorInvalidToken, data.MessageErrorTokenInvalid, nil))
		// 	return
		// }

		token := strings.Split(authorizationHeader, "Bearer ")[1]
		err := auth.ValidateToken(ctx, token)
		if err != nil {
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}

		ctx.Next()
	}
}

func (m *Authorized) TokenResetPasswordMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		token := strings.Split(authorizationHeader, "Bearer ")[1]
		auth := service.NewAuthService()
		err := auth.ValidateToken(ctx, token)
		if err != nil {
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}
		ctx.Next()
	}
}

func (m *Authorized) CheckUser(u repository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		token := strings.Split(authorizationHeader, "Bearer ")[1]
		auth := service.NewAuthService()
		userID, err := auth.GetUserId(ctx, token)
		if err != nil || userID == 0 {
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}
		user, err := u.FindUserByID(userID)
		if err != nil || user.Role != 2 {
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}
		ctx.Next()
	}
}

func (m *Authorized) CheckIT(u repository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		token := strings.Split(authorizationHeader, "Bearer ")[1]
		auth := service.NewAuthService()
		userID, err := auth.GetUserId(ctx, token)
		if err != nil || userID == 0 {
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}
		user, err := u.FindUserByID(userID)
		if err != nil || user.Role != 1 {
			fmt.Println("sos", user.Role)
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}
		ctx.Next()
	}
}
