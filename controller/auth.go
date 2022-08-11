package controller

import (
	"BE_datnd/controller/service"
	"BE_datnd/data"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"BE_datnd/controller/repository"
)

func NewAuthController(
	userRepository repository.UserRepository,
	authService service.AuthService,
) *AuthController {
	return &AuthController{
		userRepository: userRepository,
		authService:    authService,
	}
}

type AuthController struct {
	userRepository repository.UserRepository
	authService    service.AuthService
}

func (p *AuthController) Login(c *gin.Context) {
	var userLogin data.UserLoginInput
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusOK, data.NewResponse(data.CodeErrorBind, data.MessageErrorInvalidEP, nil))
		return
	}

	user, err := p.userRepository.FindUserByEmailAndPassword(userLogin.Email, userLogin.Password)
	if err != nil {
		c.JSON(http.StatusOK, data.NewResponse(data.CodeErrorInvalidEmailPassword, data.MessageErrorInvalidEP, nil))
		return
	}

	token, err := p.authService.GenerateToken(*user)
	if err != nil {
		c.JSON(http.StatusOK, data.NewResponse(data.CodeErrorToken, data.MessageErrorInvalidEP, nil))
		return
	}
	//add role, name to response
	token["role"] = fmt.Sprintf("%d", user.Role)
	token["name"] = user.Name

	c.JSON(http.StatusOK, data.NewResponse(data.CodeSuccess, data.MessageOk, token))
}


func (p *AuthController) ResetPassword(c *gin.Context) {
	var resetPasswordInput data.ResetPasswordInput
	if err := c.ShouldBindJSON(&resetPasswordInput); err != nil {
		c.JSON(http.StatusBadRequest, data.NewResponse(data.CodeErrorInvalidEmailPassword, err.Error(), nil))
		return
	}
	user, err := getCurrentUser(c, p.userRepository)
	if err != nil {
		c.JSON(http.StatusInternalServerError, data.NewResponse(data.CodeInternalError, err.Error(), nil))
		return
	}

	if resetPasswordInput.Password != resetPasswordInput.PasswordConfirm {
		c.JSON(http.StatusBadRequest, data.NewResponse(data.CodeErrorInvalidEmailPassword, "passwords are not matched", nil))
		return
	}

	hashedPassword, err := p.userRepository.HashPassword(resetPasswordInput.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, data.NewResponse(data.CodeInternalError, err.Error(), nil))
		return
	}

	err = p.userRepository.UpdateHashedPassword(hashedPassword, user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, data.NewResponse(data.CodeInternalError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, data.NewResponse(data.CodeSuccess, "password has been changed successfully", nil))
}
