package user

import (
    "github.com/gin-gonic/gin"
    "github.com/jackc/pgconn"
    "net/http"
)
import "github.com/go-playground/validator"

type UsersController struct {
    usersService    *UsersService
}

func (uc *UsersController)CreateUser(c *gin.Context) {

    var usersDto UsersDto
    err := c.BindJSON(&usersDto)

    // return exact error on each field
    if err != nil {

        if err.Error() == "EOF" {
            c.JSON(400, gin.H{"message": "provide username and password", "status": 400})
            return
        }
        errs, _ := err.(validator.ValidationErrors)

        e := make(map[string]string)

        for _, err := range errs {
            e[err.Field()] = err.Tag()
        }

        c.JSON(400, e)
        return
    }

    var user User
    err = uc.usersService.CreateUser(&user, &usersDto)

    if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {
        c.JSON(http.StatusConflict, gin.H{"message": "username is taken", "status": http.StatusConflict})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "user created successfully", "status": http.StatusCreated})
}

func NewUsersController(usersService *UsersService) *UsersController {
    return &UsersController{usersService}
}