package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type TeamBalancesRepo struct {
	DB *pg.DB
}

func (r *TeamBalancesRepo) GetTeamBalances(filter *models.TeamBalanceFilter, first, last *int, after, before *string) (*models.TeamBalanceConnection, error) {
	var items []*models.TeamBalance
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
	var edges []*models.TeamBalanceEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.TeamBalanceEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.TeamBalanceEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.TeamBalanceEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.TeamBalanceEdge, countElems)
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
				edges[count] = &models.TeamBalanceEdge{
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
				edges[count] = &models.TeamBalanceEdge{
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
				edges[count] = &models.TeamBalanceEdge{
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

	pc := &models.TeamBalanceConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *TeamBalancesRepo) GetTeamBalanceByFiled(field, value string) (*models.TeamBalance, error) {
	var item models.TeamBalance
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *TeamBalancesRepo) GetTeamBalanceByID(id string) (*models.TeamBalance, error) {
	return r.GetTeamBalanceByFiled("id", id)
}

func (r *TeamBalancesRepo) GetTeamBalanceByTitle(title string) (*models.TeamBalance, error) {
	return r.GetTeamBalanceByFiled("title", title)
}

func (r *TeamBalancesRepo) TeamBalanceSave(place *models.TeamBalance) (*models.TeamBalancePayload, error) {
	_, err := r.DB.Model(place).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.TeamBalancePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TeamBalancesRepo) TeamBalanceUpdate(place *models.TeamBalance) (*models.TeamBalancePayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.TeamBalancePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TeamBalancesRepo) TeamBalancePublish(place *models.TeamBalance) (*models.TeamBalancePayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.TeamBalancePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TeamBalancesRepo) TeamBalanceRestore(place *models.TeamBalance) (*models.TeamBalancePayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.TeamBalancePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TeamBalancesRepo) TeamBalanceDelete(place *models.TeamBalance) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
