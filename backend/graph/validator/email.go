package validator

import "regexp"

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var emailme = "kionako96@gmail.com"

func (v *Validator) IsEmail(field, email string) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	// マッチしなかったらtrue
	if !emailRegexp.MatchString(email) {
		v.Errors[field] = "not a valid email"
		return false
	}

	return true
}
