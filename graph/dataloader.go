package graph

import (
	"context"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
	"net/http"
	"time"
)

const userLoaderKey = "userLoader"

func DataLoaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		ctx := context.WithValue(r.Context(), userLoaderKey, &userLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
