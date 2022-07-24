package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type UsersRepo struct {
	DB *pg.DB
}

func (r *UsersRepo) GetUsers(filter *models.UserFilter, first, last *int, after, before *string) (*models.UserConnection, error) {
	var items []*models.User
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
	var edges []*models.UserEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.UserEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.UserEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.UserEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.UserEdge, countElems)
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
				edges[count] = &models.UserEdge{
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
				edges[count] = &models.UserEdge{
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
				edges[count] = &models.UserEdge{
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

	pc := &models.UserConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (u *UsersRepo) GetUserByFiled(field, value string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &user, err
}

func (u *UsersRepo) GetUserByID(id string) (*models.User, error) {
	return u.GetUserByFiled("id", id)

}

func (u *UsersRepo) GetUserByPhone(phone string) (*models.User, error) {
	return u.GetUserByFiled("phone", phone)
}

func (u *UsersRepo) CreateUser(tx *pg.Tx, user *models.User) (*models.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()
	return user, err
}

func (r *UsersRepo) GetUsersAll() ([]*models.UserDto, error) {
	var items []*models.User
	query := r.DB.Model(&items).Column("id", "last_name", "first_name")
	err := query.Select()
	if err != nil {
		return nil, err
	}
	var arr []*models.UserDto
	for _, v := range items {
		arr = append(arr, &models.UserDto{
			ID:   v.ID,
			Name: v.LastName + " " + v.FirstName,
		})
	}
	return arr, nil
}
