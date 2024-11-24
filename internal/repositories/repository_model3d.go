package repositories

import (
	"github.com/imnerocode/apis/api_models3D/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model3DRepository struct {
	Client *mongo.Client
}

func NewModel3DRepository(client *mongo.Client) *Model3DRepository {
	return &Model3DRepository{Client: client}
}

func (r *Model3DRepository) PostModel3D(model3d *models.Model3D) (string, error) {
	var databaseName string
	var collectionName string
}
