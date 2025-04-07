package schemas

type UserBodySchema struct {
	RoleID   string `json:"role_id" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,omitempty"`
	Fullname string `json:"fullname" validate:"required"`
	Status   string `json:"status" validate:"required"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type UserResponseSchema struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
