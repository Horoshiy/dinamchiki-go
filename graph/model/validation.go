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

func (r ArticleInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("title", r.Title)
	v.Required("description", r.Description)

	return v.IsValid(), v.Errors
}

func (r StadiumInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("name", r.Name)

	return v.IsValid(), v.Errors
}

func (r StaffInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("name", r.Name)
	return v.IsValid(), v.Errors
}

func (r TeamInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("name", r.Name)
	return v.IsValid(), v.Errors
}

func (r TrainingInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("teamId", r.TeamID)
	return v.IsValid(), v.Errors
}

func (r CreatorInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("name", r.Name)
	return v.IsValid(), v.Errors
}

func (r StudentInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("name", r.Name)
	return v.IsValid(), v.Errors
}

func (r StudentVisitInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r MoneyMoveInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r MoneyCostInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r CoachPaymentByMonthInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r CoachPaymentByTeamInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r CoachPaymentByTrainingInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (l LoginInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("phone", l.Phone)
	v.IsPhone("phone", l.Phone)

	v.Required("password", l.Password)

	return v.IsValid(), v.Errors
}

func (r RentPaymentByMonthInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r RentPaymentByTrainingInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r TrainingDayInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r ClubBalanceInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r TeamBalanceInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r TaskInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r LeadInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r KitInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r CartInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}

func (r OrderInput) Validate() (bool, map[string]string) {
	v := validator.New()

	return v.IsValid(), v.Errors
}
