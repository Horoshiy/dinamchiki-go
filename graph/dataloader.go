package graph

import (
	"context"
	"github.com/go-pg/pg/v10"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
	"net/http"
	"time"
)

const loadersKey = "dataLoaders"

type Loaders struct {
	UserLoader
	PlaceLoader
	StaffSliceLoader
	StaffLoader
	StadiumLoader
	TeamLoader
	CreatorLoader
}

func DataLoaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	userLoader := UserLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(keys []string) ([]*models.User, []error) {
			var users []*models.User

			err := db.Model(&users).Where("id in (?)", pg.In(keys)).Select()

			if err != nil {
				return nil, []error{err}
			}

			u := make(map[string]*models.User, len(users))

			for _, user := range users {
				u[user.ID] = user
			}

			result := make([]*models.User, len(keys))

			for i, id := range keys {
				result[i] = u[id]
			}

			return result, nil
		},
	}

	placeLoader := PlaceLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(keys []string) ([]*models.Place, []error) {
			var items []*models.Place

			err := db.Model(&items).Where("id in (?)", pg.In(keys)).Select()

			if err != nil {
				return nil, []error{err}
			}

			u := make(map[string]*models.Place, len(items))

			for _, item := range items {
				u[item.ID] = item
			}

			result := make([]*models.Place, len(keys))

			for i, id := range keys {
				result[i] = u[id]
			}

			return result, nil
		},
	}

	creatorLoader := CreatorLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(keys []string) ([]*models.Creator, []error) {
			var items []*models.Creator

			err := db.Model(&items).Where("id in (?)", pg.In(keys)).Select()

			if err != nil {
				return nil, []error{err}
			}

			u := make(map[string]*models.Creator, len(items))

			for _, item := range items {
				u[item.ID] = item
			}

			result := make([]*models.Creator, len(keys))

			for i, id := range keys {
				result[i] = u[id]
			}

			return result, nil
		},
	}

	stadiumLoader := StadiumLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(keys []string) ([]*models.Stadium, []error) {
			var items []*models.Stadium

			err := db.Model(&items).Where("id in (?)", pg.In(keys)).Select()

			if err != nil {
				return nil, []error{err}
			}

			u := make(map[string]*models.Stadium, len(items))

			for _, item := range items {
				u[item.ID] = item
			}

			result := make([]*models.Stadium, len(keys))

			for i, id := range keys {
				result[i] = u[id]
			}

			return result, nil
		},
	}

	staffLoader := StaffLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(keys []string) ([]*models.Staff, []error) {
			var items []*models.Staff

			err := db.Model(&items).Where("id in (?)", pg.In(keys)).Select()

			if err != nil {
				return nil, []error{err}
			}

			u := make(map[string]*models.Staff, len(items))

			for _, item := range items {
				u[item.ID] = item
			}

			result := make([]*models.Staff, len(keys))

			for i, id := range keys {
				result[i] = u[id]
			}

			return result, nil
		},
	}

	teamLoader := TeamLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(keys []string) ([]*models.Team, []error) {
			var items []*models.Team

			err := db.Model(&items).Where("id in (?)", pg.In(keys)).Select()

			if err != nil {
				return nil, []error{err}
			}

			u := make(map[string]*models.Team, len(items))

			for _, item := range items {
				u[item.ID] = item
			}

			result := make([]*models.Team, len(keys))

			for i, id := range keys {
				result[i] = u[id]
			}

			return result, nil
		},
	}

	staffSliceLoader := StaffSliceLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(keys []string) ([][]*models.Staff, []error) {
			staff := make([][]*models.Staff, len(keys))
			errors := make([]error, len(keys))

			var items []*models.Staff

			err := db.Model(&items).Where("id in (?)", pg.In(keys)).Select()

			if err != nil {
				return nil, []error{err}
			}

			u := make(map[string]*models.Staff, len(items))

			for _, item := range items {
				u[item.ID] = item
			}
			result := make([]*models.Staff, len(keys))

			for i, _ := range keys {
				staff[i] = result
			}
			return staff, errors
		},
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			UserLoader:       userLoader,
			PlaceLoader:      placeLoader,
			StaffSliceLoader: staffSliceLoader,
			StaffLoader:      staffLoader,
			StadiumLoader:    stadiumLoader,
			TeamLoader:       teamLoader,
			CreatorLoader:    creatorLoader,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
