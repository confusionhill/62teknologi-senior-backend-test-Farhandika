package business

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"project-template/model/dto"
	"project-template/model/entity"
)

func (r *Repository) InsertNewBusiness(ctx context.Context, business entity.Business, categoryIDs []uint64) (*entity.Business, error) {
	if err := r.Database.Transaction(func(tx *gorm.DB) error {
		err := r.Database.Create(&business).Error
		if err != nil {
			return err
		}
		businessCats := []entity.BusinessCategory{}
		for _, catId := range categoryIDs {
			businessCats = append(businessCats, entity.BusinessCategory{BusinessID: business.ID, CategoryID: catId})
		}
		if len(businessCats) > 0 {
			err = r.Database.Create(&businessCats).Error
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &business, nil
}

func (r *Repository) UpdateBusiness(ctx context.Context, business entity.Business, categoryIDs []uint64) (*entity.Business, error) {

	err := r.Database.Transaction(func(tx *gorm.DB) error {
		err := r.Database.Exec("DELETE FROM business_categories WHERE business_id = ?", business.ID).Error
		if err != nil {
			return err
		}
		businessCats := []entity.BusinessCategory{}
		for _, catId := range categoryIDs {
			businessCats = append(businessCats, entity.BusinessCategory{BusinessID: business.ID, CategoryID: catId})
		}
		if len(businessCats) > 0 {
			err = r.Database.Create(&businessCats).Error
			if err != nil {
				return err
			}
		}
		return r.Database.Save(&business).Error
	})

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
	query := "select * from businesses where term ~ ? and locale ~ ? and longitude = ? and latitude = ? and open_now = ? limit ? offset ?"

	if req.Price != 0 {
		query = "select * from businesses where term ~ ? and locale ~ ? and price = ? and longitude = ? and latitude = ? and open_now = ? limit ? offset ?"
		if req.Longitude == 0 && req.Latitude == 0 {
			query = "select * from businesses where price = ? and term ~ ? and locale ~ ? and location ~ ? and open_now limit ? offset ?"
			err := r.Database.Raw(query, req.Price, req.Term, req.Locale, req.Location, req.OpenNow, req.Limit, offset).Scan(&resp).Error
			return resp, err
		}
		err := r.Database.Raw(query, req.Term, req.Locale, req.Price, req.Longitude, req.Latitude, req.OpenNow, req.Limit, offset).Scan(&resp).Error
		return resp, err
	}

	if req.Longitude == 0 && req.Latitude == 0 {
		query = "select * from businesses where term ~ ? and locale ~ ? and location ~ ? and open_now limit ? offset ?"
		err := r.Database.Raw(query, req.Term, req.Locale, req.Location, req.OpenNow, req.Limit, offset).Scan(&resp).Error
		return resp, err
	}

	err := r.Database.Raw(query, req.Term, req.Locale, req.Longitude, req.Latitude, req.OpenNow, req.Limit, offset).Scan(&resp).Error
	return resp, err
}
