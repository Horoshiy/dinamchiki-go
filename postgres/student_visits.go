package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type StudentVisitsRepo struct {
	DB *pg.DB
}

func (r *StudentVisitsRepo) GetStudentVisits(filter *models.StudentVisitFilter, first, last *int, after, before *string) (*models.StudentVisitConnection, error) {
	var items []*models.StudentVisit
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
	var edges []*models.StudentVisitEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.StudentVisitEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.StudentVisitEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.StudentVisitEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.StudentVisitEdge, countElems)
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
				edges[count] = &models.StudentVisitEdge{
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
				edges[count] = &models.StudentVisitEdge{
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
				edges[count] = &models.StudentVisitEdge{
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

	pc := &models.StudentVisitConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *StudentVisitsRepo) GetStudentVisitByFiled(field, value string) (*models.StudentVisit, error) {
	var item models.StudentVisit
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *StudentVisitsRepo) GetStudentVisitByID(id string) (*models.StudentVisit, error) {
	return r.GetStudentVisitByFiled("id", id)
}

func (r *StudentVisitsRepo) GetStudentVisitByTitle(title string) (*models.StudentVisit, error) {
	return r.GetStudentVisitByFiled("title", title)
}

func (r *StudentVisitsRepo) StudentVisitSave(item *models.StudentVisit) (*models.StudentVisitPayload, error) {
	_, err := r.DB.Model(item).Returning("*").Insert()
	fmt.Println(err)
	payload := &models.StudentVisitPayload{
		Record:   item,
		RecordID: item.ID,
	}
	return payload, err
}

func (r *StudentVisitsRepo) StudentVisitUpdate(place *models.StudentVisit) (*models.StudentVisitPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.StudentVisitPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *StudentVisitsRepo) StudentVisitPublish(place *models.StudentVisit) (*models.StudentVisitPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.StudentVisitPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *StudentVisitsRepo) StudentVisitRestore(place *models.StudentVisit) (*models.StudentVisitPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.StudentVisitPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *StudentVisitsRepo) StudentVisitDelete(place *models.StudentVisit) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
