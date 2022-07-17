package domain

import (
	"errors"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
	"gitlab.com/dinamchiki/go-graphql/postgres"
)

var (
	ErrBadCredentials  = errors.New("email/password combination dont'n work")
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrForbidden       = errors.New("unauthorized")
)

type Domain struct {
	UsersRepo   postgres.UsersRepo
	MeetupsRepo postgres.MeetupsRepo
	PlacesRepo  postgres.PlacesRepo
}

func NewDomain(usersRepo postgres.UsersRepo, meetupsRepo postgres.MeetupsRepo, placesRepo postgres.PlacesRepo) *Domain {
	return &Domain{UsersRepo: usersRepo, MeetupsRepo: meetupsRepo, PlacesRepo: placesRepo}
}

type Ownable interface {
	IsOwner(user *models.User) bool
}

func checkOwnership(o Ownable, user *models.User) bool {
	return o.IsOwner(user)
}
