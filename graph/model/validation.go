package models

import "gitlab.com/dinamchiki/go-graphql/validator"

func (r RegisterInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("phone", r.Phone)
	v.IsPhone("phone", r.Phone)

	v.Required("password", r.Password)
	v.MinLength("password", r.Password, 6)

	v.Required("confirmPassword", r.ConfirmPassword)
	v.EqualToField("confirmPassword", r.ConfirmPassword, "password", r.Password)

	v.Required("firstName", r.FirstName)
	v.MinLength("firstName", r.FirstName, 2)

	v.Required("lastName", r.LastName)
	v.MinLength("lastName", r.LastName, 2)

	return v.IsValid(), v.Errors
}

func (r PlaceInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("name", r.Name)
	v.Required("address", r.Address)
	v.Required("description", r.Description)

	return v.IsValid(), v.Errors
}

func (l LoginInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("phone", l.Phone)
	v.IsPhone("phone", l.Phone)

	v.Required("password", l.Password)

	return v.IsValid(), v.Errors
}
