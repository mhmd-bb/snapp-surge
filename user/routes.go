package user

import (
"github.com/gin-gonic/gin"
)

type UsersRouter struct {
    usersController      *UsersController
}

func (sr *UsersRouter) SetupRouter(r *gin.Engine) *gin.Engine {

    surge := r.Group("/users")
    {
        surge.POST("register", sr.usersController.CreateUser)
    }

    return r
}

func NewUsersRouter (usersController *UsersController) *UsersRouter{
    return &UsersRouter{usersController: usersController}
}