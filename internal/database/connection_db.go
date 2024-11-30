package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/imnerocode/apis/api_models3D/internal/config"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectionDB interface {
	GetConnectionDB() (*mongo.Client, error)
	LoadEnv(pathToEnv string) string
}

type ConnectionDBImpl struct {
	Client *mongo.Client
}

func NewConnectionDB() *ConnectionDBImpl {
	return &ConnectionDBImpl{}
}

func (conn *ConnectionDBImpl) GetConnectionDB() (*mongo.Client, error) {
	var pathToEnv string
	var keyEnv string

	pathToEnv = config.PathToEnv
	keyEnv = config.EnvMongoURI

	uri := conn.LoadEnv(pathToEnv, keyEnv)

	if uri == "" {
		return nil, fmt.Errorf("failed to load environment variable %s", keyEnv)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	conn.Client = client
	return client, nil

}

func (conn *ConnectionDBImpl) GetBucketFs(db *mongo.Database) (*gridfs.Bucket, error) {
	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func (conn *ConnectionDBImpl) LoadEnv(pathToEnv, keyEnv string) string {
	if err := godotenv.Load(pathToEnv); err != nil {
		panic(err)
	}

	varEnv, exists := os.LookupEnv(keyEnv)

	if !exists {
		return ""
	}
	return varEnv
}
