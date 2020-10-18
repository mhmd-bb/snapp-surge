package user

type UsersDto struct {
	// username must be longer than 4 chars
	Username string `json:"username" binding:"required,min=4,max=255" example:"admin"`
	// password must be longer than 4 chars
	Password string `json:"password" binding:"required,min=4,max=255" example:"admin"`
}

type UpdatePasswordDto struct {
	// password must be longer than 4 chars
	Password string `json:"password" binding:"required,min=4,max=255"`
}
