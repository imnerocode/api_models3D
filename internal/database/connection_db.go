package database

import (
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConnectionDB interface {
	NewConnectionDB(db *mongo.Client) *ConnectionDBImpl
	GetConnectionDB() (*mongo.Client, error)
	LoadEnv(pathToEnv string) string
}

type ConnectionDBImpl struct {
	DB *mongo.Client
}

func NewConnectionDB(db *mongo.Client) *ConnectionDBImpl {
	return &ConnectionDBImpl{DB: db}
}

func (conn *ConnectionDBImpl) GetConnectionDB() *mongo.Client {

}

func (conn *ConnectionDBImpl) LoadEnv(pathToEnv, keyEnv string) string {
	if err := godotenv.Load(pathToEnv); err != nil {
		panic(err)
	}

	varEnv, exists := os.LookupEnv(keyEnv)

	if !exists {
		return " "
	}
	return varEnv
}
