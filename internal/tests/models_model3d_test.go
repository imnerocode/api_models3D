package tests

import (
	"testing"
	"time"

	"github.com/imnerocode/apis/api_models3D/internal/models"
)

func TestModel3D(t *testing.T) {
	expectedName := "test.glb"
	expectedFormat := ".glb"
	expectedSize := int64(2048)
	expectedCreatedAt := time.Now()

	modelTest := &models.Model3D{Name: expectedName, Format: expectedFormat, Size: expectedSize, CreatedAt: expectedCreatedAt}

	if modelTest.Name != expectedName {
		t.Errorf("expected Name to be %v, fot %v", expectedName, modelTest.Name)
	}

	if modelTest.Format != expectedFormat {
		t.Errorf("expected Format to be %v, fot %v", expectedFormat, modelTest.Format)
	}

	if modelTest.Size != expectedSize {
		t.Errorf("expected Size to be %v, fot %v", expectedSize, modelTest.Size)
	}

	if !modelTest.CreatedAt.Equal(expectedCreatedAt) {
		t.Errorf("expected CreatedAt to be %v, got %v", expectedCreatedAt, modelTest.CreatedAt)
	}

	t.Log("Model3D structure validation passed!")
	t.Logf("Model created: %v", *modelTest)

}
