//go:generate go run github.com/99designs/gqlgen --verbose

package graph

import "gitlab.com/dinamchiki/go-graphql/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Domain *domain.Domain
}
