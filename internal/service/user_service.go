package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/model"
	mysqlModel "github.com/go_example/internal/model/mysql"
	mysqlRepo "github.com/go_example/internal/repository/mysql"
	"github.com/spf13/cast"
	"time"
)

type UserService interface {
	GetUserInfoByFormData(ctx *gin.Context, username string, password string) (mysqlModel.User, error)
	GetUserList(ctx *gin.Context) []mysqlModel.User
	InsertUserList([]mysqlModel.User) error
	Cronjob()
	Job()
}

type userService struct {
	userMysqlRepo mysqlRepo.UserRepository
}

func NewUserService(userMysqlRepo mysqlRepo.UserRepository) UserService {
	return &userService{
		userMysqlRepo: userMysqlRepo,
	}
}

func (svc *userService) GetUserInfoByFormData(ctx *gin.Context, username string, password string) (mysqlModel.User, error) {
	formData := model.LoginForm{
		Username: username,
		Password: password,
	}
	return svc.userMysqlRepo.GetUserInfoByFormData(formData)
}

func (svc *userService) GetUserList(ctx *gin.Context) []mysqlModel.User {
	userList, err := svc.userMysqlRepo.GetUserList()
	if err != nil {
		return []mysqlModel.User{}
	}
	return userList
}

func (svc *userService) InsertUserList(userList []mysqlModel.User) error {
	return svc.userMysqlRepo.InsertRows(userList)
}

func (svc *userService) Cronjob() {
	now := time.Now().Unix()
	user := mysqlModel.User{
		Username: "crontab_" + cast.ToString(now),
		Password: "123456",
		Gender:   0,
	}
	userList := make([]mysqlModel.User, 0, 10)
	userList = append(userList, user)
	svc.userMysqlRepo.InsertRows(userList)
}

func (svc *userService) Job() {
	now := time.Now().Unix()
	user := mysqlModel.User{
		Username: "job_" + cast.ToString(now),
		Password: "123456",
		Gender:   0,
	}
	userList := make([]mysqlModel.User, 0, 10)
	userList = append(userList, user)
	svc.userMysqlRepo.InsertRows(userList)
}
