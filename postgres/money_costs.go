package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type MoneyCostsRepo struct {
	DB *pg.DB
}

func (r *MoneyCostsRepo) GetMoneyCosts(filter *models.MoneyCostFilter, first, last *int, after, before *string) (*models.MoneyCostConnection, error) {
	var items []*models.MoneyCost
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
	var edges []*models.MoneyCostEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.MoneyCostEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.MoneyCostEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.MoneyCostEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.MoneyCostEdge, countElems)
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
				edges[count] = &models.MoneyCostEdge{
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
				edges[count] = &models.MoneyCostEdge{
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
				edges[count] = &models.MoneyCostEdge{
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

	pc := &models.MoneyCostConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *MoneyCostsRepo) GetMoneyCostByFiled(field, value string) (*models.MoneyCost, error) {
	var item models.MoneyCost
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *MoneyCostsRepo) GetMoneyCostByID(id string) (*models.MoneyCost, error) {
	return r.GetMoneyCostByFiled("id", id)
}

func (r *MoneyCostsRepo) GetMoneyCostByTitle(title string) (*models.MoneyCost, error) {
	return r.GetMoneyCostByFiled("title", title)
}

func (r *MoneyCostsRepo) MoneyCostSave(item *models.MoneyCost) (*models.MoneyCostPayload, error) {
	_, err := r.DB.Model(item).Returning("*").Insert()
	fmt.Println(err)
	payload := &models.MoneyCostPayload{
		Record:   item,
		RecordID: item.ID,
	}
	return payload, err
}

func (r *MoneyCostsRepo) MoneyCostUpdate(place *models.MoneyCost) (*models.MoneyCostPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.MoneyCostPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *MoneyCostsRepo) MoneyCostPublish(place *models.MoneyCost) (*models.MoneyCostPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.MoneyCostPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *MoneyCostsRepo) MoneyCostRestore(place *models.MoneyCost) (*models.MoneyCostPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.MoneyCostPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *MoneyCostsRepo) MoneyCostDelete(place *models.MoneyCost) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
