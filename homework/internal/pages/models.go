package pages

type RegisterRequest struct {
	Name     string `form:"name" validate:"required,min=2"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=5"`
}

type LoginRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}
