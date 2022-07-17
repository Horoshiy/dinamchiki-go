package validator

import "regexp"

var phoneRegexp = regexp.MustCompile("^\\+7[0-9]{10}$")

func (v *Validator) IsPhone(field, phone string) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if !phoneRegexp.MatchString(phone) {
		v.Errors[field] = "неправильный формат номера телефона"
		return false
	}

	return true
}
