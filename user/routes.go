package user

import (
"github.com/gin-gonic/gin"
    "github.com/mhmd-bb/snapp-surge/auth"
)

type UsersRouter struct {
    usersController      *UsersController
}

func (sr *UsersRouter) SetupRouter(r *gin.Engine) *gin.Engine {

    surge := r.Group("/users", BadRequestErrorMiddleware())
    {
        surge.POST("register", auth.AuthorizeJWT(), sr.usersController.CreateUser)
        surge.POST("login", sr.usersController.Login)
        surge.POST("update/password", auth.AuthorizeJWT(), sr.usersController.UpdatePassword)

    }

    return r
}

func NewUsersRouter (usersController *UsersController) *UsersRouter{
    return &UsersRouter{usersController: usersController}
}