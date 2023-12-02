package auth

type AuthRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
