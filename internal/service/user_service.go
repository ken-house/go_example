package service

import (
	"github.com/gin-gonic/gin"
	MysqlModel "github.com/go_example/internal/model/mysql"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
)

type UserService interface {
	GetUserList(ctx *gin.Context) []MysqlModel.User
	InsertUserList([]MysqlModel.User) error
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
