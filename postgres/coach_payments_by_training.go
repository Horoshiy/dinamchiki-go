package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type CoachPaymentByTrainingRepo struct {
	DB *pg.DB
}

func (r *CoachPaymentByTrainingRepo) GetCoachPaymentByTraining(filter *models.CoachPaymentByTrainingFilter, first, last *int, after, before *string) (*models.CoachPaymentByTrainingConnection, error) {
	var items []*models.CoachPaymentByTraining
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
	var edges []*models.CoachPaymentByTrainingEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.CoachPaymentByTrainingEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.CoachPaymentByTrainingEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.CoachPaymentByTrainingEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.CoachPaymentByTrainingEdge, countElems)
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
				edges[count] = &models.CoachPaymentByTrainingEdge{
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
				edges[count] = &models.CoachPaymentByTrainingEdge{
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
				edges[count] = &models.CoachPaymentByTrainingEdge{
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

	pc := &models.CoachPaymentByTrainingConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *CoachPaymentByTrainingRepo) GetCoachPaymentByTrainingByFiled(field, value string) (*models.CoachPaymentByTraining, error) {
	var item models.CoachPaymentByTraining
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *CoachPaymentByTrainingRepo) GetCoachPaymentByTrainingByID(id string) (*models.CoachPaymentByTraining, error) {
	return r.GetCoachPaymentByTrainingByFiled("id", id)
}

func (r *CoachPaymentByTrainingRepo) GetCoachPaymentByTrainingByTitle(title string) (*models.CoachPaymentByTraining, error) {
	return r.GetCoachPaymentByTrainingByFiled("title", title)
}

func (r *CoachPaymentByTrainingRepo) CoachPaymentByTrainingSave(place *models.CoachPaymentByTraining) (*models.CoachPaymentByTrainingPayload, error) {
	_, err := r.DB.Model(place).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.CoachPaymentByTrainingPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CoachPaymentByTrainingRepo) CoachPaymentByTrainingUpdate(place *models.CoachPaymentByTraining) (*models.CoachPaymentByTrainingPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.CoachPaymentByTrainingPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CoachPaymentByTrainingRepo) CoachPaymentByTrainingPublish(place *models.CoachPaymentByTraining) (*models.CoachPaymentByTrainingPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.CoachPaymentByTrainingPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CoachPaymentByTrainingRepo) CoachPaymentByTrainingRestore(place *models.CoachPaymentByTraining) (*models.CoachPaymentByTrainingPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.CoachPaymentByTrainingPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CoachPaymentByTrainingRepo) CoachPaymentByTrainingDelete(place *models.CoachPaymentByTraining) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
