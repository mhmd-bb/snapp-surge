package user

type UsersDto struct {
    Username    string  `json:"username" binding:"required,min=4,max=255"`
    Password    string  `json:"password" binding:"required,min=4,max=255"`
}

type UpdatePasswordDto struct {
    Password    string  `json:"password" binding:"required,min=4,max=255"`
}