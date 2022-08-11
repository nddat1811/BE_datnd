package controller

import (
	"BE_datnd/controller/repository"
	"BE_datnd/data"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RoleAdmin = 0
	RoleIT    = 1
	RoleUser  = 2
)

func NewUserController(
	userRepository repository.UserRepository,
) *UserController {
	return &UserController{
		userRepository: userRepository,
	}
}

type UserController struct {
	userRepository repository.UserRepository
}

func getCurrentUser(ctx *gin.Context, userRepository repository.UserRepository) (*data.User, error) {
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		return nil, fmt.Errorf("get current user from context")
	}
	user, err := userRepository.FindUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *UserController) GetProfileUserByID(c *gin.Context) {
	user, err := getCurrentUser(c, p.userRepository)
	if err != nil {
		c.JSON(http.StatusBadRequest, data.NewResponse(data.CodeErrorBind, data.ErrorGetProfileIT, data.User{}))
		return
	}
	c.JSON(http.StatusOK, data.NewResponse(data.CodeSuccess, data.MessageOk, user))
}

// IT Profile  godoc
// @Summary IT Profile
// @Description IT Profile
// @Accept  json
// @Produce  json
// @Security Authorization
// @Success 200 {object} string
// @Router /api/it/me [get]
func (p *UserController) GetProfileIT(c *gin.Context) {
	user, err := getCurrentUser(c, p.userRepository)
	if err != nil {
		c.JSON(http.StatusBadRequest, data.NewResponse(data.CodeErrorBind, data.ErrorGetProfileIT, data.User{}))
		return
	}
	c.JSON(http.StatusOK, data.NewResponse(data.CodeSuccess, data.MessageOk, user))
}

