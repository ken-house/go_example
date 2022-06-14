package mongoClient

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type SingleClient interface {
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
type singleClient struct {
	*mongo.Client
}

type SingleConfig struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	MaxOpen  uint64 `json:"max_open" mapstructure:"max_open"`
}

func NewSingleClient(cfg SingleConfig) (SingleClient, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 得到连接信息
	uri := fmt.Sprintf("mongodb://%s", cfg.Addr)
	clientOption := options.Client().ApplyURI(uri)
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

	return &singleClient{Client: client}, func() {
		_ = client.Disconnect(ctx)
	}, nil
}
