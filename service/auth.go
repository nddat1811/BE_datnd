package service

import (
	"BE_datnd/controller/service"
	data "BE_datnd/data"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = os.Getenv("API_KEY")

const (
	expRefreshToken = 24
	expToken        = 15
	company         = "Hybrid Technologies Viet Nam"
	hostMail        = "humghuy201280@gmail.com"
	subjectMail     = "Email reset password"
	textContent     = "Struction to reset your password:"
)

func NewAuthService() service.AuthService {
	return &AuthService{}
}

type AuthService struct{}

func (m *AuthService) GenerateToken(useInfo data.User) (map[string]string, error) {
	//generate access token
	tokenJwt := jwt.New(jwt.SigningMethodHS256)
	claims := tokenJwt.Claims.(jwt.MapClaims)
	claims["id"] = useInfo.Id
	claims["exp"] = time.Now().Add(time.Hour * expToken).Unix()
	token, err := tokenJwt.SignedString([]byte(jwtKey))
	if err != nil {
		return nil, err
	}
	//generate refresh token
	// refreshTokenJWT := jwt.New(jwt.SigningMethodHS256)
	// rtClaims := refreshTokenJWT.Claims.(jwt.MapClaims)
	// rtClaims["id"] = useInfo.Id
	// rtClaims["exp"] = time.Now().Add(time.Hour * expRefreshToken).Unix()

	// refreshToken, err := refreshTokenJWT.SignedString([]byte(jwtKey))
	// if err != nil {
	// 	return nil, err
	// }

	return map[string]string{
		"access_token": token,
	}, nil
}

func (m *AuthService) ValidateToken(ctx *gin.Context, tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtKey), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["id"].(float64))
		role := uint(claims["role"].(float64))
		ctx.Set("userID", int(userID))
		ctx.Set("role", int(role))
		return nil
	} else {
		return fmt.Errorf("account invalid")
	}
}

func (m *AuthService) GetUserId(ctx *gin.Context, tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["id"].(float64))

		return int(userID), nil
	} else {
		return 0, fmt.Errorf("account invalid")
	}
}
