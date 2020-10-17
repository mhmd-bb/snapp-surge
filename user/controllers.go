package user

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "github.com/jackc/pgconn"
    "net/http"
)

type UsersController struct {
    usersService    IUserService
}

func (uc *UsersController)Login(c *gin.Context) {
    var usersDto UsersDto

    err := c.BindJSON(&usersDto)
    if err != nil {
        _ = c.Error(err)
        return
    }

    var user User
    var jwt string
    err = uc.usersService.LoginUser(&user, &jwt, &usersDto)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "unable to login", "status": http.StatusUnauthorized})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "user logged in successfully", "status": http.StatusOK, "token": jwt})
}

func (uc *UsersController)CreateUser(c *gin.Context) {

    var usersDto UsersDto

    err := c.BindJSON(&usersDto)
    if err != nil {
        _ = c.Error(err)
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

func (uc *UsersController)UpdatePassword(c *gin.Context) {

    var passwordDto UpdatePasswordDto

    err := c.BindJSON(&passwordDto)
    if err != nil {
        _ = c.Error(err)
        return
    }

    // get user info from jwt
    auth, _ := c.Get("auth")
    username := auth.(jwt.MapClaims)["username"].(string)

    // create userDto
    userDto := UsersDto{Username: username, Password: passwordDto.Password}

    var user User
    err = uc.usersService.UpdatePassword(&user, &userDto)

    c.JSON(http.StatusOK, gin.H{"message": "password updated successfully", "status": http.StatusOK})
}

func NewUsersController(usersService IUserService) *UsersController {
    return &UsersController{usersService}
}