package graph

import (
	"context"
	"github.com/go-pg/pg/v10"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
	"net/http"
	"time"
)

const loadersKey = "dataLoaders"
const userLoaderKey = "userLoader"
const placeLoaderKey = "placeLoader"

type Loaders struct {
	UserLoader
	PlaceLoader
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			userLoader,
			placeLoader,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
