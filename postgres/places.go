package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type PlacesRepo struct {
	DB *pg.DB
}

func (r *PlacesRepo) GetPlaces(filter *models.PlaceFilter, first, last *int, after, before *string) (*models.PlaceConnection, error) {
	var places []*models.Place
	query := r.DB.Model(&places)

	var decodedCursor string
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return nil, err
		}
		decodedCursor = string(b)
	}
	if before != nil {
		b, err := base64.StdEncoding.DecodeString(*before)
		if err != nil {
			return nil, err
		}
		decodedCursor = string(b)
	}
	var edges []*models.PlaceEdge
	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.PlaceEdge, *first)
	} else {
		query.Order("id DESC")
		edges = make([]*models.PlaceEdge, *last)
	}
	count := 0
	currentPage := false

	if decodedCursor == "" {
		currentPage = true
	}
	hasNextPage := false

	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
		}
	}

	err := query.Select()
	if err != nil {
		return nil, err
	}

	for i, v := range places {
		if v.ID == decodedCursor {
			currentPage = true
		}

		if first != nil {
			if currentPage && count < *first {
				edges[count] = &models.PlaceEdge{
					Cursor: base64.StdEncoding.EncodeToString([]byte(v.ID)),
					Node:   v,
				}
				count++
			}

			if count == *first && i < len(places) {
				hasNextPage = true
			}
		} else {
			if currentPage && count < *last {
				edges[count] = &models.PlaceEdge{
					Cursor: base64.StdEncoding.EncodeToString([]byte(v.ID)),
					Node:   v,
				}
				count++
			}

			if count == *last && i < len(places) {
				hasNextPage = true
			}
		}
	}

	pageInfo := models.PageInfo{
		StartCursor: base64.StdEncoding.EncodeToString([]byte(edges[0].Node.ID)),
		EndCursor:   base64.StdEncoding.EncodeToString([]byte(edges[count-1].Node.ID)),
		HasNextPage: &hasNextPage,
	}

	pc := &models.PlaceConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
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

func (r *PlacesRepo) CreatePlace(place *models.Place) (*models.PlacePayload, error) {
	_, err := r.DB.Model(place).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.PlacePayload{
		Place: &models.Place{
			ID:          place.ID,
			Address:     place.Address,
			Description: place.Description,
			Name:        place.Name,
			OrderNumber: place.OrderNumber,
			Published:   place.Published,
		},
		ID: place.ID,
	}
	return placePayload, err
}

func (r *PlacesRepo) UpdatePlace(place *models.Place) (*models.Place, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	return place, err
}

func (r *PlacesRepo) DeletePlace(place *models.Place) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
