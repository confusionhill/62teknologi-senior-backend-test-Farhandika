package business

import (
	"context"
	"errors"
	"project-template/model/dto"
	"project-template/model/entity"
)

func (r *Repository) InsertNewBusiness(ctx context.Context, business entity.Business) (*entity.Business, error) {
	if err := r.Database.Create(&business).Error; err != nil {
		return nil, err
	}
	return &business, nil
}

func (r *Repository) UpdateBusiness(ctx context.Context, business entity.Business) (*entity.Business, error) {
	err := r.Database.Save(&business).Error
	return &business, err
}

func (r *Repository) DeleteBusiness(ctx context.Context, uid uint) error {
	resp := r.Database.Exec("DELETE FROM businesses WHERE id = ?", uid)
	if resp.RowsAffected < 1 {
		return errors.New("cannot delete, id not found")
	}
	return resp.Error
}

func (r *Repository) GetBusinesses(ctx context.Context, req dto.SearchBusinessDTO) ([]entity.Business, error) {
	resp := []entity.Business{}
	offset := (req.Page - 1) * req.Limit
	query := "select * from businesses where term ~ ? and locale ~ ? and location ~ ? limit ? offset ?"
	if req.Location != "" {
		query = "select * from businesses where term ~ ? and locale ~ ? and longitude = ? and latitude = ? limit ? offset ?"
		err := r.Database.Raw(query, req.Term, req.Locale, req.Longitude, req.Latitude, req.Limit, offset).Scan(&resp).Error
		return resp, err
	}
	err := r.Database.Raw(query, req.Term, req.Locale, req.Location, req.Limit, offset).Scan(&resp).Error
	return resp, err
}
