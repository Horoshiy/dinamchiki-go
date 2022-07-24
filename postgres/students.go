package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type StudentsRepo struct {
	DB *pg.DB
}

func (r *StudentsRepo) GetStudents(filter *models.StudentFilter, first, last *int, after, before *string) (*models.StudentConnection, error) {
	var items []*models.Student
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
	var edges []*models.StudentEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.StudentEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.StudentEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.StudentEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.StudentEdge, countElems)
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
				edges[count] = &models.StudentEdge{
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
				edges[count] = &models.StudentEdge{
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
				edges[count] = &models.StudentEdge{
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

	pc := &models.StudentConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *StudentsRepo) All() ([]*models.StudentDto, error) {
	var items []*models.Student
	query := r.DB.Model(&items)
	err := query.Select()
	if err != nil {
		return nil, err
	}
	var arr []*models.StudentDto
	for _, v := range items {
		arr = append(arr, &models.StudentDto{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return arr, nil
}

func (r *StudentsRepo) GetStudentsByFiled(field, value string) (*models.Student, error) {
	var item models.Student
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *StudentsRepo) GetStudentsByID(id string) (*models.Student, error) {
	return r.GetStudentsByFiled("id", id)
}

func (r *StudentsRepo) StudentsSave(item *models.Student) (*models.StudentPayload, error) {
	_, err := r.DB.Model(item).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.StudentPayload{
		Record:   item,
		RecordID: item.ID,
	}
	return placePayload, err
}

func (r *StudentsRepo) StudentsUpdate(place *models.Student) (*models.StudentPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.StudentPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *StudentsRepo) StudentsPublish(place *models.Student) (*models.StudentPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.StudentPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *StudentsRepo) StudentsRestore(place *models.Student) (*models.StudentPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.StudentPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *StudentsRepo) StudentsDelete(place *models.Student) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}

func (r *StudentsRepo) GetTeamsForStudent(obj *models.Student) ([]*models.Team, error) {
	var items []*models.Team
	err := r.DB.Model(&items).Where("id in (?)", pg.In(obj.TeamIds)).Select()
	return items, err
}

func (r *StudentsRepo) GetCreatorsForStudent(obj *models.Student) ([]*models.Creator, error) {
	var items []*models.Creator
	err := r.DB.Model(&items).Where("id in (?)", pg.In(obj.CreatorIds)).Select()
	return items, err
}
