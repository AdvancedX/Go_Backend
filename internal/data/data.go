package data

import (
	"context"
	"fmt"
	"kratos-realworld/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo, NewProfileRepo)

// Data .
type Data struct {
	client *mongo.Client
	err    error
	db     *mongo.Database
	test   *mongo.Collection
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *mongo.Database) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) *mongo.Database {
	var (
		client *mongo.Client
		err    error
		db     *mongo.Database
	)
	// 连接MongoDB
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(c.Database.Dsn).SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Println("err")
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("err")
	}

	// 选择数据库 my_db
	db = client.Database("kratos")
	return db
}