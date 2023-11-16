package user

type RegisterRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Nama     string `json:"nama" form:"nama" validate:"required,alpha"`
	Password string `json:"password" form:"password" validate:"required"`
}

type RegisterResponse struct {
	Username string `json:"username"`
	Nama     string `json:"nama"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Token    string `json:"token"`
}

type UserListResponse struct {
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}
