package domain

import (
	"context"
	"errors"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
	"log"
)

// Login is the resolver for the login field.
func (d *Domain) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := d.UsersRepo.GetUserByPhone(input.Phone)
	if err != nil {
		return nil, ErrBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, ErrBadCredentials
	}

	token, err := user.GenToken()
	if err != nil {
		return nil, errors.New("Something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil

}

// Register is the resolver for the register field.
func (d *Domain) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	_, err := d.UsersRepo.GetUserByPhone(input.Phone)
	if err == nil {
		return nil, errors.New("Такой номер телефона уже зарегистрирован за пользователем")
	}

	//var roles []models.Role

	user := &models.User{
		Phone:     input.Phone,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Roles:     []models.Role{models.RoleUser},
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("somethig wentwrong")
	}

	//TODO: create verification code
	tx, err := d.UsersRepo.DB.Begin()
	if err != nil {
		log.Printf("error creating a transaction: %v", err)
		return nil, errors.New("something went wrong ")
	}
	defer tx.Rollback()

	if _, err := d.UsersRepo.CreateUser(tx, user); err != nil {
		log.Printf("error creating a user: %v", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("error while commiting: %v", err)
		return nil, err
	}

	token, err := user.GenToken()
	if err != nil {
		log.Printf("error while generating the token: %v", err)

		return nil, errors.New("somethin went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
