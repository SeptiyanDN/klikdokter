package users

type RegisterUserInput struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `form:"email" binding:"required,email"`
}
