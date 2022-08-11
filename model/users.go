package model

import (
	"BE_datnd/controller/repository"
	"BE_datnd/data"
	"BE_datnd/model/convert"
	"BE_datnd/model/tablemodel"
	"BE_datnd/service"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewUserModel(db *gorm.DB) repository.UserRepository {
	return &UserModel{
		db: db,
	}
}

type UserModel struct {
	db *gorm.DB
}

func (m *UserModel) FindUserByEmailAndPassword(email string, password string) (*data.User, error) {
	var user tablemodel.User
	err := m.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = m.CheckPasswordHash(password, user.Password)
	if err != nil {
		return nil, err
	}

	//check if user is deleted
	if user.DeletedAt != nil {
		return nil, errors.New("user is deleted")
	}

	return convert.User(user), nil
}

func (m *UserModel) FindUserByEmail(email string) (*data.User, error) {
	var user tablemodel.User
	err := m.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return convert.User(user), nil
}

func (m *UserModel) FindUserByID(id int) (*data.User, error) {
	var user tablemodel.User
	err := m.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return convert.User(user), nil
}

func (m *UserModel) CheckPasswordHash(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func (m *UserModel) UpdateHashedPassword(hasedPassword string, userID int) error {
	err := m.db.Model(&tablemodel.User{}).Where("id = ?", userID).Update("password", hasedPassword).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *UserModel) UpdateAccountIT(userID int, name string, phone string) error {
	err := m.db.Model(&tablemodel.User{}).Where("id = ?", userID).
		Updates(map[string]interface{}{"name": name, "phone": phone}).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *UserModel) UserUpdateAccount(userdata data.UserUpdateAccountInput) error {
	model := tablemodel.User{
		Id: userdata.Id,
	}

	var user tablemodel.User
	user.Phone = userdata.Phone
	user.Name = userdata.Name

	err := m.db.
		Model(&model).
		Updates(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) CheckRole() int {
	var ctx *gin.Context
	authorizationHeader := ctx.GetHeader("authorization")

	token := strings.Split(authorizationHeader, "Bearer ")[1]
	auth := service.NewAuthService()
	userID, err := auth.GetUserId(ctx, token)
	if err != nil || userID == 0 {
		return 0
	}
	user, err := m.FindUserByID(userID)
	if err != nil {
		return 0
	}
	return user.Role
}
