package service

import (
	"github.com/gin-gonic/gin"
	MysqlModel "github.com/go_example/internal/model/mysql"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	"github.com/spf13/cast"
	"time"
)

type UserService interface {
	GetUserList(ctx *gin.Context) []MysqlModel.User
	InsertUserList([]MysqlModel.User) error
	Cronjob()
	Job()
}

type userService struct {
	userMysqlRepo MysqlRepo.UserRepository
}

func NewUserService(userMysqlRepo MysqlRepo.UserRepository) UserService {
	return &userService{
		userMysqlRepo: userMysqlRepo,
	}
}

func (svc *userService) GetUserList(c *gin.Context) []MysqlModel.User {
	userList, err := svc.userMysqlRepo.GetUserList()
	if err != nil {
		return []MysqlModel.User{}
	}
	return userList
}

func (svc *userService) InsertUserList(userList []MysqlModel.User) error {
	return svc.userMysqlRepo.InsertRows(userList)
}

func (svc *userService) Cronjob() {
	now := time.Now().Unix()
	user := MysqlModel.User{
		Username: "crontab_" + cast.ToString(now),
		Password: "123456",
		Gender:   0,
	}
	userList := make([]MysqlModel.User, 0, 10)
	userList = append(userList, user)
	svc.userMysqlRepo.InsertRows(userList)
}

func (svc *userService) Job() {
	now := time.Now().Unix()
	user := MysqlModel.User{
		Username: "job_" + cast.ToString(now),
		Password: "123456",
		Gender:   0,
	}
	userList := make([]MysqlModel.User, 0, 10)
	userList = append(userList, user)
	svc.userMysqlRepo.InsertRows(userList)
}
