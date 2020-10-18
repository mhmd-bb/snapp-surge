package user

import (
    "errors"
    "github.com/mhmd-bb/snapp-surge/auth"
    log "github.com/sirupsen/logrus"
    "gorm.io/gorm"
)

type IUserService interface {
    CreateUser(user *User, dto *UsersDto) (err error)
    UpdatePassword(user *User, dto *UsersDto) (err error)
    LoginUser(user *User,jwt *string , dto *UsersDto) (err error)
    GetByUsername(user *User, username string) (err error)
    CreateDefaultUser(username string, password string) (err error)
}

type UsersService struct {
    DB  *gorm.DB
    authService auth.IJwtAuthService
    logger *log.Logger
}

// Get One user by username
func (us *UsersService) GetByUsername(user *User, username string) (err error) {
    err = us.DB.Where("username = ?", username).First(user).Error

    if err != nil {
        us.logger.WithFields(log.Fields{"userID": user.ID}).Info("user not found")
    }

    return err
}


// Create User and Hash plane password before save
func (us *UsersService) CreateUser(user *User, dto *UsersDto) (err error){

    *user = User{Username: dto.Username, Password: dto.Password}

    // hash user password
    err = user.setPassword(user.Password)
    if err != nil {
        return err
    }

    // save user in db
    err = us.DB.Create(user).Error
    if err != nil {
        us.logger.WithFields(log.Fields{"error": err, "userID": user.ID}).Error("user creation failed")
    }
    return err

}

func (us *UsersService) LoginUser(user *User,jwt *string , dto *UsersDto) (err error) {

    *user = User{Username: dto.Username, Password: dto.Password}

    var dbUser User
    err = us.GetByUsername(&dbUser, user.Username)

    if err != nil {
        return err
    }

    // check user password
    err = dbUser.checkPassword(user.Password)
    if err != nil {
        us.logger.WithFields(log.Fields{"userID": user.ID}).Info("user password entered wrong")
        return err
    }

    *user = dbUser

    *jwt = us.authService.GenerateJwtToken(user.Username)

    us.logger.WithFields(log.Fields{"userID": user.ID}).Info("user logged in successfully")

    return nil

}

func (us *UsersService) UpdatePassword(user *User, dto *UsersDto) (err error) {
    err = us.GetByUsername(user, dto.Username)
    if err != nil {
        return err
    }

    // hash password
    err = user.setPassword(dto.Password)
    if err != nil {
        return err
    }

    // update password
    err = us.DB.Model(user).Update("password", user.Password).Error

    us.logger.WithFields(log.Fields{"userID": user.ID}).Info("user password updated successfully")

    return err
}

func (us *UsersService) CreateDefaultUser(username string, password string) (err error) {

    var user User
    userDto := UsersDto{username, password}

    err = us.DB.First(&user).Error

    // if no user is found in database, create one
    if errors.Is(err, gorm.ErrRecordNotFound) {
        us.logger.WithFields(log.Fields{"defaultUsername": userDto.Username}).Info("no users found in database; creating a default user")

        err = us.CreateUser(&user, &userDto)

        if err != nil {
            us.logger.WithFields(log.Fields{"defaultUsername": userDto.Username}).Fatal("default user creation failed")
        }
    }
    return err
}


func NewUsersService(db *gorm.DB, authService auth.IJwtAuthService, logger *log.Logger) *UsersService{
    return &UsersService{DB: db, authService: authService, logger: logger}
}