package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v10"
	"github.com/rs/cors"
	"gitlab.com/dinamchiki/go-graphql/domain"
	"gitlab.com/dinamchiki/go-graphql/graph"
	"gitlab.com/dinamchiki/go-graphql/graph/generated"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
	customMiddleware "gitlab.com/dinamchiki/go-graphql/middleware"
	"gitlab.com/dinamchiki/go-graphql/postgres"
	"log"
	"net/http"
	"os"
)

const defaultPort = "9090"

func main() {
	DB := postgres.New(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userRepo := postgres.UsersRepo{DB: DB}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("CORS_ALLOWED_SITES")},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(userRepo))

	d := domain.NewDomain(
		userRepo,
		postgres.PlacesRepo{DB: DB},
		postgres.ArticlesRepo{DB: DB},
		postgres.StadiumsRepo{DB: DB},
	)
	c := generated.Config{
		Resolvers: &graph.Resolver{
			Domain: d,
		},
	}
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role models.Role) (res interface{}, err error) {
		currentUser, err := customMiddleware.GetCurrentUserFromCtx(ctx)
		if err != nil {
			return nil, fmt.Errorf("доступ запрещен")
		}
		if !currentUser.HasRole(role) {
			return nil, fmt.Errorf("уровень доступа недостаточный")
		}

		return next(ctx)

	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", graph.DataLoaderMiddleware(DB, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
