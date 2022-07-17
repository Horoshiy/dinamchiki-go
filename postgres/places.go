package postgres

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type PlacesRepo struct {
	DB *pg.DB
}

func (r *PlacesRepo) GetPlaces(filter *models.PlaceFilter, limit, offset *int) ([]*models.Place, error) {
	var places []*models.Place

	query := r.DB.Model(&places).Order("id")

	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
		}
	}

	if limit != nil {
		query.Limit(*limit)
	}

	if offset != nil {
		query.Offset(*offset)
	}

	err := query.Select()
	if err != nil {
		return nil, err
	}
	return places, nil
}

func (r *PlacesRepo) GetPlaceByFiled(field, value string) (*models.Place, error) {
	var place models.Place
	err := r.DB.Model(&place).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &place, err
}

func (r *PlacesRepo) GetPlaceByID(id string) (*models.Place, error) {
	return r.GetPlaceByFiled("id", id)
}

func (r *PlacesRepo) GetPlaceByName(name string) (*models.Place, error) {
	return r.GetPlaceByFiled("name", name)
}

func (r *PlacesRepo) CreatePlace(place *models.Place) (*models.Place, error) {
	_, err := r.DB.Model(place).Returning("*").Insert()
	return place, err
}

func (r *PlacesRepo) UpdatePlace(place *models.Place) (*models.Place, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	return place, err
}

func (r *PlacesRepo) DeletePlace(place *models.Place) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
