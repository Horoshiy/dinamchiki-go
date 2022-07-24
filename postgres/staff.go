package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type StaffRepo struct {
	DB *pg.DB
}

func (r *StaffRepo) GetStaff(filter *models.StaffFilter, first, last *int, after, before *string) (*models.StaffConnection, error) {
	var items []*models.Staff
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
	var edges []*models.StaffEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.StaffEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.StaffEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.StaffEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.StaffEdge, countElems)
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
				edges[count] = &models.StaffEdge{
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
				edges[count] = &models.StaffEdge{
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
				edges[count] = &models.StaffEdge{
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

	pc := &models.StaffConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *StaffRepo) GetStaffByFiled(field, value string) (*models.Staff, error) {
	var item models.Staff
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *StaffRepo) GetStaffByID(id string) (*models.Staff, error) {
	return r.GetStaffByFiled("id", id)
}

func (r *StaffRepo) GetStaffByName(name string) (*models.Staff, error) {
	return r.GetStaffByFiled("name", name)
}

func (r *StaffRepo) StaffSave(stadium *models.Staff) (*models.StaffPayload, error) {
	_, err := r.DB.Model(stadium).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.StaffPayload{
		Record:   stadium,
		RecordID: stadium.ID,
	}
	return placePayload, err
}

func (r *StaffRepo) StaffUpdate(place *models.Staff) (*models.StaffPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.StaffPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *StaffRepo) StaffPublish(place *models.Staff) (*models.StaffPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.StaffPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *StaffRepo) StaffRestore(place *models.Staff) (*models.StaffPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.StaffPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *StaffRepo) StaffDelete(place *models.Staff) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
