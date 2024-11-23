package tests

import (
	"testing"

	"github.com/imnerocode/apis/api_models3D/internal/database"
)

func TestConnectionDB(t *testing.T) {
	connectionDB := database.NewConnectionDB()

	_, err := connectionDB.GetConnectionDB()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Test passed!")
}
