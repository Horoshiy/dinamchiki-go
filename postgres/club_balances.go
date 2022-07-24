package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type ClubBalancesRepo struct {
	DB *pg.DB
}

func (r *ClubBalancesRepo) GetClubBalances(filter *models.ClubBalanceFilter, first, last *int, after, before *string) (*models.ClubBalanceConnection, error) {
	var items []*models.ClubBalance
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
	var edges []*models.ClubBalanceEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.ClubBalanceEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.ClubBalanceEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.ClubBalanceEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.ClubBalanceEdge, countElems)
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
				edges[count] = &models.ClubBalanceEdge{
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
				edges[count] = &models.ClubBalanceEdge{
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
				edges[count] = &models.ClubBalanceEdge{
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

	pc := &models.ClubBalanceConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *ClubBalancesRepo) GetClubBalanceByFiled(field, value string) (*models.ClubBalance, error) {
	var item models.ClubBalance
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *ClubBalancesRepo) GetClubBalanceByID(id string) (*models.ClubBalance, error) {
	return r.GetClubBalanceByFiled("id", id)
}

func (r *ClubBalancesRepo) GetClubBalanceByTitle(title string) (*models.ClubBalance, error) {
	return r.GetClubBalanceByFiled("title", title)
}

func (r *ClubBalancesRepo) ClubBalanceSave(place *models.ClubBalance) (*models.ClubBalancePayload, error) {
	_, err := r.DB.Model(place).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.ClubBalancePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *ClubBalancesRepo) ClubBalanceUpdate(place *models.ClubBalance) (*models.ClubBalancePayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.ClubBalancePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *ClubBalancesRepo) ClubBalancePublish(place *models.ClubBalance) (*models.ClubBalancePayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.ClubBalancePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *ClubBalancesRepo) ClubBalanceRestore(place *models.ClubBalance) (*models.ClubBalancePayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.ClubBalancePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *ClubBalancesRepo) ClubBalanceDelete(place *models.ClubBalance) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
