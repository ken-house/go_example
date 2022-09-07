package cache

import (
	"errors"
	"time"

	"github.com/go_example/internal/meta"
	MysqlModel "github.com/go_example/internal/model/mysql"
)

type UserRepository interface {
	SetUserInfo(uid int, userInfo MysqlModel.User) error
	GetUserInfo(uid int) (MysqlModel.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (repo *userRepository) SetUserInfo(uid int, userInfo MysqlModel.User) error {
	cacheKey := GetCacheKey(UserInfoKey, uid)
	meta.CacheDriver.Set(cacheKey, userInfo, time.Hour)
	return nil
}

func (repo *userRepository) GetUserInfo(uid int) (MysqlModel.User, error) {
	cacheKey := GetCacheKey(UserInfoKey, uid)
	data, isExist := meta.CacheDriver.Get(cacheKey)
	if !isExist {
		return MysqlModel.User{}, errors.New("key不存在")
	}
	return data.(MysqlModel.User), nil
}
