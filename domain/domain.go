package domain

import (
	"errors"
	"gitlab.com/dinamchiki/go-graphql/postgres"
)

var (
	ErrBadCredentials  = errors.New("email/password combination dont'n work")
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrForbidden       = errors.New("unauthorized")
)

type Domain struct {
	UsersRepo    postgres.UsersRepo
	PlacesRepo   postgres.PlacesRepo
	ArticlesRepo postgres.ArticlesRepo
	StadiumsRepo postgres.StadiumsRepo
	StaffRepo    postgres.StaffRepo
}

func NewDomain(
	usersRepo postgres.UsersRepo,
	placesRepo postgres.PlacesRepo,
	articlesRepo postgres.ArticlesRepo,
	stadiumsRepo postgres.StadiumsRepo,
	staffRepo postgres.StaffRepo,
) *Domain {
	return &Domain{
		UsersRepo:    usersRepo,
		PlacesRepo:   placesRepo,
		ArticlesRepo: articlesRepo,
		StadiumsRepo: stadiumsRepo,
		StaffRepo:    staffRepo,
	}
}
