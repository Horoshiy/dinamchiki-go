package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type CoachPaymentByTeamRepo struct {
	DB *pg.DB
}

func (r *CoachPaymentByTeamRepo) GetCoachPaymentByTeam(filter *models.CoachPaymentByTeamFilter, first, last *int, after, before *string) (*models.CoachPaymentByTeamConnection, error) {
	var items []*models.CoachPaymentByTeam
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
	var edges []*models.CoachPaymentByTeamEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.CoachPaymentByTeamEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.CoachPaymentByTeamEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.CoachPaymentByTeamEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.CoachPaymentByTeamEdge, countElems)
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
				edges[count] = &models.CoachPaymentByTeamEdge{
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
				edges[count] = &models.CoachPaymentByTeamEdge{
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
				edges[count] = &models.CoachPaymentByTeamEdge{
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

	pc := &models.CoachPaymentByTeamConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *CoachPaymentByTeamRepo) GetCoachPaymentByTeamByFiled(field, value string) (*models.CoachPaymentByTeam, error) {
	var item models.CoachPaymentByTeam
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *CoachPaymentByTeamRepo) GetCoachPaymentByTeamByID(id string) (*models.CoachPaymentByTeam, error) {
	return r.GetCoachPaymentByTeamByFiled("id", id)
}

func (r *CoachPaymentByTeamRepo) GetCoachPaymentByTeamByTitle(title string) (*models.CoachPaymentByTeam, error) {
	return r.GetCoachPaymentByTeamByFiled("title", title)
}

func (r *CoachPaymentByTeamRepo) CoachPaymentByTeamSave(place *models.CoachPaymentByTeam) (*models.CoachPaymentByTeamPayload, error) {
	_, err := r.DB.Model(place).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.CoachPaymentByTeamPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CoachPaymentByTeamRepo) CoachPaymentByTeamUpdate(place *models.CoachPaymentByTeam) (*models.CoachPaymentByTeamPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.CoachPaymentByTeamPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CoachPaymentByTeamRepo) CoachPaymentByTeamPublish(place *models.CoachPaymentByTeam) (*models.CoachPaymentByTeamPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.CoachPaymentByTeamPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CoachPaymentByTeamRepo) CoachPaymentByTeamRestore(place *models.CoachPaymentByTeam) (*models.CoachPaymentByTeamPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.CoachPaymentByTeamPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CoachPaymentByTeamRepo) CoachPaymentByTeamDelete(place *models.CoachPaymentByTeam) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
