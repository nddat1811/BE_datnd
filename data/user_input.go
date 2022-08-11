package data

type UserLoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type ForgotPasswordInput struct {
	Email string `form:"email" binding:"required,email"`
}

type ResetPasswordInput struct {
	Password        string `json:"password" binding:"required,min=6"`
	PasswordConfirm string `json:"password_confirm" binding:"required,min=6"`
}
type UserUpdateAccountInput struct {
	Id    int    `json:"-"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UserChangePasswordInput struct {
	Id              int    `json:"id"`
	CurrentPassword string `json:"current_password"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
