package business

import (
	"context"
	"gorm.io/gorm"
	"project-template/model/dto"
	"project-template/model/entity"
)

func (c *Controller) InsertNewBusiness(ctx context.Context, req dto.InsertBusinessDTO) (*dto.InsertBusinessResponseDTO, error) {
	business := entity.Business{
		Name:      req.Name,
		Term:      req.Term,
		Price:     req.Price,
		Location:  req.Location,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Radius:    req.Radius,
		Locale:    req.Locale,
		OpenNow:   req.OpenNow,
	}
	resp, err := c.repository.InsertNewBusiness(ctx, business, req.CategoryIDS)
	if err != nil {
		return nil, err
	}
	return &dto.InsertBusinessResponseDTO{
		ID:   resp.ID,
		Name: resp.Name,
	}, nil
}

func (c *Controller) UpdateBusiness(ctx context.Context, req dto.UpdateBusinessDTO) (*dto.InsertBusinessResponseDTO, error) {
	business := entity.Business{
		Model: gorm.Model{
			ID: req.ID,
		},
		Name:      req.Name,
		Term:      req.Term,
		Price:     req.Price,
		OpenNow:   req.OpenNow,
		Location:  req.Location,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Radius:    req.Radius,
		Locale:    req.Locale,
	}
	resp, err := c.repository.UpdateBusiness(ctx, business, req.CategoryIDS)
	if err != nil {
		return nil, err
	}
	return &dto.InsertBusinessResponseDTO{
		ID:   resp.ID,
		Name: resp.Name,
	}, nil
}

func (c *Controller) DeleteBusiness(ctx context.Context, req dto.DeleteBusinessDTO) error {
	return c.repository.DeleteBusiness(ctx, req.ID)
}

func (c *Controller) SearchBusiness(ctx context.Context, req dto.SearchBusinessDTO) ([]dto.BusinessDTO, error) {
	result := []dto.BusinessDTO{}
	resp, err := c.repository.GetBusinesses(ctx, req)
	if err != nil {
		return result, err
	}
	for _, bus := range resp {
		loc := dto.BusinessDTO{
			ID:        bus.ID,
			Name:      bus.Name,
			Term:      bus.Term,
			Price:     bus.Price,
			Location:  bus.Location,
			Latitude:  bus.Latitude,
			Longitude: bus.Longitude,
			Radius:    bus.Radius,
			Locale:    bus.Locale,
			OpenNow:   bus.OpenNow,
		}
		if bus.OpenAt != nil {
			loc.OpenAt = &bus.OpenAt.Time
		}
		result = append(result, loc)
	}
	return result, nil
}
