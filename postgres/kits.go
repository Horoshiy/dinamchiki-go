package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type KitsRepo struct {
	DB *pg.DB
}

func (r *KitsRepo) GetKits(filter *models.KitFilter, first, last *int, after, before *string) (*models.KitConnection, error) {
	var items []*models.Kit
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
	var edges []*models.KitEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.KitEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.KitEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.KitEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.KitEdge, countElems)
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
				edges[count] = &models.KitEdge{
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
				edges[count] = &models.KitEdge{
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
				edges[count] = &models.KitEdge{
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

	pc := &models.KitConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *KitsRepo) GetKitsByFiled(field, value string) (*models.Kit, error) {
	var item models.Kit
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *KitsRepo) GetKitsByID(id string) (*models.Kit, error) {
	return r.GetKitsByFiled("id", id)
}

func (r *KitsRepo) KitsSave(item *models.Kit) (*models.KitPayload, error) {
	_, err := r.DB.Model(item).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.KitPayload{
		Record:   item,
		RecordID: item.ID,
	}
	return placePayload, err
}

func (r *KitsRepo) KitsUpdate(place *models.Kit) (*models.KitPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.KitPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *KitsRepo) KitsPublish(place *models.Kit) (*models.KitPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.KitPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *KitsRepo) KitsRestore(place *models.Kit) (*models.KitPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.KitPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *KitsRepo) KitsDelete(place *models.Kit) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
