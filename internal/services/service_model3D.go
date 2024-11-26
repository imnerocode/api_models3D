package services

import (
	"github.com/imnerocode/apis/api_models3D/internal/models"
	"github.com/imnerocode/apis/api_models3D/internal/repositories"
)

type ServiceModel3D struct {
	model3dRepository *repositories.Model3DRepository
}

func NewModel3DService(model3dRepository *repositories.Model3DRepository) *ServiceModel3D {
	return &ServiceModel3D{model3dRepository: model3dRepository}
}

func (s *ServiceModel3D) PostModel3D(model3d *models.Model3D) (string, error) {
	idModel, err := s.model3dRepository.PostModel3D(model3d)
	if err != nil {
		return "", err
	}

	return idModel, nil
}

func (s *ServiceModel3D) GetModel3D(idModel string) (*models.Model3D, error) {
	model3d, err := s.model3dRepository.GetModel3D(idModel)
	if err != nil {
		return nil, err
	}

	return model3d, nil
}
func (s *ServiceModel3D) DeleteModel3D(idModel string) (bool, error) {
	ok, err := s.model3dRepository.DeleteModel3D(idModel)
	if err != nil && !ok {
		return false, err
	}

	return true, nil
}
