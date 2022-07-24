package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type TeamsRepo struct {
	DB *pg.DB
}

func (r *TeamsRepo) GetTeams(filter *models.TeamFilter, first, last *int, after, before *string) (*models.TeamConnection, error) {
	var items []*models.Team
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
	var edges []*models.TeamEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.TeamEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.TeamEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.TeamEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.TeamEdge, countElems)
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
				edges[count] = &models.TeamEdge{
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
				edges[count] = &models.TeamEdge{
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
				edges[count] = &models.TeamEdge{
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

	pc := &models.TeamConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *TeamsRepo) All() ([]*models.TeamDto, error) {
	var items []*models.Team
	query := r.DB.Model(&items).Column("id", "name")
	err := query.Select()
	if err != nil {
		return nil, err
	}
	var arr []*models.TeamDto
	for _, v := range items {
		arr = append(arr, &models.TeamDto{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return arr, nil
}

func (r *TeamsRepo) GetTeamsByFiled(field, value string) (*models.Team, error) {
	var item models.Team
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *TeamsRepo) GetTeamsByID(id string) (*models.Team, error) {
	return r.GetTeamsByFiled("id", id)
}

func (r *TeamsRepo) GetTeamsByName(name string) (*models.Team, error) {
	return r.GetTeamsByFiled("name", name)
}

func (r *TeamsRepo) TeamsSave(stadium *models.Team) (*models.TeamPayload, error) {
	_, err := r.DB.Model(stadium).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.TeamPayload{
		Record:   stadium,
		RecordID: stadium.ID,
	}
	return placePayload, err
}

func (r *TeamsRepo) TeamsUpdate(place *models.Team) (*models.TeamPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.TeamPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TeamsRepo) TeamsPublish(place *models.Team) (*models.TeamPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.TeamPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TeamsRepo) TeamsRestore(place *models.Team) (*models.TeamPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.TeamPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *TeamsRepo) TeamsDelete(place *models.Team) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}

func (r *TeamsRepo) GetCoachesForTeam(obj *models.Team) ([]*models.Staff, error) {
	var coaches []*models.Staff
	err := r.DB.Model(&coaches).Where("id in (?)", pg.In(obj.CoachIds)).Select()
	return coaches, err
}
