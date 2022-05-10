package mysql

import (
	"github.com/go_example/internal/meta"
	MysqlModel "github.com/go_example/internal/model/mysql"
)

type UserRepository interface {
	GetUserInfo(int) (MysqlModel.User, error)
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

func (repo *userRepository) GetUserInfo(id int) (user MysqlModel.User, err error) {
	_, err = repo.EngineGroup.Table(repo.Table).Where("id=?", id).Get(&user)
	return user, err
}
