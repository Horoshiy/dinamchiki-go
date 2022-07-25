package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type CartsRepo struct {
	DB *pg.DB
}

func (r *CartsRepo) GetCarts(filter *models.CartFilter, first, last *int, after, before *string) (*models.CartConnection, error) {
	var items []*models.Cart
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
	var edges []*models.CartEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.CartEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.CartEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.CartEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.CartEdge, countElems)
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
				edges[count] = &models.CartEdge{
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
				edges[count] = &models.CartEdge{
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
				edges[count] = &models.CartEdge{
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

	pc := &models.CartConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *CartsRepo) GetCartsByFiled(field, value string) (*models.Cart, error) {
	var item models.Cart
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *CartsRepo) GetCartsByID(id string) (*models.Cart, error) {
	return r.GetCartsByFiled("id", id)
}

func (r *CartsRepo) CartsSave(item *models.Cart) (*models.CartPayload, error) {
	_, err := r.DB.Model(item).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.CartPayload{
		Record:   item,
		RecordID: item.ID,
	}
	return placePayload, err
}

func (r *CartsRepo) CartsUpdate(place *models.Cart) (*models.CartPayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.CartPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CartsRepo) CartsPublish(place *models.Cart) (*models.CartPayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.CartPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CartsRepo) CartsRestore(place *models.Cart) (*models.CartPayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.CartPayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *CartsRepo) CartsDelete(place *models.Cart) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}

func (r *CartsRepo) GetKitsFoCart(obj *models.Cart) ([]*models.Kit, error) {
	var items []*models.Kit
	err := r.DB.Model(&items).Where("id in (?)", pg.In(obj.KitIds)).Select()
	return items, err
}
