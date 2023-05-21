package business

import (
	"context"
	"project-template/model/dto"
	"project-template/model/entity"
)

type Dependency interface {
	InsertNewBusiness(ctx context.Context, business entity.Business, categoryIDs []uint64) (*entity.Business, error)
	UpdateBusiness(ctx context.Context, business entity.Business, categoryIDs []uint64) (*entity.Business, error)
	DeleteBusiness(ctx context.Context, uid uint) error
	GetBusinesses(ctx context.Context, req dto.SearchBusinessDTO) ([]entity.Business, error)
}
