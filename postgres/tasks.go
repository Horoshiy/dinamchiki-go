package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type TasksRepo struct {
	DB *pg.DB
}

func (r *TasksRepo) GetTasks(filter *models.TaskFilter, first, last *int, after, before *string) (*models.TaskConnection, error) {
	var items []*models.Task
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
	var edges []*models.TaskEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.TaskEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.TaskEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.TaskEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.TaskEdge, countElems)
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
				edges[count] = &models.TaskEdge{
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
				edges[count] = &models.TaskEdge{
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
				edges[count] = &models.TaskEdge{
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

	pc := &models.TaskConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *TasksRepo) GetTasksByFiled(field, value string) (*models.Task, error) {
	var item models.Task
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *TasksRepo) GetTasksByID(id string) (*models.Task, error) {
	return r.GetTasksByFiled("id", id)
}

func (r *TasksRepo) TasksSave(item *models.Task) (*models.TaskPayload, error) {
	_, err := r.DB.Model(item).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.TaskPayload{
		Record:   item,
		RecordID: item.ID,
	}
	return placePayload, err
}

func (r *TasksRepo) TasksUpdate(place *models.Task) (*models.TaskPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.TaskPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TasksRepo) TasksPublish(place *models.Task) (*models.TaskPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.TaskPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TasksRepo) TasksRestore(place *models.Task) (*models.TaskPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.TaskPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TasksRepo) TasksDelete(place *models.Task) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}

func (r *TasksRepo) GetLeadsForTask(obj *models.Task) ([]*models.Lead, error) {
	var items []*models.Lead
	err := r.DB.Model(&items).Where("id in (?)", pg.In(obj.LeadIds)).Select()
	return items, err
}

func (r *TasksRepo) GetStudentsForTask(obj *models.Task) ([]*models.Student, error) {
	var items []*models.Student
	err := r.DB.Model(&items).Where("id in (?)", pg.In(obj.StudentIds)).Select()
	return items, err
}

func (r *TasksRepo) GetWorkersForTask(obj *models.Task) ([]*models.Staff, error) {
	var items []*models.Staff
	err := r.DB.Model(&items).Where("id in (?)", pg.In(obj.WorkerIds)).Select()
	return items, err
}
