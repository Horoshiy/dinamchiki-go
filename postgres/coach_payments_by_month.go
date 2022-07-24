package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type CoachPaymentByMonthRepo struct {
	DB *pg.DB
}

func (r *CoachPaymentByMonthRepo) GetCoachPaymentByMonth(filter *models.CoachPaymentByMonthFilter, first, last *int, after, before *string) (*models.CoachPaymentByMonthConnection, error) {
	var items []*models.CoachPaymentByMonth
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
	var edges []*models.CoachPaymentByMonthEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.CoachPaymentByMonthEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.CoachPaymentByMonthEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.CoachPaymentByMonthEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.CoachPaymentByMonthEdge, countElems)
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
				edges[count] = &models.CoachPaymentByMonthEdge{
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
				edges[count] = &models.CoachPaymentByMonthEdge{
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
				edges[count] = &models.CoachPaymentByMonthEdge{
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

	pc := &models.CoachPaymentByMonthConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *CoachPaymentByMonthRepo) GetCoachPaymentByMonthByFiled(field, value string) (*models.CoachPaymentByMonth, error) {
	var item models.CoachPaymentByMonth
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *CoachPaymentByMonthRepo) GetCoachPaymentByMonthByID(id string) (*models.CoachPaymentByMonth, error) {
	return r.GetCoachPaymentByMonthByFiled("id", id)
}

func (r *CoachPaymentByMonthRepo) GetCoachPaymentByMonthByTitle(title string) (*models.CoachPaymentByMonth, error) {
	return r.GetCoachPaymentByMonthByFiled("title", title)
}

func (r *CoachPaymentByMonthRepo) CoachPaymentByMonthSave(place *models.CoachPaymentByMonth) (*models.CoachPaymentByMonthPayload, error) {
	_, err := r.DB.Model(place).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.CoachPaymentByMonthPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CoachPaymentByMonthRepo) CoachPaymentByMonthUpdate(place *models.CoachPaymentByMonth) (*models.CoachPaymentByMonthPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.CoachPaymentByMonthPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CoachPaymentByMonthRepo) CoachPaymentByMonthPublish(place *models.CoachPaymentByMonth) (*models.CoachPaymentByMonthPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.CoachPaymentByMonthPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CoachPaymentByMonthRepo) CoachPaymentByMonthRestore(place *models.CoachPaymentByMonth) (*models.CoachPaymentByMonthPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.CoachPaymentByMonthPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CoachPaymentByMonthRepo) CoachPaymentByMonthDelete(place *models.CoachPaymentByMonth) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
