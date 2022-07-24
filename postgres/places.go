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
	if before != nil && after == nil {
		b, err := base64.StdEncoding.DecodeString(*before)
		if err != nil {
			return nil, err
		}
		decodedCursor = string(b)
	}
	var edges []*models.PlaceEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.PlaceEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.PlaceEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.PlaceEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.PlaceEdge, countElems)
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

	err = query.Select()
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
		}
		if last != nil && first == nil {
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
		if last == nil && first == nil {
			if currentPage && count < countElems {
				edges[count] = &models.PlaceEdge{
					Cursor: base64.StdEncoding.EncodeToString([]byte(v.ID)),
					Node:   v,
				}
				count++
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

func (r *PlacesRepo) All() ([]*models.PlaceDto, error) {
	var items []*models.Place
	query := r.DB.Model(&items).Column("id", "name")
	err := query.Select()
	if err != nil {
		return nil, err
	}
	var arr []*models.PlaceDto
	for _, v := range items {
		arr = append(arr, &models.PlaceDto{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return arr, nil
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

func (r *PlacesRepo) PlaceSave(place *models.Place) (*models.PlacePayload, error) {
	_, err := r.DB.Model(place).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.PlacePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *PlacesRepo) PlaceUpdate(place *models.Place) (*models.PlacePayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.PlacePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *PlacesRepo) PlacePublish(place *models.Place) (*models.PlacePayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.PlacePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *PlacesRepo) PlaceRestore(place *models.Place) (*models.PlacePayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.PlacePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *PlacesRepo) PlaceDelete(place *models.Place) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
