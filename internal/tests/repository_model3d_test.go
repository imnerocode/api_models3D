package tests

import (
	"fmt"
	"testing"

	"github.com/imnerocode/apis/api_models3D/internal/database"
	"github.com/imnerocode/apis/api_models3D/internal/models"
	"github.com/imnerocode/apis/api_models3D/internal/repositories"
)

func TestRepository3D(t *testing.T) {
	db := database.NewConnectionDB()
	client, err := db.GetConnectionDB()

	if err != nil {
		t.Error(err)
		return
	}
	model3dRepository := repositories.NewModel3DRepository(client)

	model3d := &models.Model3D{Name: "TestName", Format: ".glb", Size: 2040, Tags: []string{"NSFW", "Adult", "Porn"}}
	id, err := model3dRepository.PostModel3D(model3d)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Test repository passed")
	t.Logf("ID model: %s", id)

	t.Run("test get model3D", func(t *testing.T) {
		model3d, err := model3dRepository.GetModel3D(id)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("Test repository passed")
		fmt.Printf("Model3D: %v", model3d)
	})
	t.Run("Test delete model3D", func(t *testing.T) {
		isDeleted, err := model3dRepository.DeleteModel3D(id)
		if err != nil && isDeleted != true {
			t.Error(err)
			return
		}
		t.Log("Test repository passed")

	})
}

func TestRepositoryDelete(t *testing.T) {
	db := database.NewConnectionDB()
	client, err := db.GetConnectionDB()

	if err != nil {
		t.Error(err)
		return
	}
	model3dRepository := repositories.NewModel3DRepository(client)
	id := "6743c3fa53fb1b790eb817ca"
	isDeleted, err := model3dRepository.DeleteModel3D(id)
	if err != nil && isDeleted != true {
		t.Error(err)
		return
	}
	t.Log("Test passed")
}
