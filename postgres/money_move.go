package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type MoneyMovesRepo struct {
	DB *pg.DB
}

func (r *MoneyMovesRepo) GetMoneyMoves(filter *models.MoneyMoveFilter, first, last *int, after, before *string) (*models.MoneyMoveConnection, error) {
	var items []*models.MoneyMove
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
	var edges []*models.MoneyMoveEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.MoneyMoveEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.MoneyMoveEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.MoneyMoveEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.MoneyMoveEdge, countElems)
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
				edges[count] = &models.MoneyMoveEdge{
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
				edges[count] = &models.MoneyMoveEdge{
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
				edges[count] = &models.MoneyMoveEdge{
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

	pc := &models.MoneyMoveConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *MoneyMovesRepo) GetMoneyMoveByFiled(field, value string) (*models.MoneyMove, error) {
	var item models.MoneyMove
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *MoneyMovesRepo) GetMoneyMoveByID(id string) (*models.MoneyMove, error) {
	return r.GetMoneyMoveByFiled("id", id)
}

func (r *MoneyMovesRepo) GetMoneyMoveByTitle(title string) (*models.MoneyMove, error) {
	return r.GetMoneyMoveByFiled("title", title)
}

func (r *MoneyMovesRepo) MoneyMoveSave(item *models.MoneyMove) (*models.MoneyMovePayload, error) {
	_, err := r.DB.Model(item).Returning("*").Insert()
	fmt.Println(err)
	payload := &models.MoneyMovePayload{
		Record:   item,
		RecordID: item.ID,
	}
	return payload, err
}

func (r *MoneyMovesRepo) MoneyMoveUpdate(place *models.MoneyMove) (*models.MoneyMovePayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.MoneyMovePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *MoneyMovesRepo) MoneyMovePublish(place *models.MoneyMove) (*models.MoneyMovePayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.MoneyMovePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *MoneyMovesRepo) MoneyMoveRestore(place *models.MoneyMove) (*models.MoneyMovePayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.MoneyMovePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *MoneyMovesRepo) MoneyMoveDelete(place *models.MoneyMove) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
