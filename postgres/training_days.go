package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type TrainingDaysRepo struct {
	DB *pg.DB
}

func (r *TrainingDaysRepo) GetTrainingDays(filter *models.TrainingDayFilter, first, last *int, after, before *string) (*models.TrainingDayConnection, error) {
	var items []*models.TrainingDay
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
	var edges []*models.TrainingDayEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.TrainingDayEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.TrainingDayEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.TrainingDayEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.TrainingDayEdge, countElems)
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
				edges[count] = &models.TrainingDayEdge{
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
				edges[count] = &models.TrainingDayEdge{
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
				edges[count] = &models.TrainingDayEdge{
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

	pc := &models.TrainingDayConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *TrainingDaysRepo) GetTrainingDaysByFiled(field, value string) (*models.TrainingDay, error) {
	var item models.TrainingDay
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *TrainingDaysRepo) GetTrainingDaysByID(id string) (*models.TrainingDay, error) {
	return r.GetTrainingDaysByFiled("id", id)
}

func (r *TrainingDaysRepo) TrainingDaysSave(item *models.TrainingDay) (*models.TrainingDayPayload, error) {
	_, err := r.DB.Model(item).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.TrainingDayPayload{
		Record:   item,
		RecordID: item.ID,
	}
	return placePayload, err
}

func (r *TrainingDaysRepo) TrainingDaysUpdate(place *models.TrainingDay) (*models.TrainingDayPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.TrainingDayPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TrainingDaysRepo) TrainingDaysPublish(place *models.TrainingDay) (*models.TrainingDayPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.TrainingDayPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TrainingDaysRepo) TrainingDaysRestore(place *models.TrainingDay) (*models.TrainingDayPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.TrainingDayPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TrainingDaysRepo) TrainingDaysDelete(place *models.TrainingDay) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
