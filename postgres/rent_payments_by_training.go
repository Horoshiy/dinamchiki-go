package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type RentPaymentByTrainingRepo struct {
	DB *pg.DB
}

func (r *RentPaymentByTrainingRepo) GetRentPaymentByTraining(filter *models.RentPaymentByTrainingFilter, first, last *int, after, before *string) (*models.RentPaymentByTrainingConnection, error) {
	var items []*models.RentPaymentByTraining
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
	var edges []*models.RentPaymentByTrainingEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.RentPaymentByTrainingEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.RentPaymentByTrainingEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.RentPaymentByTrainingEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.RentPaymentByTrainingEdge, countElems)
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
				edges[count] = &models.RentPaymentByTrainingEdge{
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
				edges[count] = &models.RentPaymentByTrainingEdge{
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
				edges[count] = &models.RentPaymentByTrainingEdge{
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

	pc := &models.RentPaymentByTrainingConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *RentPaymentByTrainingRepo) GetRentPaymentByTrainingByFiled(field, value string) (*models.RentPaymentByTraining, error) {
	var item models.RentPaymentByTraining
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *RentPaymentByTrainingRepo) GetRentPaymentByTrainingByID(id string) (*models.RentPaymentByTraining, error) {
	return r.GetRentPaymentByTrainingByFiled("id", id)
}

func (r *RentPaymentByTrainingRepo) GetRentPaymentByTrainingByTitle(title string) (*models.RentPaymentByTraining, error) {
	return r.GetRentPaymentByTrainingByFiled("title", title)
}

func (r *RentPaymentByTrainingRepo) RentPaymentByTrainingSave(place *models.RentPaymentByTraining) (*models.RentPaymentByTrainingPayload, error) {
	_, err := r.DB.Model(place).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.RentPaymentByTrainingPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *RentPaymentByTrainingRepo) RentPaymentByTrainingUpdate(place *models.RentPaymentByTraining) (*models.RentPaymentByTrainingPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.RentPaymentByTrainingPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *RentPaymentByTrainingRepo) RentPaymentByTrainingPublish(place *models.RentPaymentByTraining) (*models.RentPaymentByTrainingPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.RentPaymentByTrainingPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *RentPaymentByTrainingRepo) RentPaymentByTrainingRestore(place *models.RentPaymentByTraining) (*models.RentPaymentByTrainingPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.RentPaymentByTrainingPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *RentPaymentByTrainingRepo) RentPaymentByTrainingDelete(place *models.RentPaymentByTraining) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}

func (r *RentPaymentByTrainingRepo) GetTrainingsForRentPaymentByTraining(obj *models.RentPaymentByTraining) ([]*models.Training, error) {
	var items []*models.Training
	err := r.DB.Model(&items).Where("id in (?)", pg.In(obj.TrainingIds)).Select()
	return items, err
}
