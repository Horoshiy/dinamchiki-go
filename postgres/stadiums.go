package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type StadiumsRepo struct {
	DB *pg.DB
}

func (r *StadiumsRepo) GetStadiums(filter *models.StadiumFilter, first, last *int, after, before *string) (*models.StadiumConnection, error) {
	var items []*models.Stadium
	query := r.DB.Model(&items)

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
	var edges []*models.StadiumEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.StadiumEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.StadiumEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.StadiumEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.StadiumEdge, countElems)
	}

	count := 0
	currentPage := false

	if decodedCursor == "" {
		currentPage = true
	}
	hasNextPage := false

	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query.Where("title ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
		}
	}

	err = query.Select()
	if err != nil {
		return nil, err
	}

	for i, v := range items {
		if v.ID == decodedCursor {
			currentPage = true
		}
		if err != nil {
			return nil, err
		}
		if first != nil {
			if currentPage && count < *first {
				edges[count] = &models.StadiumEdge{
					Cursor: base64.StdEncoding.EncodeToString([]byte(v.ID)),
					Node:   v,
				}
				count++
			}

			if count == *first && i < len(items) {
				hasNextPage = true
			}
		}
		if last != nil && first == nil {
			if currentPage && count < *last {
				edges[count] = &models.StadiumEdge{
					Cursor: base64.StdEncoding.EncodeToString([]byte(v.ID)),
					Node:   v,
				}
				count++
			}

			if count == *last && i < len(items) {
				hasNextPage = true
			}
		}
		if last == nil && first == nil {
			if currentPage && count < countElems {
				edges[count] = &models.StadiumEdge{
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

	pc := &models.StadiumConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *StadiumsRepo) All() ([]*models.StadiumDto, error) {
	var items []*models.Stadium
	query := r.DB.Model(&items).Column("id", "name")
	err := query.Select()
	if err != nil {
		return nil, err
	}
	var arr []*models.StadiumDto
	for _, v := range items {
		arr = append(arr, &models.StadiumDto{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return arr, nil
}

func (r *StadiumsRepo) GetStadiumByFiled(field, value string) (*models.Stadium, error) {
	var item models.Stadium
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *StadiumsRepo) GetStadiumByID(id string) (*models.Stadium, error) {
	return r.GetStadiumByFiled("id", id)
}

func (r *StadiumsRepo) GetStadiumByName(name string) (*models.Stadium, error) {
	return r.GetStadiumByFiled("name", name)
}

func (r *StadiumsRepo) StadiumSave(stadium *models.Stadium) (*models.StadiumPayload, error) {
	_, err := r.DB.Model(stadium).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.StadiumPayload{
		Record:   stadium,
		RecordID: stadium.ID,
	}
	return placePayload, err
}

func (r *StadiumsRepo) StadiumUpdate(place *models.Stadium) (*models.StadiumPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.StadiumPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *StadiumsRepo) StadiumPublish(place *models.Stadium) (*models.StadiumPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.StadiumPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *StadiumsRepo) StadiumRestore(place *models.Stadium) (*models.StadiumPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.StadiumPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *StadiumsRepo) StadiumDelete(place *models.Stadium) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
