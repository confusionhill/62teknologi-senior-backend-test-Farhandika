package business

import (
	"context"
	"project-template/model/dto"
)

type Dependency interface {
	InsertNewBusiness(ctx context.Context, req dto.InsertBusinessDTO) (*dto.InsertBusinessResponseDTO, error)
	UpdateBusiness(ctx context.Context, req dto.UpdateBusinessDTO) (*dto.InsertBusinessResponseDTO, error)
	DeleteBusiness(ctx context.Context, req dto.DeleteBusinessDTO) error
	SearchBusiness(ctx context.Context, req dto.SearchBusinessDTO) ([]dto.BusinessDTO, error)
}
