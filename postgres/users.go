package postgres

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type UsersRepo struct {
	DB *pg.DB
}

func (u *UsersRepo) GetUserByFiled(field, value string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &user, err
}

func (u *UsersRepo) GetUserByID(id string) (*models.User, error) {
	return u.GetUserByFiled("id", id)

}

func (u *UsersRepo) GetUserByPhone(phone string) (*models.User, error) {
	return u.GetUserByFiled("phone", phone)
}

func (u *UsersRepo) CreateUser(tx *pg.Tx, user *models.User) (*models.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()
	return user, err
}
