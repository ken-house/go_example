package mongoClient

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	StartSession(opts ...*options.SessionOptions) (mongo.Session, error)
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
	ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error)
	ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error)
	UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
	UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo.SessionContext) error) error
	Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
	NumberSessionsInProgress() int
}
type mongoClient struct {
	*mongo.Client
}

type MongoConfig struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	MaxOpen  uint64 `json:"max_open" mapstructure:"max_open"`
}

func NewMongoClient(cfg MongoConfig) (MongoClient, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 得到连接信息
	clientOption := options.Client().ApplyURI(cfg.Addr)
	if cfg.MaxOpen > 0 {
		clientOption.SetMaxPoolSize(cfg.MaxOpen)
	}

	// 连接
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return &mongoClient{Client: client}, func() {
		_ = client.Disconnect(ctx)
	}, nil
}
