package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type LeadsRepo struct {
	DB *pg.DB
}

func (r *LeadsRepo) GetLeads(filter *models.LeadFilter, first, last *int, after, before *string) (*models.LeadConnection, error) {
	var items []*models.Lead
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
	var edges []*models.LeadEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.LeadEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.LeadEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.LeadEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.LeadEdge, countElems)
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
				edges[count] = &models.LeadEdge{
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
				edges[count] = &models.LeadEdge{
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
				edges[count] = &models.LeadEdge{
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

	pc := &models.LeadConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *LeadsRepo) GetLeadsByFiled(field, value string) (*models.Lead, error) {
	var item models.Lead
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *LeadsRepo) GetLeadsByID(id string) (*models.Lead, error) {
	return r.GetLeadsByFiled("id", id)
}

func (r *LeadsRepo) LeadsSave(item *models.Lead) (*models.LeadPayload, error) {
	_, err := r.DB.Model(item).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.LeadPayload{
		Record:   item,
		RecordID: item.ID,
	}
	return placePayload, err
}

func (r *LeadsRepo) LeadsUpdate(place *models.Lead) (*models.LeadPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.LeadPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *LeadsRepo) LeadsPublish(place *models.Lead) (*models.LeadPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.LeadPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *LeadsRepo) LeadsRestore(place *models.Lead) (*models.LeadPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.LeadPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *LeadsRepo) LeadsDelete(place *models.Lead) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}

func (r *LeadsRepo) GetStudentsForLead(obj *models.Lead) ([]*models.Student, error) {
	var items []*models.Student
	err := r.DB.Model(&items).Where("id in (?)", pg.In(obj.StudentIds)).Select()
	return items, err
}
