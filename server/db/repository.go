package db

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionProducts = "products"
	collectionUsers    = "users"
)

type MongoRepository struct {
	URI      string
	Database string
	PoolSize int

	log *logrus.Logger

	conn *mongo.Client

	Products *ProductsWrapper
	Users    *UserWrapper
}

func (r *MongoRepository) Connect(ctx context.Context) error {
	var err error
	if r.log == nil {
		r.log = logrus.New()
	}

	options := &options.ClientOptions{}
	options.ApplyURI(r.URI)
	options.SetMaxPoolSize(uint64(r.PoolSize))
	options.SetConnectTimeout(5 * time.Second)

	r.log.Infof("connecting to mongo: %s %s", r.URI, r.Database)

	connCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	r.conn, err = mongo.Connect(connCtx, options)
	if err != nil {
		return err
	}

	r.Products = &ProductsWrapper{
		conn: r.conn.Database(r.Database).Collection(collectionProducts),
		log:  r.log,
	}

	r.Users = &UserWrapper{
		conn: r.conn.Database(r.Database).Collection(collectionUsers),
		log:  r.log,
	}

	return nil
}
