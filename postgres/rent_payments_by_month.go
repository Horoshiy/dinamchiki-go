package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type RentPaymentByMonthRepo struct {
	DB *pg.DB
}

func (r *RentPaymentByMonthRepo) GetRentPaymentByMonth(filter *models.RentPaymentByMonthFilter, first, last *int, after, before *string) (*models.RentPaymentByMonthConnection, error) {
	var items []*models.RentPaymentByMonth
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
	var edges []*models.RentPaymentByMonthEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.RentPaymentByMonthEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.RentPaymentByMonthEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.RentPaymentByMonthEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.RentPaymentByMonthEdge, countElems)
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
				edges[count] = &models.RentPaymentByMonthEdge{
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
				edges[count] = &models.RentPaymentByMonthEdge{
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
				edges[count] = &models.RentPaymentByMonthEdge{
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

	pc := &models.RentPaymentByMonthConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *RentPaymentByMonthRepo) GetRentPaymentByMonthByFiled(field, value string) (*models.RentPaymentByMonth, error) {
	var item models.RentPaymentByMonth
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *RentPaymentByMonthRepo) GetRentPaymentByMonthByID(id string) (*models.RentPaymentByMonth, error) {
	return r.GetRentPaymentByMonthByFiled("id", id)
}

func (r *RentPaymentByMonthRepo) GetRentPaymentByMonthByTitle(title string) (*models.RentPaymentByMonth, error) {
	return r.GetRentPaymentByMonthByFiled("title", title)
}

func (r *RentPaymentByMonthRepo) RentPaymentByMonthSave(place *models.RentPaymentByMonth) (*models.RentPaymentByMonthPayload, error) {
	_, err := r.DB.Model(place).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.RentPaymentByMonthPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *RentPaymentByMonthRepo) RentPaymentByMonthUpdate(place *models.RentPaymentByMonth) (*models.RentPaymentByMonthPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.RentPaymentByMonthPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *RentPaymentByMonthRepo) RentPaymentByMonthPublish(place *models.RentPaymentByMonth) (*models.RentPaymentByMonthPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.RentPaymentByMonthPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *RentPaymentByMonthRepo) RentPaymentByMonthRestore(place *models.RentPaymentByMonth) (*models.RentPaymentByMonthPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.RentPaymentByMonthPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *RentPaymentByMonthRepo) RentPaymentByMonthDelete(place *models.RentPaymentByMonth) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
