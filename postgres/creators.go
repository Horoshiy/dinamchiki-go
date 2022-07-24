package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type CreatorsRepo struct {
	DB *pg.DB
}

func (r *CreatorsRepo) GetCreators(filter *models.CreatorFilter, first, last *int, after, before *string) (*models.CreatorConnection, error) {
	var items []*models.Creator
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
	var edges []*models.CreatorEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.CreatorEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.CreatorEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.CreatorEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.CreatorEdge, countElems)
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
				edges[count] = &models.CreatorEdge{
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
				edges[count] = &models.CreatorEdge{
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
				edges[count] = &models.CreatorEdge{
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

	pc := &models.CreatorConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *CreatorsRepo) GetCreatorByFiled(field, value string) (*models.Creator, error) {
	var item models.Creator
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *CreatorsRepo) GetCreatorByID(id string) (*models.Creator, error) {
	return r.GetCreatorByFiled("id", id)
}

func (r *CreatorsRepo) GetCreatorByTitle(title string) (*models.Creator, error) {
	return r.GetCreatorByFiled("title", title)
}

func (r *CreatorsRepo) CreatorSave(place *models.Creator) (*models.CreatorPayload, error) {
	_, err := r.DB.Model(place).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.CreatorPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CreatorsRepo) CreatorUpdate(place *models.Creator) (*models.CreatorPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.CreatorPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CreatorsRepo) CreatorPublish(place *models.Creator) (*models.CreatorPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.CreatorPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CreatorsRepo) CreatorRestore(place *models.Creator) (*models.CreatorPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.CreatorPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CreatorsRepo) CreatorDelete(place *models.Creator) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
