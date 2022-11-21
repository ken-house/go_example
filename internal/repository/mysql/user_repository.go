package mysql

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel/attribute"

	"github.com/go_example/internal/meta"
	"github.com/go_example/internal/model"
	MysqlModel "github.com/go_example/internal/model/mysql"
)

type UserRepository interface {
	GetUserInfoByFormData(model.LoginForm) (MysqlModel.User, error)
	GetUserInfoById(ctx context.Context, id int) (MysqlModel.User, error)
	GetUserList() ([]MysqlModel.User, error)
	InsertRows([]MysqlModel.User) error
}

type userRepository struct {
	EngineGroup meta.MysqlGroupClient
	Table       string
}

func NewUserRepository(
	eg meta.MysqlGroupClient,
) UserRepository {
	return &userRepository{
		EngineGroup: eg,
		Table:       "user",
	}
}

func (repo *userRepository) GetUserInfoByFormData(formData model.LoginForm) (user MysqlModel.User, err error) {
	exist, err := repo.EngineGroup.Table(repo.Table).Where("username=?", formData.Username).Get(&user)
	if !exist {
		return MysqlModel.User{}, errors.New("用户名不正确")
	}

	// 密码判断，这里可以自定义加密方法
	if user.Password != formData.Password {
		return MysqlModel.User{}, errors.New("密码不正确")
	}

	return user, nil
}

func (repo *userRepository) GetUserInfoById(ctx context.Context, id int) (user MysqlModel.User, err error) {
	_, span := meta.HttpTracer.Start(ctx, "mysql_userRepository_GetUserInfoById")
	defer span.End()
	exist, err := repo.EngineGroup.Table(repo.Table).Where("id=?", id).Get(&user)
	span.SetAttributes(attribute.Int("id", id), attribute.Bool("exist", exist))
	if !exist {
		return MysqlModel.User{}, errors.New("用户名或密码不正确")
	}
	return user, err
}

func (repo *userRepository) GetUserList() (userList []MysqlModel.User, err error) {
	userList = make([]MysqlModel.User, 0, 100)
	err = repo.EngineGroup.Table(repo.Table).Find(&userList)
	return
}

func (repo *userRepository) InsertRows(userList []MysqlModel.User) (err error) {
	_, err = repo.EngineGroup.Table(repo.Table).Insert(&userList)
	return
}
