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
	UsersRepo                  postgres.UsersRepo
	PlacesRepo                 postgres.PlacesRepo
	ArticlesRepo               postgres.ArticlesRepo
	StadiumsRepo               postgres.StadiumsRepo
	StaffRepo                  postgres.StaffRepo
	TeamsRepo                  postgres.TeamsRepo
	TrainingsRepo              postgres.TrainingsRepo
	CreatorsRepo               postgres.CreatorsRepo
	StudentsRepo               postgres.StudentsRepo
	StudentVisitsRepo          postgres.StudentVisitsRepo
	MoneyMoveRepo              postgres.MoneyMovesRepo
	MoneyCostRepo              postgres.MoneyCostsRepo
	CoachPaymentByMonthRepo    postgres.CoachPaymentByMonthRepo
	CoachPaymentByTeamRepo     postgres.CoachPaymentByTeamRepo
	CoachPaymentByTrainingRepo postgres.CoachPaymentByTrainingRepo
}

func NewDomain(
	usersRepo postgres.UsersRepo,
	placesRepo postgres.PlacesRepo,
	articlesRepo postgres.ArticlesRepo,
	stadiumsRepo postgres.StadiumsRepo,
	staffRepo postgres.StaffRepo,
	teamsRepo postgres.TeamsRepo,
	trainingsRepo postgres.TrainingsRepo,
	creatorsRepo postgres.CreatorsRepo,
	studentsRepo postgres.StudentsRepo,
	studentVisitsRepo postgres.StudentVisitsRepo,
	moneyMoveRepo postgres.MoneyMovesRepo,
	moneyCostsRepo postgres.MoneyCostsRepo,
	coachPaymentByMonthRepo postgres.CoachPaymentByMonthRepo,
	coachPaymentByTeamRepo postgres.CoachPaymentByTeamRepo,
	coachPaymentByTrainingRepo postgres.CoachPaymentByTrainingRepo,
) *Domain {
	return &Domain{
		UsersRepo:                  usersRepo,
		PlacesRepo:                 placesRepo,
		ArticlesRepo:               articlesRepo,
		StadiumsRepo:               stadiumsRepo,
		StaffRepo:                  staffRepo,
		TeamsRepo:                  teamsRepo,
		TrainingsRepo:              trainingsRepo,
		CreatorsRepo:               creatorsRepo,
		StudentsRepo:               studentsRepo,
		StudentVisitsRepo:          studentVisitsRepo,
		MoneyMoveRepo:              moneyMoveRepo,
		MoneyCostRepo:              moneyCostsRepo,
		CoachPaymentByMonthRepo:    coachPaymentByMonthRepo,
		CoachPaymentByTeamRepo:     coachPaymentByTeamRepo,
		CoachPaymentByTrainingRepo: coachPaymentByTrainingRepo,
	}
}
