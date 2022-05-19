package mysql

import (
	"github.com/go_example/internal/meta"
	"github.com/go_example/internal/model"
	MysqlModel "github.com/go_example/internal/model/mysql"
)

type UserRepository interface {
	GetUserInfoByFormData(model.LoginForm) (MysqlModel.User, error)
	GetUserInfoById(int) (MysqlModel.User, error)
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
	_, err = repo.EngineGroup.Table(repo.Table).Where("username=? and password=?", formData.Username, formData.Password).Get(&user)
	return user, err
}

func (repo *userRepository) GetUserInfoById(id int) (user MysqlModel.User, err error) {
	_, err = repo.EngineGroup.Table(repo.Table).Where("id=?", id).Get(&user)
	return user, err
}
