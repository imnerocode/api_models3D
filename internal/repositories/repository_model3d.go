package repositories

import (
	"bytes"
	"context"
	"time"

	"github.com/imnerocode/apis/api_models3D/internal/config"
	"github.com/imnerocode/apis/api_models3D/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type Model3DRepository struct {
	Client *mongo.Client
}

func NewModel3DRepository(client *mongo.Client) *Model3DRepository {
	return &Model3DRepository{Client: client}
}

func (r *Model3DRepository) UploadFile(filename string, data []byte, bucket *gridfs.Bucket) (fileId interface{}, err error) {
	idFile, err := bucket.UploadFromStream(filename, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return idFile, nil
}

func (r *Model3DRepository) DownloadFile(filename string, dest *bytes.Buffer, bucket *gridfs.Bucket) ([]byte, error) {
	_, err := bucket.DownloadToStreamByName(filename, dest)
	if err != nil {
		return nil, err
	}

	return dest.Bytes(), nil

}

func (r *Model3DRepository) PostModel3D(model3d *models.Model3D) (string, error) {
	db := r.Client.Database(config.DatabaseName)
	collection := db.Collection(config.CollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, model3d)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (r *Model3DRepository) GetModel3D(id string) (*models.Model3D, error) {
	var model3D *models.Model3D
	db := r.Client.Database(config.DatabaseName)
	collection := db.Collection(config.CollectionName)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&model3D); err != nil {
		return nil, err
	}
	return model3D, nil
}

func (r *Model3DRepository) DeleteModel3D(id string) (bool, error) {
	var deletedModel3D *models.Model3D
	db := r.Client.Database(config.DatabaseName)
	collection := db.Collection(config.CollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	resultErr := collection.FindOneAndDelete(ctx, bson.M{"_id": objectID}).Decode(&deletedModel3D)

	if resultErr != nil {
		return false, nil
	}

	return true, nil

}
