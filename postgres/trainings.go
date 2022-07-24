package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type TrainingsRepo struct {
	DB *pg.DB
}

func (r *TrainingsRepo) GetTrainings(filter *models.TrainingFilter, first, last *int, after, before *string) (*models.TrainingConnection, error) {
	var items []*models.Training
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
	var edges []*models.TrainingEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.TrainingEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.TrainingEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.TrainingEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.TrainingEdge, countElems)
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
				edges[count] = &models.TrainingEdge{
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
				edges[count] = &models.TrainingEdge{
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
				edges[count] = &models.TrainingEdge{
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

	pc := &models.TrainingConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *TrainingsRepo) All() ([]*models.TrainingDto, error) {
	var items []*models.Training
	query := r.DB.Model(&items).
		ColumnExpr("trainings.*").
		ColumnExpr("squad.id AS squad__id").
		ColumnExpr("squad.name AS squad__name").
		Join("LEFT JOIN teams AS squad ON squad.id = trainings.team")
	err := query.Select()
	if err != nil {
		return nil, err
	}
	var arr []*models.TrainingDto
	for _, v := range items {
		arr = append(arr, &models.TrainingDto{
			ID:   v.ID,
			Name: v.Squad.Name,
		})
	}
	return arr, nil
}

func (r *TrainingsRepo) GetTrainingsByFiled(field, value string) (*models.Training, error) {
	var item models.Training
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *TrainingsRepo) GetTrainingsByID(id string) (*models.Training, error) {
	return r.GetTrainingsByFiled("id", id)
}

func (r *TrainingsRepo) TrainingsSave(item *models.Training) (*models.TrainingPayload, error) {
	_, err := r.DB.Model(item).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.TrainingPayload{
		Record:   item,
		RecordID: item.ID,
	}
	return placePayload, err
}

func (r *TrainingsRepo) TrainingsUpdate(place *models.Training) (*models.TrainingPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.TrainingPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TrainingsRepo) TrainingsPublish(place *models.Training) (*models.TrainingPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.TrainingPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TrainingsRepo) TrainingsRestore(place *models.Training) (*models.TrainingPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.TrainingPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TrainingsRepo) TrainingsDelete(place *models.Training) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}

func (r *TrainingsRepo) GetCoachesForTraining(obj *models.Training) ([]*models.Staff, error) {
	var coaches []*models.Staff
	err := r.DB.Model(&coaches).Where("id in (?)", pg.In(obj.CoachIds)).Select()
	return coaches, err
}

func (r *TrainingsRepo) GetCoachesForTeam(obj *models.Training) ([]*models.Staff, error) {
	var coaches []*models.Staff
	err := r.DB.Model(&coaches).Where("id in (?)", pg.In(obj.CoachIds)).Select()
	return coaches, err
}
