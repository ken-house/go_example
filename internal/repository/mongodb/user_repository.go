package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/go_example/internal/meta"
	MongoModel "github.com/go_example/internal/model/mongodb"
)

type UserRepository interface {
	SetUserInfo(int, string, string) error
	GetUserInfo(int) (MongoModel.User, error)
}

type userRepository struct {
	client     meta.MongoSingleClient
	database   string
	collection string
}

func NewUserRepository(client meta.MongoSingleClient) UserRepository {
	return &userRepository{
		client:     client,
		database:   "test",
		collection: "user",
	}
}

func (repo *userRepository) SetUserInfo(id int, username, password string) error {
	userDocument := MongoModel.User{
		Id:       id,
		Username: username,
		Password: password,
	}
	_, err := repo.client.Database(repo.database).Collection(repo.collection).InsertOne(context.TODO(), userDocument)
	if err != nil {
		return nil
	}
	return nil
}

func (repo *userRepository) GetUserInfo(uid int) (MongoModel.User, error) {
	var userInfo MongoModel.User

	err := repo.client.Database(repo.database).Collection(repo.collection).FindOne(context.TODO(), bson.D{}).Decode(&userInfo)
	if err != nil {
		return userInfo, err
	}
	return userInfo, nil
}
